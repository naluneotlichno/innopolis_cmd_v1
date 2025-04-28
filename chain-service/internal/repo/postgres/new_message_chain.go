package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/szaluzhanskaya/Innopolis/chain-service/internal/repo"
)

type MessageChainPostgres struct {
	DB *pgxpool.Pool
}

func New(db *pgxpool.Pool) repo.MessageChainRepository {
	return &MessageChainPostgres{DB: db}
}
