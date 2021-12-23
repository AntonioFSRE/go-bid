CREATE TABLE public.bid (
    bidId int PRIMARY KEY,
    ttl  int,
    price int,
    setAt timestamp,
    userId int
);

CREATE TABLE public.user (
    userId int PRIMARY KEY,
    username  varchar(64),
    password  varchar(64),
    role varchar(64)
);

INSERT INTO public.user 
    (userId, username, password, role)
VALUES
    (1, 'admin', 'admin', 'admin'),
    (2, 'user', 'user', 'user');
