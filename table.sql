CREATE TABLE app_users (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(255),
    birthdate TIMESTAMP NOT NULL,
    isverified BOOLEAN NOT NULL
);


CREATE TABLE promos (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(255) NOT NULL,
    promo_type VARCHAR(255) NOT NULL
);


CREATE TABLE user_promo_relations (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    user_id INTEGER NOT NULL,
    promo_id INTEGER NOT NULL,
    amount DOUBLE PRECISION NOT NULL,
    promo_code VARCHAR(255) NOT NULL,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES app_users(id),
    CONSTRAINT fk_promo FOREIGN KEY (promo_id) REFERENCES promos(id),
    CONSTRAINT idx_user_promo UNIQUE (user_id, promo_id)
);


INSERT INTO app_users (created_at, updated_at, name, email, phone, birthdate, isverified)
VALUES
    ('2024-02-26 13:30:22', '2024-02-26 13:30:22', 'Alice', 'alice@ona.com', '+123456789', '2024-01-27', true),
    ('2024-02-26 13:30:22', '2024-02-26 13:30:22', 'Bob', 'bob@pol.com', NULL, '2024-02-01', false),
    ('2024-02-26 13:30:22', '2024-02-26 13:30:22', 'Clark', 'clark@kad.com', NULL, '2024-02-26', true);


INSERT INTO promo (created_at, updated_at, name, promo_type)
VALUES ('2024-02-26 13:30:22', '2024-02-26 13:30:22', 'Birthday Promo', 'birthday');
