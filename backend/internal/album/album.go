package album

type Album struct {
	Id     string `json:"id"`
	Artist string `json:"artist"`
	Title  string `json:"title"`
}

func New(artist string, title string) *Album {
	a := new(Album)
	a.Artist = artist
	a.Title = title
	return a
}
