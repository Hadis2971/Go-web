CREATE TABLE IF NOT EXISTS Sales (
    `user_id` INT UNSIGNED NOT NULL,
    `product_id` INT UNSIGNED NOT NULL,
    FOREIGN KEY (`user_id`) REFERENCES `User` (`id`)
    ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (`product_id`) REFERENCES `Product` (`id`)
    ON DELETE CASCADE ON UPDATE CASCADE
)