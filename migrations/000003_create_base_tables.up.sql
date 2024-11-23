CREATE TABLE users2
(
    id   uuid         NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL
);