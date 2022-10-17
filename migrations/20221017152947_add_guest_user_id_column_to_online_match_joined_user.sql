-- +goose Up
alter table online_match_joined_users add guest_user_id string;

-- +goose Down
alter table online_match_joined_users drop column guest_user_id;