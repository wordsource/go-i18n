package i18n

import "golang.org/x/text/language"

type Localizer interface {
	GetBundle(language.Tag) LanguageBundle
	Register(language.Tag, LanguageBundle)
}

type localizerImpl struct {
	bundles map[language.Tag]LanguageBundle
}

var (
	instance Localizer
)

func GetLocalizer() Localizer {
	if instance == nil {
		instance = &localizerImpl{
			bundles: make(map[language.Tag]LanguageBundle),
		}
	}
	return instance
}

func (l *localizerImpl) GetBundle(locale language.Tag) LanguageBundle {
	return l.bundles[locale]
}

func (l *localizerImpl) Register(locale language.Tag, bundle LanguageBundle) {
	if l.bundles == nil {
		l.bundles = make(map[language.Tag]LanguageBundle)
	}
	if l.bundles[locale] != nil {
		l.bundles[locale].Merge(bundle)
	} else {
		l.bundles[locale] = bundle
	}
}
