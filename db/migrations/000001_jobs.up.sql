CREATE TABLE IF NOT EXISTS "jobs" (
    id uuid default gen_random_uuid() not null
        constraint jobs_pk
            primary key,
    office varchar(50) not null,
    description text not null,
    still_working bool default false,
    start_at timestamp not null,
    end_at timestamp,
    "as" varchar(100) default 'Backend Developer' not null,
    is_full_time bool default true not null,
    created_at timestamp default now() not null,
    updated_at timestamp,
    deleted_at timestamp
);

CREATE INDEX jobs_index ON "jobs" (is_full_time asc);
