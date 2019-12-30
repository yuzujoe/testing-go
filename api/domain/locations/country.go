package locations

type Country struct {
	ID            string        `json:"id"`
	Name          string        `json:"name"`
	TimeZone      string        `json:"time_zone"`
	GeoInfomation GeoInfomation `json:"geo_infomation"`
	States        []State       `json:"states"`
}

type GeoInfomation struct {
	Location GeoLocation `json:"location"`
}

type GeoLocation struct {
	Latitude  float64 `json:"latitube"`
	Longitude float64 `json:"longitube"`
}

type State struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
