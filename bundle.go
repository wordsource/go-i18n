package i18n

import (
	"embed"
	"encoding/hex"
	"encoding/json"
	"hash/fnv"
	"path/filepath"
	"strings"

	"golang.org/x/text/language"
)

type Message struct {
	Message     string `json:"message"`
	Description string `json:"description,omitempty"`

	// Source is the file path where this message was extracted from,
	// relative to the working directory at the time of extraction.
	Source string `json:"source,omitempty"`
	// Parser is the parser used to extract this message from the source.
	Parser string `json:"parser,omitempty"`
}

func (m Message) Hash() string {
	hash := fnv.New64a()
	_, _ = hash.Write([]byte(m.Message + m.Description))
	return hex.EncodeToString(hash.Sum(nil))
}

func (m Message) String() string {
	return m.Message
}

type LanguageBundle map[string]Message

func (b LanguageBundle) Merge(other LanguageBundle) {
	for k, v := range other {
		b[k] = v
	}
}

func parseBundle(byt []byte) (LanguageBundle, error) {
	var (
		bundle = make(LanguageBundle)
		err    = json.Unmarshal(byt, &bundle)
	)
	return bundle, err
}

func LoadBundles(fs embed.FS) error {
	entries, err := fs.ReadDir(".")
	if err != nil {
		return err
	}

	var filePath = ""
	if len(entries) == 1 && entries[0].IsDir() && entries[0].Name() == ".i18n" {
		entries, err = fs.ReadDir(entries[0].Name())
		if err != nil {
			return err
		}
		filePath = ".i18n/"
	}

	for _, entry := range entries {
		var name = entry.Name()

		if entry.IsDir() || name == "defaultMessages.json" {
			continue
		}

		byt, err := fs.ReadFile(filepath.Join(filePath, name))
		if err != nil {
			return err
		}

		tag, err := language.Parse(strings.TrimSuffix(name, ".json"))
		if err != nil {
			return err
		}

		bundle, err := parseBundle(byt)
		if err != nil {
			return err
		}

		GetLocalizer().Register(tag, bundle)
	}

	return nil
}

func MustLoadBundles(fs embed.FS) {
	if err := LoadBundles(fs); err != nil {
		panic(err)
	}
}
