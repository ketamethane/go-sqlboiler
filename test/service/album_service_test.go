package service

import (
	"testing"
	"go-sqlboiler/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAlbumService is a mock of AlbumService for testing purposes.
type MockAlbumService struct {
	mock.Mock
}

// GetAllAlbums mocks the GetAllAlbums method.
func (m *MockAlbumService) GetAllAlbums() []models.Album {
	args := m.Called()
	return args.Get(0).([]models.Album)
}

// GetAlbumByID mocks the GetAlbumByID method.
func (m *MockAlbumService) GetAlbumByID(id string) (models.Album, error) {
	args := m.Called(id)
	return args.Get(0).(models.Album), args.Error(1)
}

// AddNewAlbum mocks the AddNewAlbum method.
func (m *MockAlbumService) AddNewAlbum(album models.Album) {
	m.Called(album)
}

// UpdateAlbumByID mocks the UpdateAlbumByID method.
func (m *MockAlbumService) UpdateAlbumByID(id string, updatedAlbum models.Album) error {
	args := m.Called(id, updatedAlbum)
	return args.Error(0)
}

// DeleteAlbumByID mocks the DeleteAlbumByID method.
func (m *MockAlbumService) DeleteAlbumByID(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// TestAlbumService tests the AlbumService methods.
func TestAlbumService(t *testing.T) {
	mockAlbumService := new(MockAlbumService)

	// Test GetAllAlbums
	mockAlbumService.On("GetAllAlbums").Return([]models.Album{{ID: "1", Title: "Test Album"}})
	assert.Equal(t, []models.Album{{ID: "1", Title: "Test Album"}}, mockAlbumService.GetAllAlbums())

	// Test GetAlbumByID
	mockAlbumService.On("GetAlbumByID", "1").Return(models.Album{ID: "1", Title: "Test Album"}, nil)
	_, err := mockAlbumService.GetAlbumByID("1")
	assert.NoError(t, err)

	// Test AddNewAlbum
	mockAlbumService.On("AddNewAlbum", models.Album{ID: "2", Title: "Another Test Album"}).Return()
	mockAlbumService.AddNewAlbum(models.Album{ID: "2", Title: "Another Test Album"})

	// Test UpdateAlbumByID
	mockAlbumService.On("UpdateAlbumByID", "1", models.Album{ID: "1", Title: "Updated Test Album"}).Return(nil)
	err = mockAlbumService.UpdateAlbumByID("1", models.Album{ID: "1", Title: "Updated Test Album"})
	assert.NoError(t, err)

	// Test DeleteAlbumByID
	mockAlbumService.On("DeleteAlbumByID", "1").Return(nil)
	err = mockAlbumService.DeleteAlbumByID("1")
	assert.NoError(t, err)
}
