package controllers

/*****************************CONTROLLERS**********************************/
import (
	"awesomeProject/domain"
	"awesomeProject/services"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

//ASDASDASD
func GetCharacters(c *gin.Context) {
	character, err := services.GetCharacters()
	if err != nil {
		apiErr := parseErrors(err)
		c.String(apiErr.Status, apiErr.Error())
		return
	}

	c.JSON(http.StatusOK, character)
}

func GetCharactersID(c *gin.Context) {
	id := c.Param("id")

	character, err := services.GetCharactersID(id)
	if err != nil {
		apiErr := parseErrors(err)
		c.String(apiErr.Status, apiErr.Error())
		return
	}

	c.JSON(http.StatusOK, character)
}

func GetEpisodes(c *gin.Context) {
	episodes, err := services.GetEpisodes()
	if err != nil {
		apiErr := parseErrors(err)
		c.String(apiErr.Status, apiErr.Error())
		return
	}
	c.JSON(http.StatusOK, episodes)
}

func GetEpisodesID(c *gin.Context) {
	p := c.Param("id")
	pi, err := strconv.Atoi(p)
	episodes, err := services.GetEpisodesID(pi)
	if err != nil {
		apiErr := parseErrors(err)
		c.String(apiErr.Status, apiErr.Error())
		return
	}
	c.JSON(http.StatusOK, episodes)
}

func GetQuotes(c *gin.Context) {
	quotes, err := services.GetQuotes()
	if err != nil {
		apiErr := parseErrors(err)
		c.String(apiErr.Status, apiErr.Error())
		return
	}
	c.JSON(http.StatusOK, quotes)
}

func GetQuoutesID(c *gin.Context) {
	p := c.Param("id")
	pi, err := strconv.Atoi(p)
	quotes, err := services.GetQuotesID(pi)
	if err != nil {
		apiErr := parseErrors(err)
		c.String(apiErr.Status, apiErr.Error())
		return
	}
	c.JSON(http.StatusOK, quotes)
}

func GetDeaths(c *gin.Context) {
	deaths, err := services.GetDeaths()
	if err != nil {
		apiErr := parseErrors(err)
		c.String(apiErr.Status, apiErr.Error())
		return
	}
	c.JSON(http.StatusOK, deaths)
}

func PostCharacters(c *gin.Context) {
	var character domain.CharactersBB
	err := c.BindJSON(&character)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	character, err = services.PostCharacters(character)
	if err != nil {
		apiErr := parseErrors(err)
		c.String(apiErr.Status, apiErr.Error())
		return
	}
	c.JSON(http.StatusOK, character)
}

func DeleteCharactersID(c *gin.Context) {
	id := c.Param("id")

	user, err := services.DeleteCharactersID(id)
	if err != nil {
		apiErr := parseErrors(err)
		c.String(apiErr.Status, apiErr.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func PutCharacters(c *gin.Context) {
	var character domain.CharactersBB

	err := c.BindJSON(&character)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	character, err = services.PutCharacters(character)
	if err != nil {
		apiErr := parseErrors(err)
		c.String(apiErr.Status, apiErr.Error())
		return
	}

	c.JSON(http.StatusOK, character)
}
