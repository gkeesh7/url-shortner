CREATE DATABASE url_shortner;

USE url_shortner;

DROP TABLE IF EXISTS `redirection`;

CREATE TABLE `redirection`
(
    `id`         INT(11)       NOT NULL AUTO_INCREMENT,
    `url_id`     varchar(20)   NOT NULL UNIQUE,
    `url`        varchar(1024) NOT NULL,
    `created_at` TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `expiry`     TIMESTAMP     NULL     DEFAULT NULL,
    PRIMARY KEY (`id`)
);


DROP TABLE IF EXISTS `url_stats`;

CREATE TABLE `url_stats`
(
    `id`         INT(11)     NOT NULL AUTO_INCREMENT,
    `url_id`     varchar(20) NOT NULL,
    `count`      INT(11)     NOT NULL,
    `created_at` TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

INSERT INTO `redirection` (`id`, `url_id`, `url`, `created_at`, `updated_at`, `expiry`) VALUES ('1', 'test', 'https://youtu.be/dQw4w9WgXcQ?t=43', '2020-01-08 14:05:37', '2020-01-08 14:05:37', '2030-01-19 14:05:37');