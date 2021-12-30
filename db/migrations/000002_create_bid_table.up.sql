CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE bid (
    bidId          uuid PRIMARY KEY,
    userId    uuid NOT NULL REFERENCES user (userId) ON DELETE CASCADE ON UPDATE CASCADE,
    ttl       varchar(250) NOT NULL CHECK (title <> ''),
    price  int NOT NULL CHECK (price <> ''),
    setAt  timestamp with time zone NOT NULL DEFAULT current_timestamp
);
