CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE public.bid (
    id          uuid PRIMARY KEY,
    user_id    uuid NOT NULL REFERENCES user (id),
    ttl       int(50) NOT NULL,
    price  int NOT NULL,
    setAt  timestamp with time zone NOT NULL DEFAULT current_timestamp
);
