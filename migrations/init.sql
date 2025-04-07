-- +migrate Up
CREATE TABLE message_blocks (
    id INT PRIMARY KEY AUTO_INCREMENT,
    -- title TEXT NOT NULL, -- Можно добавить короткое название (для удобства в админке)
    content_type ENUM('text', 'audio', 'video', 'image') NOT NULL,
    content TEXT NOT NULL,
    display_parameters JSON,
    -- created_at TIMESTAMP DEFAULT now(), -- Можно добавить дату создания (тоже удобно в админке)
    -- updated_at TIMESTAMP -- Можно добавить дату изменения (это тоже больше для админов)
);

CREATE TABLE chain_block_links (
    id INT PRIMARY KEY AUTO_INCREMENT,
    chain_id INT NOT NULL,
    block_id INT NOT NULL,
    FOREIGN KEY (chain_id) REFERENCES message_chains(id),
    FOREIGN KEY (block_id) REFERENCES message_blocks(id)
);

-- Прямо добавляем индексы для ускорения работы
CREATE INDEX idx_chain_block_links_chain_id ON chain_block_links(chain_id);
CREATE INDEX idx_chain_block_links_block_id ON chain_block_links(block_id);

-- +migrate Down
DROP INDEX IF EXISTS idx_chain_block_links_block_id;
DROP INDEX IF EXISTS idx_chain_block_links_chain_id;

DROP TABLE IF EXISTS chain_block_links;
DROP TABLE IF EXISTS message_blocks;