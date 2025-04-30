package i18n

import (
	"context"
	"golang.org/x/text/language"
	"testing"
)

func hashMessage(message string) string {
	return Message{Message: message}.Hash()
}

func Test_Localizef(t *testing.T) {
	var testTable = map[string]struct {
		Context context.Context
		Message string
		Values  Values
		Bundles map[language.Tag]LanguageBundle
		Expect  string
	}{
		"handles_no_registered_language_bundles": {
			Context: WithLocale(context.Background(), DefaultLocale),
			Message: "Hello, World!",
			Expect:  "Hello, World!",
		},
		"handles_no_registered_language_bundles_with_values": {
			Context: WithLocale(context.Background(), DefaultLocale),
			Message: "Hello, {name}!",
			Values:  Values{"name": "Charles"},
			Expect:  "Hello, Charles!",
		},
		"handles_registered_language_bundles": {
			Context: WithLocale(context.Background(), language.Spanish),
			Message: "Hello, World!",
			Bundles: map[language.Tag]LanguageBundle{
				language.Spanish: {
					hashMessage("Hello, World!"): {Message: "Hola, Mundo!"},
				},
			},
			Expect: "Hola, Mundo!",
		},
		"handles_registered_language_bundles_with_values": {
			Context: WithLocale(context.Background(), language.Spanish),
			Message: "Hello, {name}!",
			Values:  Values{"name": "Charles"},
			Bundles: map[language.Tag]LanguageBundle{
				language.Spanish: {
					hashMessage("Hello, {name}!"): {Message: "Hola, {name}!"},
				},
			},
			Expect: "Hola, Charles!",
		},
	}

	for name, test := range testTable {
		t.Run(name, func(t *testing.T) {
			for tag, bundle := range test.Bundles {
				GetLocalizer().Register(tag, bundle)
			}

			result := Localizef(test.Context, test.Message, test.Values)
			if result != test.Expect {
				t.Fatalf("expected %s, got %s", test.Expect, result)
			}
		})
	}
}
