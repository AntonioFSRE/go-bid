CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE user (
    userId          uuid PRIMARY KEY,
    name       varchar(50) UNIQUE NOT NULL CHECK (name <> ''),
    "password"  varchar(50) NOT NULL CHECK ("password" <> ''),
    role  varchar(50) NOT NULL CHECK ("role" <> '')
);
