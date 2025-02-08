package album

type AlbumService struct {
	repo *AlbumRepository
}

func NewAlbumService(repo *AlbumRepository) *AlbumService {
	service := new(AlbumService)
	service.repo = repo
	return service
}

func (s *AlbumService) GetAlbums() []Album {
	return s.repo.GetAll()
}
