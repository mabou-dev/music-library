package album

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mabou-dev/music-library/pkg/log"
)

type AlbumHandler struct {
	service *AlbumService
}

func NewAlbumHandler(service *AlbumService) *AlbumHandler {
	handler := new(AlbumHandler)
	handler.service = service
	return handler
}

func (h *AlbumHandler) GetAlbums(c *gin.Context) {
	albums := h.service.GetAlbums()
	log.Debugf("retrieves albums are %+v", albums)
	c.JSON(http.StatusOK, albums)
}
