-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `users`(
    `id`              INT(11)         NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name`            VARCHAR(255)    NOT NULL,
    `email`           VARCHAR(255)    NOT NULL,
    `password`        VARCHAR(255)    NOT NULL,
    `role`            VARCHAR(255)    NOT NULL,
    `created_at`      TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`      TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`      TIMESTAMP       NULL DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `users`;
-- +goose StatementEnd
