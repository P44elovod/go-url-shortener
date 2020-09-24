package url

import (
	_httpURLDelivery "go-url-shortener/url/delivery/http"
	_urlRepository "go-url-shortener/url/repository/sql"
	_urlUsecase "go-url-shortener/url/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InitURL(db *gorm.DB, srv *gin.Engine) {
	urlRepo := _urlRepository.NewURLRepository(db)
	_httpURLDelivery.NewURLHandler(srv)
	_urlUsecase.NewURLUSeCase(urlRepo)
}
