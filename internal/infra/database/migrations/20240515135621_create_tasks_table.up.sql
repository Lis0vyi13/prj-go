CREATE TABLE IF NOT EXISTS public.tasks
(
    id              serial PRIMARY KEY,
    user_id         integer references public.users(id),
    "name"          text NOT NULL,
    "description"   text NOT NULL,
    deadline        timestamp NOT NULL,
    "status"        varchar(20) NOT NULL,
    created_date    timestamp NOT NULL,
    updated_date    timestamp NOT NULL,
    deleted_date    timestamp NULL
);
