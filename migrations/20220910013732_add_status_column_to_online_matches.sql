-- +goose Up
alter table online_matches add status string;

-- +goose Down
alter table online_matches drop column status;