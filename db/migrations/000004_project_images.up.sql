CREATE TABLE IF NOT EXISTS "project_images" (
    id serial not null
        constraint project_images_pk
            primary key,
    project_id uuid not null
        constraint project_fk
            references "projects"
                on delete cascade on update cascade,
    media_file_id int not null
        constraint media_file_fk
            references "media_file"
                on delete cascade on update cascade,
    created_at timestamp default now()
);

CREATE UNIQUE INDEX project_images_id_uindex ON "project_images" (id);

CREATE INDEX project_images_index
    ON "project_images" (project_id ASC);