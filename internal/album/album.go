package album

type Album struct {
	artist string
	name   string
	genres []string
}

func New(artist string, name string) *Album {
	a := new(Album)
	a.artist = artist
	a.name = name
	return a
}

func (a *Album) GetArtist() string {
	return a.artist
}

func (a *Album) GetName() string {
	return a.name
}

func (a *Album) SetGenre(g string) {
	a.genres = append(a.genres, g)
}
