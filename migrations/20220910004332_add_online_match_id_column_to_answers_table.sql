-- +goose Up
alter table answers add online_match_id integer;

-- +goose Down
alter table answers drop column online_match_id;
