CREATE TABLE message_blocks (
    id INT PRIMARY KEY AUTO_INCREMENT,
    title TEXT NOT NULL,
    content_type ENUM('text', 'audio', 'video', 'image') NOT NULL,
    content TEXT NOT NULL,
    display_parameters JSON,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
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