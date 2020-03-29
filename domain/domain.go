package domain
//DOMAIN
type CharactersBB struct {
	CharID                   int           `json:"char_id"`
	Name                     string        `json:"name"`
	Birthday                 string        `json:"birthday"`
	Occupation               []string      `json:"occupation"`
	Img                      string        `json:"img"`
	Status                   string        `json:"status"`
	Nickname                 string        `json:"nickname"`
	Appearance               []int         `json:"appearance"`
	Portrayed                string        `json:"portrayed"`
	Category                 string        `json:"category"`
	BetterCallSaulAppearance []interface{} `json:"better_call_saul_appearance"`
}

func (c CharactersBB) ValidCharID() bool {
	return c.CharID > 0
}

func (c CharactersBB) HasName() bool {
	return c.Name != ""
}

func (c CharactersBB) HasBirthday() bool {
	return c.Birthday != ""
}

func (c CharactersBB) HasIMG() bool {
	return c.Img != ""
}

func (c CharactersBB) HasStatus() bool {
	return c.Status != ""
}
func (c CharactersBB) HasNickname() bool {
	return c.Nickname != ""
}
func (c CharactersBB) HasPortrayed() bool {
	return c.Portrayed != ""
}
func (c CharactersBB) HasCategory() bool {
	return c.Category != ""
}

/*This a struct of episodes of Breaking Bad*/
type EpisodesBB []struct {
	EpisodeID  int      `json:"episode_id"`
	Title      string   `json:"title"`
	Season     string   `json:"season"`
	AirDate    string   `json:"air_date"`
	Characters []string `json:"characters"`
	Episode    string   `json:"episode"`
	Series     string   `json:"series"`
}

/*This a struct of Quotes of Breaking Bad*/
type QuotesBB []struct {
	QuoteID int    `json:"quote_id"`
	Quote   string `json:"quote"`
	Author  string `json:"author"`
	Series  string `json:"series"`
}

type DeathsBB []struct {
	DeathID        int    `json:"death_id"`
	Death          string `json:"death"`
	Cause          string `json:"cause"`
	Responsible    string `json:"responsible"`
	LastWords      string `json:"last_words"`
	Season         int    `json:"season"`
	Episode        int    `json:"episode"`
	NumberOfDeaths int    `json:"number_of_deaths"`
}
