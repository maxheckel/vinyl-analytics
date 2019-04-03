package services

import (
	"app/internal/discogs/models"
	"encoding/json"
	"fmt"
	"gopkg.in/resty.v1"
)

type MastersService interface {
	GetMaster(id uint) (*models.Master, error)
}

type masterService struct {
	client *resty.Client
}

func NewMasterService(client *resty.Client) MastersService {
	return &masterService{client: client}
}

func (m *masterService) GetMaster(id uint) (*models.Master, error) {
	results, err := m.client.R().Get(fmt.Sprintf("https://api.discogs.com/masters/%d", id))
	if err != nil{
		 return nil, err
	}
	master := &models.Master{}
	err = json.Unmarshal(results.Body(), master)
	if err != nil {
		return nil, err
	}
	return master, nil
}

