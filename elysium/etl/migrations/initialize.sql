CREATE TABLE `accounts`
(
    `id`              VARCHAR(255) NOT NULL PRIMARY KEY,
    `hashed_email`    VARCHAR(255),
    `encrypt_email`   VARBINARY(255),
    `hashed_password` VARCHAR(255),
    `created_at`      DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`      DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

CREATE TABLE `permissions`
(
    `id`         INT AUTO_INCREMENT PRIMARY KEY,
    `account_id` VARCHAR(255),
    `route`      VARCHAR(255)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

CREATE TABLE `posts`
(
    `id`         INT PRIMARY KEY AUTO_INCREMENT,
    `author`     VARCHAR(255),
    `content`    TEXT,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    INDEX `idx_author` (`author`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

CREATE TABLE `users`
(
    `id`           VARCHAR(255) NOT NULL PRIMARY KEY,
    `name`         VARCHAR(255),
    `alias`        VARCHAR(255),
    `avatar`       VARCHAR(255),
    `introduction` VARCHAR(255),
    `workplace`    VARCHAR(255),
    `hometown`     VARCHAR(255),
    `created_at`   DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`   DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

CREATE TABLE `interactions`
(
    `id`         INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `post_id`    INT UNSIGNED NOT NULL,
    `author`     VARCHAR(255) NOT NULL,
    `type`       VARCHAR(255) NOT NULL,
    `content`    TEXT         NOT NULL,
    `created_at` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    INDEX `post_id_idx` (`post_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;