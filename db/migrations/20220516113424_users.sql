-- +goose Up
-- +goose StatementBegin
create table users(
    user_id bigint primary key ,
    show bool,
    chat_id bigint,
    state text,
    lastUpdate bigint
);

create table lastDate(
    id SERIAL primary key ,
    name text,
    date bigint
);

insert into lastDate(name,date) values('update',0);

create table channels(
    id SERIAL PRIMARY KEY,
    channel_url text,
    is_telegram bool,
    last_updated bigint
);

create table sub(
    id SERIAL primary key ,
    user_id int references users(user_id),
    channel_id int references channels(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table sub
drop table channels
drop table lastDate
drop table users
-- +goose StatementEnd
