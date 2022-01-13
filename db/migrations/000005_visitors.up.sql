CREATE TABLE IF NOT EXISTS "visitors" (
    id uuid default gen_random_uuid() not null
        constraint visitors_pk
            primary key,
    ip varchar(50) not null,
    user_agent varchar not null,
    total_visit int default 1 not null,
    created_at timestamp not null,
    updated_at timestamp
);

CREATE INDEX visitors_index
    ON "visitors" (ip ASC, user_agent ASC);