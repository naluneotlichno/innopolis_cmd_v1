package postgres

import (
	"context"
	"log"

	"github.com/szaluzhanskaya/Innopolis/chain-service/internal/entity"
)

const createMessageChainQuery = `
INSERT INTO message_chains (user_id, created_at, updated_at, status, title)
VALUES ($1, $2, $3, $4, $5)
RETURNING uuid`

func (r *MessageChainPostgres) CreateMessageChain(chain *entity.MessageChain) error {

	tx, err := r.DB.Begin(context.Background())
	if err != nil {
		log.Printf("Failed to begin transaction: %v", err)
		return err
	}

	defer func() {
		if err := recover(); err != nil {
			if err = tx.Rollback(context.Background()); err != nil {
				log.Printf("Transaction rollback failed: %v", err)
			}
			log.Printf("Transaction rolled back because panic: %v", err)
		}
	}()

	err = tx.QueryRow(
		context.Background(),
		createMessageChainQuery,
		chain.UserID,
		chain.CreatedAt,
		chain.UpdatedAt,
		chain.Status,
		chain.Title,
	).Scan(&chain.UUID)

	if err != nil {
		if err = tx.Rollback(context.Background()); err != nil {
			log.Printf("Transaction rollback failed: %v", err)
		}
		log.Printf("Failed to create message chain: %v", err)
		return err
	}

	if err := tx.Commit(context.Background()); err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return err
	}

	return nil
}
