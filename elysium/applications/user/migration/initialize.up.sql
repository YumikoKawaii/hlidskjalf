CREATE TABLE `users`
(
    `id`              VARCHAR(255) NOT NULL PRIMARY KEY,
    `name`    VARCHAR(255),
    `alias`   VARCHAR(255),
    `avatar` VARCHAR(255),
    `introduction` VARCHAR(255),
    `workplace` VARCHAR(255),
    `hometown` VARCHAR(255),
    `created_at`      DATETIME     NOT NULL,
    `updated_at`      DATETIME     NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;