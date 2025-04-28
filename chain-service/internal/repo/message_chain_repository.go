package repo

import (
	"github.com/szaluzhanskaya/Innopolis/chain-service/internal/entity"
)

type MessageChainRepository interface {
	CreateMessageChain(chain *entity.MessageChain) error
	DeleteMessageChain(uuid string) error
}
