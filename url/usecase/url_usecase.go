package usecase

import "go-url-shortener/domain"

type urlUsecase struct {
	urlRepo domain.URLRepository
}

func NewURLUSeCase(r domain.URLRepository) domain.URLUseCase {
	return &urlUsecase{
		urlRepo: r,
	}
}

func (u *urlUsecase) Create(url *domain.URL) map[string]interface{} {
	return map[string]interface{}{}
}
func (u *urlUsecase) GetByHash(id string) map[string]interface{} {
	return map[string]interface{}{}

}
