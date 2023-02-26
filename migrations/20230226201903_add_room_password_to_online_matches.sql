-- +goose Up
alter table online_matches add room_password string;

-- +goose Down
alter table online_matches drop column room_password;
