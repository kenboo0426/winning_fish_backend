-- +goose Up
alter table online_match_joined_users add online_match_id integer;

-- +goose Down
alter table online_match_joined_users drop column online_match_id;
