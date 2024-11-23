CREATE TABLE users
(
    id   uuid         NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL
);

COMMENT ON TABLE users IS
    'Таблица для хранения пользователей';

COMMENT ON COLUMN users.id IS 'Ид записи';
COMMENT ON COLUMN users.name IS 'Имя';

ALTER TABLE users
    ADD
        CONSTRAINT users_pk
            PRIMARY KEY
                (id);