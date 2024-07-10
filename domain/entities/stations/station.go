package station_entity

type Station struct {
	Id              int      `json:"id" db:"station_id"`
	Municipality_id int      `json:"municipality_id"`
	Name            string   `json:"name"`
	Coordinates     string   `json:"coordinates"`
	Description     string   `json:"description"`
	Code            string   `json:"code"`
	Abbreviation    string   `json:"abbreviation"`
	Phones          []Phones `json:"phones"`
	State_id        int      `json:"state_id"`
	Parish_id       int      `json:"parish_id"`
	Sector          string   `json:"sector"`
	Community       string   `json:"community"`
	Street          string   `json:"street"`
	Address         string   `json:"address"`
}

type StationDto struct {
	Id              int      `json:"id" db:"station_id"`
	Municipality_id int      `json:"municipality_id"`
	Name            string   `json:"name"`
	Coordinates     string   `json:"coordinates"`
	Description     string   `json:"description"`
	Code            string   `json:"code"`
	Abbreviation    string   `json:"abbreviation"`
	Phones          []Phones `json:"phones"`
	State_id        int      `json:"state_id"`
	Parish_id       int      `json:"parish_id"`
	Sector          string   `json:"sector"`
	Community       string   `json:"community"`
	Street          string   `json:"street"`
	Address         string   `json:"address"`
}

type Phones struct {
	AreaCode string `json:"area_code"`
	Number   string `json:"number"`
}