-- +goose Up
alter table users add icon string;

-- +goose Down
alter table users drop column icon;
