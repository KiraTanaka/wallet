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

CREATE TABLE wallets
(
    id      uuid          NOT NULL DEFAULT gen_random_uuid(),
    user_id uuid          NOT NULL,
    balance NUMERIC(7, 2) NOT NULL
);

COMMENT ON TABLE wallets IS
    'Таблица для хранения кошельков пользователей';

COMMENT ON COLUMN wallets.id IS 'Ид записи';
COMMENT ON COLUMN wallets.user_id IS 'Ид пользователя (FK users.id)';
COMMENT ON COLUMN wallets.balance IS 'Баланс';

ALTER TABLE wallets
    ADD
        CONSTRAINT wallets_pk
            PRIMARY KEY
                (id);

ALTER TABLE wallets
    ADD
        CONSTRAINT wallets_fk1
            FOREIGN KEY (user_id)
                REFERENCES users (id);

CREATE TYPE wallet_operation_type AS ENUM (
    'DEPOSIT',
    'WITHDRAW'
    );

CREATE TABLE wallet_operations
(
    id             uuid                  NOT NULL DEFAULT gen_random_uuid(),
    wallet_id      uuid                  NOT NULL,
    operation_type wallet_operation_type NOT NULL,
    amount         NUMERIC(7, 2)         NOT NULL,
    operation_date TIMESTAMP             NOT NULL DEFAULT NOW()
);

COMMENT ON TABLE wallet_operations IS
    'Таблица для хранения кошельков пользователей';

COMMENT ON COLUMN wallet_operations.id IS 'Ид записи';
COMMENT ON COLUMN wallet_operations.wallet_id IS 'Ид кошелька (FK wallets.id)';
COMMENT ON COLUMN wallet_operations.operation_type IS 'Тип операции';
COMMENT ON COLUMN wallet_operations.amount IS 'Сумма';
COMMENT ON COLUMN wallet_operations.operation_date IS 'Дата операции';

ALTER TABLE wallet_operations
    ADD
        CONSTRAINT wallet_operations_pk
            PRIMARY KEY
                (id);

ALTER TABLE wallet_operations
    ADD
        CONSTRAINT wallet_operations_fk1
            FOREIGN KEY (wallet_id)
                REFERENCES wallets (id);