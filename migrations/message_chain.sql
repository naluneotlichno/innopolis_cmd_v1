CREATE EXTENSION IF NOT EXISTS "uuid-ossp"; -- Включаем расширение для работы с UUID, не поддерживается на сервере default
CREATE TYPE chain_status AS ENUM ('created', 'archived'); -- создаём перечесление
CREATE TABLE message_chains (
id INT PRIMARY KEY, -- Уникальный идентификатор цепочки сообщений
uuid UUID DEFAULT uuid_generate_v4() UNIQUE, -- Уникальный идентификатор в формате UUID
user_id INT NOT NULL, -- Идентификатор пользователя создавшего цепочку
creation_date DATE NOT NULL NOW(), -- Дата и время создания цепочки
status chain_status NOT NULL, -- Статус цепочки (например, active, archived)
title VARCHAR(255), -- Заголовок цепочки
FOREIGN KEY (user_id) REFERENCES users(id)
);