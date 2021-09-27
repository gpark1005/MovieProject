package entities

import "github.com/google/uuid"

type Movie struct {
	Id       string   `json:"Id"`
	Title    string   `json:"Title"`
	Genre    []string `json:"Genre"`
	Director string   `json:"Director"`
	Rating   float32  `json:"Rating"`
}

func (m *Movie) GetId() {
	m.Id = uuid.New().String()
}
