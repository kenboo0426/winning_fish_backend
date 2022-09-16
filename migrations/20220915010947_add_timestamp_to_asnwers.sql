-- +goose Up
alter table answers add created_at datetime;
alter table answers add updated_at datetime;

-- +goose Down
alter table answers drop column created_at;
alter table answers drop column updated_at;
