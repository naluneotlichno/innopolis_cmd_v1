-- 🔥 Таблица для хранения информации о загруженных медиа-файлах 🔥

-- Расширение для генерации UUID v4, брат. 
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Создане таблицы для схрона мультимедиа файлов в MinIo/S3
CREATE TABLE IF NOT EXISTS storage (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),   -- Автогенерация уникального UUID
    s3_path TEXT NOT NULL,                              -- Путь к медиа-файлу в S3
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP      -- Время создания медиа-файла
);