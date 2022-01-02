CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE public.user (
    id          uuid PRIMARY KEY,
    name       varchar(50) UNIQUE NOT NULL,
    password  varchar(50) NOT NULL,
    role  varchar(50) NOT NULL
);
