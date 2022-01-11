CREATE TABLE IF NOT EXISTS "media_file" (
    id serial not null
        constraint media_file_pk
            primary key,
    filename varchar not null,
    url varchar,
    "type" varchar(25) default 'image' not null,
    source varchar(25) default 'internal' not null,
    created_at timestamp default now() not null
);

CREATE UNIQUE INDEX media_file_id_uindex ON "media_file" (id);