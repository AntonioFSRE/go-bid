DROP TABLE IF EXISTS public.user CASCADE;
DROP TABLE IF EXISTS public.bid CASCADE;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE public.user (
    id         uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name       varchar(50),
    password   varchar(50),
    role       varchar(50)
);

CREATE TABLE public.bid (
    id        int PRIMARY KEY,
    user_id   uuid NOT NULL REFERENCES public.user (id),
    ttl       int NOT NULL,
    price     int NOT NULL,
    setAt     timestamp with time zone NOT NULL DEFAULT current_timestamp
);

INSERT INTO public.user 
    (name, password, role)
VALUES
    ('admin', 'admin', 'admin'),
    ('user', 'user', 'user');
