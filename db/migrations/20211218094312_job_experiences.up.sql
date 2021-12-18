CREATE TABLE IF NOT EXISTS "job_experiences" (
    id serial,
    office varchar(50) not null,
    start_at date not null,
    end_at date,
    "description" text not null,
    still_working boolean default false not null,
    created_at timestamp default now() not null,
    updated_at timestamp
);

CREATE UNIQUE INDEX job_experiences_id_uindex ON "job_experiences" (id);

ALTER TABLE "job_experiences" ADD CONSTRAINT job_experiences_pk PRIMARY KEY (id);