create database webook;

CREATE TABLE `users` (
                         `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
                         `email` VARCHAR(255) NOT NULL UNIQUE,
                         `password` VARCHAR(255) NOT NULL,
                         `ctime` BIGINT(20) NOT NULL,
                         `utime` BIGINT(20) NOT NULL,
                         PRIMARY KEY (`id`)
);