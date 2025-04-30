package i18n

import (
	"context"

	"github.com/wordsource/go-messageformat"
	"golang.org/x/text/language"
)

var DefaultLocale = language.AmericanEnglish

// Localize localizes the given message using the locale in context. It returns the given message
// if the locale is not found or a translated version of the message is not found in that locale.
func Localize(ctx context.Context, message string) string {
	bundle := GetLocalizer().GetBundle(GetLocale(ctx))
	if bundle == nil {
		return message
	}

	messageKey := Message{Message: message}.Hash()
	localizedMessage, ok := bundle[messageKey]
	if !ok {
		return message
	}

	return localizedMessage.Message
}

// LocalizeD is like Localize but allows the developer to pass a description of the message
// to be used during translation, like additional context to a human translator (or LLM prompt)
//
// The description is extracted along with the message, and both are hashed together to create
// the message identifier. Both are stored in the defaultMessage.json file.
func LocalizeD(ctx context.Context, message string, description string) string {
	bundle := GetLocalizer().GetBundle(GetLocale(ctx))
	if bundle == nil {
		return message
	}

	messageKey := Message{Message: message, Description: description}.Hash()
	localizedMessage, ok := bundle[messageKey]
	if !ok {
		return message
	}

	return localizedMessage.Message
}

// Values is a map of values to be used in message formatting. The key is the name of the value
// to be replaced in the message. Suitable values are strings, numbers, and other primitive types.
type Values map[string]interface{}

// Localizef localizes the given message using the locale in context, then formats the message
// by applying the given Values.
func Localizef(ctx context.Context, message string, values Values) string {
	localizedMessage := Localize(ctx, message)
	pt, err := messageformat.NewParser().Parse(localizedMessage)
	if err != nil {
		return localizedMessage
	}

	f, err := messageformat.NewFormatter()
	if err != nil {
		return localizedMessage
	}

	formattedMessage, err := f.FormatMap(pt, values)
	if err != nil {
		return localizedMessage
	}

	return formattedMessage
}

// LocalizeDf is like Localizef but allows the developer to pass a description of the message
// to be used during translation, like additional context to a human translator (or LLM prompt)
//
// The description is extracted along with the message, and both are hashed together to create
// the message identifier. Both are stored in the defaultMessage.json file.
func LocalizeDf(ctx context.Context, message string, description string, values Values) string {
	return Localizef(ctx, message, values)
}
