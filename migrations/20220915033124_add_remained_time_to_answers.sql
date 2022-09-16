-- +goose Up
alter table answers add remained_time float;

-- +goose Down
alter table answers drop column remained_time;
