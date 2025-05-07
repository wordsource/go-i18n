package i18n

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/text/language"
)

func Test_Localizef(t *testing.T) {
	t.Run("handles_no_registered_language_bundles", func(t *testing.T) {
		ctx := WithLocale(context.Background(), language.AmericanEnglish)
		message := "Hello, World!"
		expect := "Hello, World!"

		actual := Localizef(ctx, message, nil)
		require.Equal(t, expect, actual)
	})

	t.Run("handles_no_registered_language_bundles_with_values", func(t *testing.T) {
		ctx := WithLocale(context.Background(), language.AmericanEnglish)
		message := "Hello, {name}!"
		values := Data{"name": "Charles"}
		expect := "Hello, Charles!"

		actual := Localizef(ctx, message, values)
		require.Equal(t, expect, actual)
	})

	t.Run("handles_registered_language_bundles", func(t *testing.T) {
		ctx := WithLocale(context.Background(), language.Spanish)
		message := "Hello, World!"
		expect := "Hola, Mundo!"

		bundle := LanguageBundle{
			Message{Message: message}.Hash(): {Message: "Hola, Mundo!"},
		}

		GetLocalizer().RegisterBundle(language.Spanish, bundle)

		actual := Localizef(ctx, message, nil)
		require.Equal(t, expect, actual)
	})

	t.Run("handles_formatted_data", func(t *testing.T) {
		ctx := WithLocale(context.Background(), language.Spanish)
		message := "Hello, {name}!"
		values := Data{"name": "Charles"}
		expect := "Hola, Charles!"

		bundle := LanguageBundle{
			Message{Message: message}.Hash(): {Message: "Hola, {name}!"},
		}

		GetLocalizer().RegisterBundle(language.Spanish, bundle)

		actual := Localizef(ctx, message, values)
		require.Equal(t, expect, actual)
	})

	t.Run("handles_formatted_data_with_description", func(t *testing.T) {
		ctx := WithLocale(context.Background(), language.Spanish)
		message := "Hello, {name}!"
		description := "A friendly greeting"
		values := Data{"name": "Charles"}
		expect := "Hola, Charles!"

		bundle := LanguageBundle{
			Message{Message: message, Description: description}.Hash(): {Message: "Hola, {name}!", Description: description},
		}

		GetLocalizer().RegisterBundle(language.Spanish, bundle)

		actual := LocalizeDf(ctx, message, description, values)
		require.Equal(t, expect, actual)
	})

	t.Run("icu_formatting", func(t *testing.T) {
		t.Run("handles_plural", func(t *testing.T) {
			ctx := WithLocale(context.Background(), language.Spanish)
			enMessage := `You have {itemCount, plural,
			    =0 {no projects}
			    one {# project}
				=2 {# projects}
			    other {# projects}
			}.`
			esMessage := `Tienes {itemCount, plural,
			    =0 {ningún proyecto}
			    one {# proyecto}
			    other {# proyectos}
			}.`
			bundle := LanguageBundle{
				Message{Message: enMessage}.Hash(): {Message: esMessage},
			}

			GetLocalizer().RegisterBundle(language.Spanish, bundle)

			require.Equal(t, "Tienes ningún proyecto.", Localizef(ctx, enMessage, Data{"itemCount": 0}))
			require.Equal(t, "Tienes 1 proyecto.", Localizef(ctx, enMessage, Data{"itemCount": 1}))
			require.Equal(t, "Tienes 2 proyectos.", Localizef(ctx, enMessage, Data{"itemCount": 2}))
			require.Equal(t, "Tienes 5 proyectos.", Localizef(ctx, enMessage, Data{"itemCount": 5}))
		})

		t.Run("handles_select_ordinal", func(t *testing.T) {
			ctx := WithLocale(context.Background(), language.Spanish)
			enMessage := `Congrats! It's your {year, selectordinal,
			    one {#st}
			    two {#nd}
			    few {#rd}
			    other {#th}
			} subscription anniversary!`
			esMessage := `¡Felicidades! Es tu {year, selectordinal,
			    one {#er}
			    two {#do}
			    few {#er}
			    other {#to}
			} aniversario de suscripción.`
			bundle := LanguageBundle{
				Message{Message: enMessage}.Hash(): {Message: esMessage},
			}

			GetLocalizer().RegisterBundle(language.Spanish, bundle)

			require.Equal(t, "¡Felicidades! Es tu 1er aniversario de suscripción.", Localizef(ctx, enMessage, Data{"year": 1}))
			require.Equal(t, "¡Felicidades! Es tu 2do aniversario de suscripción.", Localizef(ctx, enMessage, Data{"year": 2}))
			require.Equal(t, "¡Felicidades! Es tu 3er aniversario de suscripción.", Localizef(ctx, enMessage, Data{"year": 3}))
			require.Equal(t, "¡Felicidades! Es tu 5to aniversario de suscripción.", Localizef(ctx, enMessage, Data{"year": 5}))
		})

		t.Run("handles_select_ordinal", func(t *testing.T) {
			ctx := WithLocale(context.Background(), language.Spanish)
			enMessage := `My pronouns are {gender, select,
			    male {He/him}
			    female {She/her}
			    other {They/them}
			}.`
			esMessage := `Mis pronombres son {gender, select,
			    male {Él}
			    female {Ella}
			    other {Elle}
			}.`
			bundle := LanguageBundle{
				Message{Message: enMessage}.Hash(): {Message: esMessage},
			}

			GetLocalizer().RegisterBundle(language.Spanish, bundle)

			require.Equal(t, "Mis pronombres son Él.", Localizef(ctx, enMessage, Data{"gender": "male"}))
			require.Equal(t, "Mis pronombres son Ella.", Localizef(ctx, enMessage, Data{"gender": "female"}))
			require.Equal(t, "Mis pronombres son Elle.", Localizef(ctx, enMessage, Data{"gender": "other"}))
		})
	})
}
