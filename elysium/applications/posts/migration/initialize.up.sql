CREATE TABLE `posts`
(
    `id`         INT PRIMARY KEY AUTO INCREMENT,
    `author`     VARCHAR(255),
    `content`    TEXT,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;