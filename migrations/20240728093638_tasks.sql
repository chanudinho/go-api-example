-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `tasks`(
    `id`              INT(11)         NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `user_id`         INT(11)         NOT NULL,
    `title`           VARCHAR(255)    NOT NULL,
    `summary`         VARCHAR(2500)   NOT NULL,
    `completed_at`    TIMESTAMP       NULL,
    `created_at`      TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`      TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`      TIMESTAMP       NULL DEFAULT NULL,
    CONSTRAINT `fk_tasks_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `tasks`;
-- +goose StatementEnd
