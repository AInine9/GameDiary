CREATE TABLE `users`
(
    `id`         int UNIQUE   NOT NULL,
    `user_name`  varchar(255) NOT NULL,
    `created_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

CREATE TABLE `games`
(
    `id`         int(20) NOT NULL AUTO_INCREMENT,
    `name`       varchar(255) NOT NULL,
    `created_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

CREATE TABLE `diaries`
(
    `id`               int      NOT NULL AUTO_INCREMENT,
    `user_id`          int      NOT NULL,
    `game_id`          int(20) NOT NULL,
    `start_playing_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `end_playing_at`   datetime,
    `created_at`       datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`       datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES users (`id`),
    FOREIGN KEY (`game_id`) REFERENCES games (`id`)
);
