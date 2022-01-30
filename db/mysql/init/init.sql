-- SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
-- SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
-- SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';
SET CHARSET UTF8;
DROP SCHEMA IF EXISTS `app`;
CREATE SCHEMA IF NOT EXISTS `app` DEFAULT CHARACTER SET utf8;
USE `app`;

-- drop ----
DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `users_info`;
DROP TABLE IF EXISTS `articles`;
DROP TABLE IF EXISTS `articles_tag`;
DROP TABLE IF EXISTS `articles_comments`;
DROP TABLE IF EXISTS `articles_nice_status`;
DROP TABLE IF EXISTS `articles_image`;

-- create ----
CREATE TABLE IF NOT EXISTS `app`.`users` (
    `id` VARCHAR(32) NOT NULL COMMENT 'ユーザID',
    `name` VARCHAR(32) NOT NULL COMMENT 'ユーザ名',
    `age` int NOT NULL COMMENT '生年月日(年)',
    PRIMARY KEY (`id`),
    INDEX `idx_auth_token` (`id` ASC)
)
ENGINE = InnoDB
COMMENT = 'ユーザプロフィール';

CREATE TABLE IF NOT EXISTS `app`.`user_infos` (
    `id` VARCHAR(64) NOT NULL COMMENT 'ユーザID',
    `password` VARCHAR(64) NOT NULL COMMENT 'パスワード',
    `token` VARCHAR(64) NOT NULL COMMENT 'トークン',
    PRIMARY KEY (`id`)
)
ENGINE = InnoDB
COMMENT = 'ユーザー認証情報';

CREATE TABLE IF NOT EXISTS `app`.`articles` (
    `id` VARCHAR(64) NOT NULL COMMENT '記事識別ID',
    `title` VARCHAR(32) NOT NULL COMMENT '記事のタイトル',
    `description` VARCHAR(64) NOT NULL COMMENT '記事の内容',
    `nice` int NOT NULL COMMENT 'いいね数',
    `contents` text NOT NULL COMMENT '記事の内容',
    `user_id` VARCHAR(64) NOT NULL COMMENT 'ユーザーID',
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`)
    REFERENCES `app`.`user_infos` (`id`)
)
ENGINE = InnoDB
COMMENT = '記事の内容';

CREATE TABLE IF NOT EXISTS `app`.`article_images` (
    `tag_id` VARCHAR(64) NOT NULL COMMENT 'タグID',
    `article_id` VARCHAR(64) NOT NULL COMMENT '記事識別ID',
    `article_tag` VARCHAR(32) NOT NULL COMMENT '記事のタグ',
    PRIMARY KEY (`tag_id`)
)
ENGINE = InnoDB
COMMENT = '記事のタグ';

CREATE TABLE IF NOT EXISTS `app`.`article_tags` (
    `tag_id` VARCHAR(64) NOT NULL COMMENT 'タグID',
    `article_id` VARCHAR(64) NOT NULL COMMENT '記事識別ID',
    `article_tag` VARCHAR(32) NOT NULL COMMENT '記事のタグ',
    PRIMARY KEY (`tag_id`)
)
ENGINE = InnoDB
COMMENT = '記事のタグ';

CREATE TABLE IF NOT EXISTS `app`.`article_comments` (
    `article_id` VARCHAR(64) NOT NULL COMMENT '記事識別ID',
    `comments_id` VARCHAR(64) NOT NULL COMMENT 'コメントのID',
    `comments_contents` VARCHAR(64) NOT NULL COMMENT 'コメントの内容',
    `user_name` VARCHAR(64) NOT NULL COMMENT 'ユーザネーム',
    `user_image` VARCHAR(128) NOT NULL COMMENT 'ユーザーの画像',
    PRIMARY KEY (`comments_id`)
)
ENGINE = InnoDB
COMMENT = '記事へのコメント';

CREATE TABLE IF NOT EXISTS `app`.`article_nice_statuss` (
    `id` VARCHAR(64) NOT NULL COMMENT 'NiceID',
    `article_id` VARCHAR(64) NOT NULL COMMENT '記事識別ID',
    `user_id` VARCHAR(64) NOT NULL COMMENT 'ユーザID',
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`)
    REFERENCES `app`.`users` (`id`)
)
ENGINE = InnoDB
COMMENT = '記事にいいねした人';
