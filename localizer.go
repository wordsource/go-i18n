package i18n

import "golang.org/x/text/language"

type Localizer interface {
	GetBundle(language.Tag) LanguageBundle
	RegisterBundle(language.Tag, LanguageBundle)
	SupportsLocale(language.Tag) bool
	ParseAcceptLanguage(string) ([]language.Tag, []float32, error)
}

type localizerImpl struct {
	bundles map[language.Tag]LanguageBundle
}

var (
	instance      Localizer
	DefaultLocale = language.AmericanEnglish
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

func (l *localizerImpl) RegisterBundle(locale language.Tag, bundle LanguageBundle) {
	if l.bundles == nil {
		l.bundles = make(map[language.Tag]LanguageBundle)
	}
	if l.bundles[locale] != nil {
		l.bundles[locale].Merge(bundle)
	} else {
		l.bundles[locale] = bundle
	}
}

func (l *localizerImpl) SupportsLocale(locale language.Tag) bool {
	if locale == DefaultLocale {
		return true
	}
	if l.bundles == nil {
		return false
	}
	_, ok := l.bundles[locale]
	return ok
}

func (l *localizerImpl) ParseAcceptLanguage(accept string) ([]language.Tag, []float32, error) {
	acceptLocales, acceptWeights, err := language.ParseAcceptLanguage(accept)
	if err != nil {
		return nil, nil, err
	}

	supportedLocales := make([]language.Tag, 0)
	supportedWeights := make([]float32, 0)
	for idx, tag := range acceptLocales {
		if l.SupportsLocale(tag) {
			supportedLocales = append(supportedLocales, tag)
			supportedWeights = append(supportedWeights, acceptWeights[idx])
		}
	}

	return supportedLocales, supportedWeights, nil
}
