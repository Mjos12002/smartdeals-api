package model

// Address represents the address model in the database
type Address struct {
	ID          int    `json:"id"`
	Street      string `json:"street"`
	PopularName string `json:"popularname"`
	Province    string `json:"province"`
	District    string `json:"district"`
	Sector      string `json:"sector"`
	LongLat     string `json:"longlat"`
}
