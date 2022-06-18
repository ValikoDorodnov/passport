-- +goose Up
-- +goose StatementBegin
-- auto-generated definition
create table article
(
    id   serial
        constraint article_pk
            primary key,
    name varchar,
    text text
);

create unique index article_id_uindex
    on article (id);

create unique index article_name_uindex
    on article (name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table article;
-- +goose StatementEnd
