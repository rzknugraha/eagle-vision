package services

import (
	"github.com/rzknugraha/eagle-vision/models"
	"github.com/rzknugraha/eagle-vision/repositories"
)

// IPlayerService is
type IPlayerService interface {
	StorePlayer(models.Player) (models.Player, error)
}

// PlayerService is
type PlayerService struct {
	PlayerRepository repositories.IPlayerRepository
}

// StorePlayer is
func (p *PlayerService) StorePlayer(data models.Player) (result models.Player, err error) {
	result, err = p.PlayerRepository.StorePlayer(data)
	return result, err
}
