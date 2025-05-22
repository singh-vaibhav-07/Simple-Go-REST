package repository

import (
	"errors"
	"example/web-service-gin/internal/models"
)

type AlbumRepository interface {
	GetAll() []models.Album
	GetById(id string) (*models.Album, error)
	Create(album models.Album) error
	Update(id string, album models.Album) error
	Delete(id string) error
}

type albumRepository struct {
	albums []models.Album
}

func NewAlbumRepository() AlbumRepository {
	return &albumRepository{
		albums: []models.Album{
			{Id: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
			{Id: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
			{Id: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
			{Id: "4", Title: "Angels for each other", Artist: "Arijit Singh", Price: 69.99},
		},
	}
}

func (r *albumRepository) GetAll() []models.Album {
	return r.albums
}

func (r *albumRepository) GetById(id string) (*models.Album, error) {
	for _, a := range r.albums {
		if a.Id == id {
			return &a, nil
		}
	}
	return nil, errors.New("album not found")
}

func (r *albumRepository) Create(album models.Album) error {
	r.albums = append(r.albums, album)
	return nil
}

func (r *albumRepository) Update(id string, album models.Album) error {
	for i, a := range r.albums {
		if a.Id == id {
			album.Id = id
			r.albums[i] = album
			return nil
		}
	}
	return errors.New("album not found")
}

func (r *albumRepository) Delete(id string) error {
	for i, a := range r.albums {
		if a.Id == id {
			r.albums = append(r.albums[:i], r.albums[i+1:]...)
			return nil
		}
	}
	return errors.New("album not found")
}
