package entities

type Movie struct {
	Id          string
	Title       string
	Genre       []string
	Description string
	Actors      []string
	Rating      float32
	Director    string
}
