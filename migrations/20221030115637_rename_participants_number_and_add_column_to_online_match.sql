-- +goose Up
alter table online_matches drop column person_number;
alter table online_matches rename column participants_number to max_participate_number;
alter table online_matches add question_number integer not null default 5;
alter table online_matches add with_bot boolean not null default false;
alter table online_matches add room_id string;

-- +goose Down
alter table online_matches add person_number integer;
alter table online_matches rename column max_participate_number to participants_number;
alter table online_matches drop column question_number;
alter table online_matches drop column with_bot;
alter table online_matches drop column room_id;
