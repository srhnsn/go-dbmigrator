CREATE TABLE `user_to_role` (
    `user_id` INT(10) UNSIGNED NOT NULL,
    `role_id` INT(10) UNSIGNED NOT NULL,
    INDEX `FK_user_to_role_user` (`user_id`),
    INDEX `FK_user_to_role_role` (`role_id`),
    CONSTRAINT `FK_user_to_role_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT `FK_user_to_role_role` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
)
ENGINE=InnoDB;
