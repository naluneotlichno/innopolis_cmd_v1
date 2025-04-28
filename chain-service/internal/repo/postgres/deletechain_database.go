package postgres

import (
	"context"
	"log"
)

func (db *MessageChainPostgres) DeleteMessageChain(uuid string) error {
	_, err := db.DB.Exec(context.Background(), "DELETE FROM chain_block_list WHERE uuid = $1", uuid)
	if err != nil {
		log.Printf("Failed to delete from database:%v", err)
		return err
	}
	return nil
}
