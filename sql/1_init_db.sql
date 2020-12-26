create table if not exists users (
    id  serial NOT NULL primary key,
    username varchar(64) NOT NULL,
    password varchar(128) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp,
    deleted_at timestamp
);
