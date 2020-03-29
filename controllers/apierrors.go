package controllers
import (
	"awesomeProject/services"
)

/*This a struct of API ERROR for have a best control of the errors in the services.go  */

type ApiErrors struct {
	Status  int
	Message string
}

func (e *ApiErrors) Error() string {
	return e.Message
}

func parseErrors(e error) ApiErrors {
	switch e {
	case services.ErrorInvalidID, services.ErrorExistsCharacters:
		return ApiErrors{400, e.Error()}
	case services.ErrorCharactersNotFound, services.ErrorQuotesNotFound, services.ErrorEpisodesNotFound, services.ErrorDeathsNotFound:
		return ApiErrors{404, e.Error()}
	default:
		return ApiErrors{500, e.Error()}

	}
}
