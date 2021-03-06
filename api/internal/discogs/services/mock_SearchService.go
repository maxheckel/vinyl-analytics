// Code generated by mockery v1.0.0. DO NOT EDIT.

package services

import mock "github.com/stretchr/testify/mock"
import models "app/internal/discogs/models"

// MockSearchService is an autogenerated mock type for the SearchService type
type MockSearchService struct {
	mock.Mock
}

// Albums provides a mock function with given fields: query, pagination
func (_m *MockSearchService) Albums(query string, pagination *models.Pagination) ([]models.AlbumSearch, *models.Pagination, error) {
	ret := _m.Called(query, pagination)

	var r0 []models.AlbumSearch
	if rf, ok := ret.Get(0).(func(string, *models.Pagination) []models.AlbumSearch); ok {
		r0 = rf(query, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.AlbumSearch)
		}
	}

	var r1 *models.Pagination
	if rf, ok := ret.Get(1).(func(string, *models.Pagination) *models.Pagination); ok {
		r1 = rf(query, pagination)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*models.Pagination)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, *models.Pagination) error); ok {
		r2 = rf(query, pagination)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Artists provides a mock function with given fields: query, pagination
func (_m *MockSearchService) Artists(query string, pagination *models.Pagination) ([]models.ArtistSearch, *models.Pagination, error) {
	ret := _m.Called(query, pagination)

	var r0 []models.ArtistSearch
	if rf, ok := ret.Get(0).(func(string, *models.Pagination) []models.ArtistSearch); ok {
		r0 = rf(query, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ArtistSearch)
		}
	}

	var r1 *models.Pagination
	if rf, ok := ret.Get(1).(func(string, *models.Pagination) *models.Pagination); ok {
		r1 = rf(query, pagination)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*models.Pagination)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, *models.Pagination) error); ok {
		r2 = rf(query, pagination)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
