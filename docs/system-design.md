### 💥 Техническое задание на проект (Go)

**Цель:**  
Создать систему управления цепочками сообщений (Message Chains), состоящих из блоков с мультимедийным контентом. Предусмотреть авторизацию и безопасный доступ к админке.

---

### 🔐 Авторизация и безопасность

- Пользователь логинится по паре `логин + пароль`.
- Пароли храним безопасно (bcrypt или аналогичный способ, с солью).
- Доступ в административную панель — **только авторизованным**.
- Роли пользователей: `admin`, `user`.

---

### 🧩 Работа с цепочками сообщений

- Просмотр списка всех цепочек.
- Фильтрация по дате, статусу (`active | archived`) и другим полям.
- Создание новой цепочки.
- Редактирование существующей (в том числе — изменение порядка блоков внутри цепочки).

---

### 🔗 Работа с блоками сообщений

- Блок может содержать: текст, аудио, видео, картинку.
- Возможность создать, редактировать, удалить блок.
- Добавление блока в цепочку.
- Удаление блока из цепочки (без удаления из базы).
- Полное удаление блока из базы (если он не используется).

---

### 🎛️ Отображение блоков

- Возможность настраивать порядок блоков внутри цепочки.
- Параметры отображения (размер, выравнивание, и т.п.) задаются через JSON.

---

### ⚙️ Технические детали

- Язык: Go
- Архитектура: чистая, модульная
- Библиотеки: на выбор (Gin, Echo, sqlx, GORM и т.д.)
- Кроссплатформенность (можно запускать на любой ОС)

---

### 📦 Ожидаемый результат

- Полностью рабочая админка с авторизацией.
- Возможность управлять цепочками и блоками.
- Надёжная работа с контентом и отображением.
- Безопасность доступа и хранения данных.

---

### 🧱 Таблица данных (SQL схема)
```sql
CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    login VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    salt VARCHAR(255) NOT NULL,
    role ENUM('admin', 'user') NOT NULL
);

CREATE TABLE message_chains (
    id INT PRIMARY KEY AUTO_INCREMENT,
    creation_date DATE NOT NULL,
    status ENUM('active', 'archived') NOT NULL,
    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE message_blocks (
    id INT PRIMARY KEY AUTO_INCREMENT,
    content_type ENUM('text', 'audio', 'video', 'image') NOT NULL,
    content TEXT NOT NULL,
    display_parameters JSON
);

CREATE TABLE chain_block_links (
    id INT PRIMARY KEY AUTO_INCREMENT,
    chain_id INT NOT NULL,
    block_id INT NOT NULL,
    FOREIGN KEY (chain_id) REFERENCES message_chains(id),
    FOREIGN KEY (block_id) REFERENCES message_blocks(id)
);
```
