package i18n

import (
	"context"
	"html/template"

	"golang.org/x/text/language"
)

// FuncMap returns a template.FuncMap that can be used with the html/template package
// to localize messages in templates. The FuncMap provides two functions:
// - localize: localizes a message using the locale in context
// - localizeD: localizes a message using the locale in context, with a description
//
// The i18n extractor knows to look for template files when configured.
func FuncMap(ctx context.Context) template.FuncMap {
	return template.FuncMap{
		"localize": func(message string) string {
			return Localize(ctx, message)
		},
		"localizeD": func(message string, description string) string {
			return LocalizeD(ctx, message, description)
		},
	}
}

func DefaultFuncMap(tag language.Tag) template.FuncMap {
	return template.FuncMap{
		"localize": func(message string) string {
			return ""
		},
		"localizeD": func(message string, description string) string {
			return ""
		},
	}
}
