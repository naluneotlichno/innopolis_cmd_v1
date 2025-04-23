package postgres

import (
	"database/sql"
	"log"

	"github.com/szaluzhanskaya/Innopolis/chain-service/internal/entity"
	"github.com/szaluzhanskaya/Innopolis/chain-service/internal/repo"
)

const createMessageChainQuery = `
INSERT INTO message_chains (user_id, created_at, updated_at, status, title)
VALUES ($1, $2, $3, $4, $5)
RETURNING uuid`

type MessageChainPostgres struct {
	DB *sql.DB
}

func New(db *sql.DB) repo.MessageChainRepository {
	return &MessageChainPostgres{DB: db}
}

func (r *MessageChainPostgres) CreateMessageChain(chain *entity.MessageChain) error {

	tx, err := r.DB.Begin()
	if err != nil {
		log.Printf("Failed to begin transaction: %v", err)
		return err
	}

	defer func() {
		if err := recover(); err != nil {
			if err = tx.Rollback(); err != nil {
				log.Printf("Transaction rollback failed: %v", err)
			}
			log.Printf("Transaction rolled back because panic: %v", err)
		}
	}()

	err = tx.QueryRow(
		createMessageChainQuery,
		chain.UserID,
		chain.CreatedAt,
		chain.UpdatedAt,
		chain.Status,
		chain.Title,
	).Scan(&chain.UUID)

	if err != nil {
		if err = tx.Rollback(); err != nil {
			log.Printf("Transaction rollback failed: %v", err)
		}
		log.Printf("Failed to create message chain: %v", err)
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return err
	}

	return nil
}
