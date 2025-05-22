package service

import (
	"example/web-service-gin/internal/models"
	"example/web-service-gin/internal/repository"
)

type AlbumService interface {
	GetAllAlbums() []models.Album
	GetAlbumById(id string) (*models.Album, error)
	CreateAlbum(album models.Album) error
	UpdateAlbum(id string, album models.Album) error
	DeleteAlbum(id string) error
}

type albumService struct {
	repo repository.AlbumRepository
}

func NewAlbumService(repo repository.AlbumRepository) AlbumService {
	return &albumService{repo: repo}
}

func(s *albumService) GetAllAlbums() []models.Album {
	return s.repo.GetAll()
}

func (s *albumService) GetAlbumById(id string) (*models.Album, error) {
	return s.repo.GetById(id)
}

func (s *albumService) CreateAlbum(album models.Album) error {
	return s.repo.Create(album)
}

func (s *albumService) UpdateAlbum(id string, album models.Album) error {
	return s.repo.Update(id, album)
}

func (s *albumService) DeleteAlbum(id string) error {
	return s.repo.Delete(id)
}