package sql

import (
	"go-url-shortener/domain"

	"github.com/jinzhu/gorm"
)

type sqlURLRepository struct {
	DB *gorm.DB
}

func NewURLRepository(db *gorm.DB) domain.URLRepository {
	return &sqlURLRepository{
		DB: db,
	}
}

func (r *sqlURLRepository) Store(url *domain.URL) (err error)
func (r *sqlURLRepository) GetAll() (urlList []domain.URL)
func (r *sqlURLRepository) GetByHash(hash string) (url domain.URL)
func (r *sqlURLRepository) GetByID(id string) (url domain.URL)
