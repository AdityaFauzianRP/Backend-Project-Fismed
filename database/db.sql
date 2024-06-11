-- Membuat tabel user_category
CREATE TABLE user_category (
                               id SERIAL PRIMARY KEY,
                               role VARCHAR(50) NOT NULL,
                               created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                               created_by VARCHAR(255),
                               updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                               updated_by VARCHAR(255)
);

-- Membuat tabel user
CREATE TABLE "user" (
                        id SERIAL PRIMARY KEY,
                        username VARCHAR(50) UNIQUE NOT NULL,
                        password VARCHAR(255),
                        role_id INT REFERENCES user_category(id),
                        token TEXT NOT NULL,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        created_by VARCHAR(255),
                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        updated_by VARCHAR(255)
);

INSERT INTO user_category (role, created_at, created_by, updated_at, updated_by)
VALUES
    ('SALES', NOW(), 'system', NOW(), 'system'),
    ('ADMIN', NOW(), 'system', NOW(), 'system'),
    ('SUPER ADMIN', NOW(), 'system', NOW(), 'system'),
    ('LOGISTIK', NOW(), 'system', NOW(), 'system'),
    ('KEUANGAN', NOW(), 'system', NOW(), 'system');

select * from user_category uc

-- Menambahkan data dummy ke tabel user dengan password yang dienkripsi menggunakan MD5
    INSERT INTO "user" (username, password, role_id, token, created_at, created_by, updated_at, updated_by)
VALUES
    ('sales1', md5('sales123'), 1, 'some_random_token', NOW(), 'system', NOW(), 'system');

select * from users u