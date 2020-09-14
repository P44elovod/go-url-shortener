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

func (r *sqlURLRepository) Store(url *domain.URL) (id uint, err error) {
	res := r.DB.Create(&url)

	err = res.Error
	id = url.ID

	return
}
func (r *sqlURLRepository) GetAll() (urlList []domain.URL, err error) {
	if r.DB.Table("url").Find(&urlList).RecordNotFound() {
		err = gorm.ErrRecordNotFound
	}

	return
}

func (r *sqlURLRepository) GetByHash(hash string) (url domain.URL, err error) {
	if r.DB.Table("url").Where("hash = ?", hash).First(&url).RecordNotFound() {
		err = gorm.ErrRecordNotFound
	}

	return
}
func (r *sqlURLRepository) GetByID(id string) (url domain.URL, err error) {
	if r.DB.Table("url").Where("id = ?", id).First(&url).RecordNotFound() {
		err = gorm.ErrRecordNotFound
	}

	return
}
