-- +goose Up
alter table online_matches add remaining_wait_time float;

-- +goose Down
alter table online_matches drop column remaining_wait_time;
