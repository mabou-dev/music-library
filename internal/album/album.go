package album

type Album struct {
	Uuid   string `json:"uuid"`
	Artist string `json:"artist"`
	Name   string `json:"name"`
}

func New(artist string, name string) *Album {
	a := new(Album)
	a.Artist = artist
	a.Name = name
	return a
}
