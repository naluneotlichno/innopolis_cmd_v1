package http

import "github.com/szaluzhanskaya/Innopolis/chain-service/internal/usecase"

type MessageChainHandler struct {
	service usecase.MessageChainUsecase
}

func New(s usecase.MessageChainUsecase) *MessageChainHandler {
	return &MessageChainHandler{service: s}
}
