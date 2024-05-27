package service

import (
	"go-sqlboiler/internal/models"
)

type AlbumService struct {
	albums []models.Album
}

func NewAlbumService() *AlbumService {
	sampleAlbums := []models.Album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}

	return &AlbumService{
		albums: sampleAlbums,
	}
}

func (as *AlbumService) GetAllAlbums() []models.Album {
	return as.albums
}

func (as *AlbumService) GetAlbumByID(id string) (models.Album, error) {
	for _, album := range as.albums {
		if album.ID == id {
			return album, nil
		}
	}
	return models.Album{}, nil // Not found
}

func (as *AlbumService) AddNewAlbum(album models.Album) {
	as.albums = append(as.albums, album)
}

func (as *AlbumService) UpdateAlbumByID(id string, updatedAlbum models.Album) error {
	for i, album := range as.albums {
		if album.ID == id {
			as.albums[i] = updatedAlbum
			return nil
		}
	}
	return nil // Not found
}

func (as *AlbumService) DeleteAlbumByID(id string) error {
	for i, album := range as.albums {
		if album.ID == id {
			as.albums = append(as.albums[:i], as.albums[i+1:]...)
			return nil
		}
	}
	return nil // Not found
}