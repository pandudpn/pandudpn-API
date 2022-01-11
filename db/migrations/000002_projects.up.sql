CREATE TABLE IF NOT EXISTS "projects" (
    id uuid default gen_random_uuid() not null
        constraint projects_pk
            primary key,
    name varchar not null,
    description text not null,
    demo_link varchar,
    slug varchar not null,
    start_at timestamp not null,
    created_at timestamp default now() not null,
    updated_at timestamp,
    deleted_at timestamp
);

CREATE UNIQUE INDEX projects_id_uindex ON "projects" (id);