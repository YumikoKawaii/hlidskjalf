CREATE TABLE `accounts` (
    `id` VARCHAR(255) NOT NULL PRIMARY KEY,
    `hashed_email` VARCHAR(255),
    `encrypt_email` VARBINARY(255),
    `hashed_password` VARCHAR(255),
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `permissions` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `account_id` VARCHAR(255),
    `route` VARCHAR(255)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
