package i18n

import (
	"context"

	"golang.org/x/text/language"
)

type contextKey string

const (
	localeKey contextKey = "locale"
)

func GetLocale(ctx context.Context) language.Tag {
	locale, ok := ctx.Value(localeKey).(language.Tag)
	if !ok {
		return language.English
	}
	return locale
}

func WithLocale(ctx context.Context, locale language.Tag) context.Context {
	return context.WithValue(ctx, localeKey, locale)
}
