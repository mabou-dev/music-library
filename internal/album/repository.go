package album

var albums = []Album{
	{
		Id:   "a",
		Title:   "album_A",
		Artist: "artist_A",
	},
	{
		Id:   "b",
		Title:   "album_B",
		Artist: "artist_B",
	},
}

type AlbumRepository struct{}

func NewAlbumRepository() *AlbumRepository {
	repository := new(AlbumRepository)
	return repository
}

func (r *AlbumRepository) GetAll() []Album {
	return albums
}
