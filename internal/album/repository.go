package album

var albums = []Album{
	{
		Uuid:   "a",
		Name:   "album_A",
		Artist: "artist_A",
	},
	{
		Uuid:   "b",
		Name:   "album_B",
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
