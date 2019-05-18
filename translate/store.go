package translate

// Language is defined for all available languages
type Language string

const (
	German  Language = "de_DE"
	English Language = "en_US"
)

// Key must be defined for all texts retuned by the API that will be read by end users
type Key string

const (
	LoginUserNotFound        Key = "Login.Error.UserNotFound"
	ImportParseError         Key = "Import.Error.Parse"
	ImportCreateProductError Key = "Import.Error.Transform"
	RegisterMailerError      Key = "Register.Error.Mailer"
	ServerError              Key = "Error.Server"
	UnauthorizedError        Key = "Error.Unauthorized"
	NotFoundError            Key = "Error.NotFound"
	MissingParameter         Key = "Error.MissingParameter"
	MissingBody              Key = "Error.MissingBody"
	BodyParserError          Key = "Error.BodyParser"
)

var mapping = map[Language]map[Key]string{
	German: {
		LoginUserNotFound:        "Benutzer konnte nicht gefunden werden.",
		ServerError:              "Es ist ein Fehler aufgetreten. Bitte versuchen Sie es später erneut.",
		UnauthorizedError:        "Zugriff nicht gestattet.",
		ImportCreateProductError: "HLK kann nicht zu Produkt umgeformed werden",
		ImportParseError:         "HLK kann nicht gelesen werden",
		NotFoundError:            "Die angefragte Resource kann nicht gefunden werden",
		MissingParameter:         "Fehlender Anfrageparameter",
		MissingBody:              "Fehlender Anfrageinhalt",
		BodyParserError:          "Anfragedaten können nicht gelesen werden",
		RegisterMailerError:      "Es gab Probleme beim Versand der Bestätigungs EMail. Wir arbeiten an einer Lösung für dieses Problem und verschicken die Mail sobald wie möglich.",
	},
	English: {
		LoginUserNotFound:        "User could not be found.",
		ServerError:              "Error processing request. Please try again at a later date.",
		UnauthorizedError:        "Access denied.",
		ImportCreateProductError: "HLK node cannot be transformed to product",
		ImportParseError:         "HLK cannot be read from file",
		NotFoundError:            "The requested resource could not be found",
		MissingParameter:         "Missing request parameter",
		MissingBody:              "Missing request body",
		BodyParserError:          "Error parsing request body",
		RegisterMailerError:      "OptIn mail could not be send. We are working on a solution and will send the mail as soon as possible.",
	},
}

// Get fetches the translation for the given language
func Get(lang Language, key Key) string {
	langMapping, ok := mapping[lang]
	if ok == false {
		return string(key)
	}

	if v, ok := langMapping[key]; ok == true {
		return v
	}

	return string(key)
}
