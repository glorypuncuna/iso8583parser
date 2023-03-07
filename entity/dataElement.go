package entity

type Field struct {
	Id     int    `json:"id"`
	Label  string `json:"label"`
	Length int    `json:"length"`
	Value  string `json:"value"`
}
