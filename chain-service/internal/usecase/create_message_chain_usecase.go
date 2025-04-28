package usecase

import (
	"time"

	"github.com/szaluzhanskaya/Innopolis/chain-service/internal/entity"
	"github.com/szaluzhanskaya/Innopolis/chain-service/internal/repo"
)

type MessageChainUsecase interface {
	CreateMessageChain(userID int, title string) (*entity.MessageChain, error)
	DeleteMessageChain(uuid string) error
}

type messageChainUsecase struct {
	repo repo.MessageChainRepository
}

func New(repo repo.MessageChainRepository) MessageChainUsecase {
	return &messageChainUsecase{repo: repo}
}

func (uc *messageChainUsecase) CreateMessageChain(userID int, title string) (*entity.MessageChain, error) {
	now := time.Now()
	chain := &entity.MessageChain{
		UserID:    userID,
		CreatedAt: now,
		UpdatedAt: now,
		Status:    entity.Created,
		Title:     title,
	}

	err := uc.repo.CreateMessageChain(chain)
	if err != nil {
		return nil, err
	}

	return chain, nil
}

func (uc *messageChainUsecase) DeleteMessageChain(uuid string) error {
	return uc.repo.DeleteMessageChain(uuid)
}
