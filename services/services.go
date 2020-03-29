package services
import (
	"awesomeProject/domain"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv")


//BASEURL,ERRORS AND GLOBAL VARIABLES
var (
	baseurl                 string = "https://www.breakingbadapi.com/api/"
	ErrorEpisodesNotFound          = errors.New("Episodes NOT-FOUND ")
	ErrorCharactersNotFound        = errors.New("Characters NOT-FOUND")
	ErrorQuotesNotFound            = errors.New("Quotes NOT-FOUND")
	ErrorDeathsNotFound            = errors.New("Deaths NOT-FOUND")

	ErrorInvalidID        = errors.New("Invalid ID")
	ErrorExistsCharacters = errors.New("Characters Exists")

	characters []domain.CharactersBB
	episodes   []domain.EpisodesBB
	quotes     []domain.QuotesBB
	deaths          []domain.DeathsBB
)

/*****************************************************************************/
/*******************************BREAKING BAD API******************************/
/*****************************************************************************/

//DEVUELVE TODOS LOS PERSONAJES DE LA SERIE BREAKING BAD
func GetCharacters() ([]domain.CharactersBB, error) {
	Data, _ := ioutil.ReadFile("./DB.json")
	json.Unmarshal(Data, &characters)
	return characters, nil
}

//This a function about the Delete Characters for ID of serie the Breaking Bad
func DeleteCharactersID(i string) (domain.CharactersBB, error) {
	id, err := validateID(i)
	if err != nil {
		return domain.CharactersBB{}, ErrorInvalidID
	}

	character, err := SearchCharacters(id)
	if err != nil {
		return domain.CharactersBB{}, err
	}

	for i := 0; i < len(characters); i++ {
		if characters[i].CharID == character.CharID {
			characters = append(characters[:i], characters[i+1:]...)
			break
		}
	}
	return character, nil
}

//this a function returns
func GetCharactersID(i string) (domain.CharactersBB, error) {
	id, err := validateID(i)
	if err != nil {
		return domain.CharactersBB{}, ErrorInvalidID
	}

	character, err := SearchCharacters(id)
	if err != nil {
		return domain.CharactersBB{}, err
	}

	return character, nil

}

func PostCharacters(character domain.CharactersBB) (domain.CharactersBB, error) {
	archivo, _ := json.Marshal(character)
	_ = ioutil.WriteFile("DB.json", archivo, 0664)
	characters = append(characters, character)

	return character, nil
}

func PutCharacters(character domain.CharactersBB) (domain.CharactersBB, error) {
	err := ExistsCharacters(character)
	if err != nil {
		for i := 0; i < len(characters); i++ {
			if characters[i].CharID == character.CharID {
				characters[i].Name = character.Name
				characters[i].Birthday = character.Birthday
				characters[i].Occupation = character.Occupation
				characters[i].Img = character.Img
				characters[i].Status = character.Status
				characters[i].Appearance = character.Appearance
				characters[i].Portrayed = character.Portrayed
				characters[i].Category = character.Category
				characters[i].BetterCallSaulAppearance = character.BetterCallSaulAppearance

				return characters[i], nil
			}
		}
	}

	return domain.CharactersBB{}, ErrorCharactersNotFound
}

/*there are function auxiliary for Characters*/
/*This a auxiliary function about the search Characters of Breaking Bad*/
func SearchCharacters(id int) (domain.CharactersBB, error) {
	for _, c := range characters {
		if c.CharID == id {
			return c, nil
		}
	}
	return domain.CharactersBB{}, ErrorCharactersNotFound
}
func validateID(id string) (int, error) {
	num, err := strconv.Atoi(id)
	if err != nil {
		return -1, err
	}
	return num, nil
}

/*This a function about if exists some characters storages */
func ExistsCharacters(character domain.CharactersBB) error {
	for _, c := range characters {
		if c.CharID == character.CharID {
			return ErrorExistsCharacters
		}
	}
	return nil
}

/*This a function that returns all episodes of the serie Breaking Bad*/
func GetEpisodes() (domain.EpisodesBB, error) {
	var episodes domain.EpisodesBB
	url := baseurl + "episodes"
	resp, err1 := http.Get(url)

	Data, err2 := ioutil.ReadAll(resp.Body)

	if err1 != nil || err2 != nil {
		return episodes, ErrorEpisodesNotFound
	}
	json.Unmarshal(Data, &episodes)
	return episodes, nil

}

/*This a function that returns an episodes with id of the serie Breaking Bad*/
func GetEpisodesID(id int) (domain.EpisodesBB, error) {
	var episodes domain.EpisodesBB
	url := baseurl + "episodes/" + strconv.Itoa(id)
	resp, err1 := http.Get(url)

	Data, err2 := ioutil.ReadAll(resp.Body)

	if err1 != nil || err2 != nil {
		return episodes, ErrorEpisodesNotFound
	}
	json.Unmarshal(Data, &episodes)
	return episodes, nil

}

/*This a function that returns all quotes of the serie Breaking Bad*/
func GetQuotes() (domain.QuotesBB, error) {
	var quotes domain.QuotesBB
	url := baseurl + "quotes"
	resp, err1 := http.Get(url)

	Data, err2 := ioutil.ReadAll(resp.Body)

	if err1 != nil || err2 != nil {
		return quotes, ErrorEpisodesNotFound
	}
	json.Unmarshal(Data, &quotes)
	return quotes, nil

}

/*This a function that returns an quootes with id of the serie Breaking Bad*/
func GetQuotesID(id int) (domain.QuotesBB, error) {
	var quotes domain.QuotesBB
	url := baseurl + "quotes/" + strconv.Itoa(id)
	resp, err1 := http.Get(url)

	Data, err2 := ioutil.ReadAll(resp.Body)

	if err1 != nil || err2 != nil {
		return quotes, ErrorQuotesNotFound
	}
	json.Unmarshal(Data, &quotes)
	return quotes, nil

}

/*This a function that returns all deaths of the serie Breaking Bad*/

func GetDeaths() (domain.DeathsBB, error) {
	var deaths domain.DeathsBB
	url := baseurl + "deaths"
	resp, err1 := http.Get(url)

	Data, err2 := ioutil.ReadAll(resp.Body)

	if err1 != nil || err2 != nil {
		return deaths, ErrorDeathsNotFound
	}
	json.Unmarshal(Data, &deaths)
	return deaths, nil

}
