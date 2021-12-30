DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS bid CASCADE;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE user (
    id        uuid PRIMARY KEY,
    name       varchar(50) UNIQUE NOT NULL CHECK (name <> ''),
    "password"  varchar(50) NOT NULL CHECK ("password" <> ''),
    role  varchar NOT NULL CHECK
);

CREATE TABLE bid (
    id          uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id   uuid NOT NULL REFERENCES user (userId) ON DELETE CASCADE ON UPDATE CASCADE,
    ttl       int NOT NULL CHECK (ttl <> ''),
    price  int NOT NULL CHECK (price <> ''),
    setAt  timestamp with time zone NOT NULL DEFAULT current_timestamp
);

INSERT INTO user 
    (id, name, password, role)
VALUES
    (1, 'admin', 'admin', 'admin'),
    (2, 'user', 'user', 'user');
