CREATE TABLE users ( -- Создаём таблицу с пользователями
    id INT PRIMARY KEY AUTO_INCREMENT, -- ID пользователя
    username VARCHAR(20) NOT NULL,
    first_name VARCHAR(20) NOT NULL, -- Имя
    last_name VARCHAR(20) NOT NULL, -- И фамилия, возможно так будет удобнее
    created_at DATETIME NOT NULL DEFAULT NOW() -- время создания пользователя
);

CREATE EXTENSION IF NOT EXISTS "uuid-ossp"; -- Включаем расширение для работы с UUID, не поддерживается на сервере default
CREATE TYPE chain_status AS ENUM ('created', 'archived'); -- создаём перечесление
CREATE TABLE message_chains (
id INT PRIMARY KEY AUTO_INCREMENT, -- Уникальный идентификатор цепочки сообщений
uuid UUID DEFAULT uuid_generate_v4() UNIQUE, -- Уникальный идентификатор в формате UUID
user_id INT NOT NULL, -- Идентификатор пользователя создавшего цепочку
created_at DATETIME NOT NULL NOW(), -- Дата и время создания цепочки
updated_at DATETIME NOT NULL, -- Дата и время обновления цепочки
status chain_status NOT NULL, -- Статус цепочки (например, active, archived)
title VARCHAR(255), -- Заголовок цепочки
FOREIGN KEY (user_id) REFERENCES users(id)
);