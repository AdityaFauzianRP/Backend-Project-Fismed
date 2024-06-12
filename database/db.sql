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


CREATE TABLE tax_code (
                          id SERIAL PRIMARY KEY,
                          tax VARCHAR(50),
                          created_at TIMESTAMP,
                          created_by VARCHAR,
                          updated_at TIMESTAMP,
                          updated_by VARCHAR
);

INSERT INTO tax_code (tax, created_at, created_by, updated_at, updated_by)
VALUES
    ('SWASTA', NOW(), 'admin', NOW(), 'admin'),
    ('PEMERINTAH', NOW(), 'admin', NOW(), 'admin'),
    ('PEMERINTAH NON', NOW(), 'admin', NOW(), 'admin');

select id , tax  from tax_code tc


CREATE TABLE customer (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR,
                          address_company TEXT,
                          npwp_address TEXT,
                          npwp VARCHAR,
                          ipak_number VARCHAR,
                          facture_address VARCHAR,
                          city_facture VARCHAR,
                          zip_code_facture VARCHAR,
                          number_phone_facture VARCHAR,
                          email_facture VARCHAR,
                          fax_facture VARCHAR,
                          pic_facture VARCHAR,
                          item_address VARCHAR,
                          city_item VARCHAR,
                          zip_code_item VARCHAR,
                          number_phone_item VARCHAR,
                          email_item VARCHAR,
                          fax_item VARCHAR,
                          pic_item VARCHAR,
                          contact_person VARCHAR,
                          tax_code_id INTEGER,
                          top VARCHAR,
                          created_at TIMESTAMP,
                          created_by VARCHAR,
                          updated_at TIMESTAMP,
                          updated_by VARCHAR,
                          FOREIGN KEY (tax_code_id) REFERENCES tax_code(id)
);


INSERT INTO customer (
    name, address_company, npwp_address, npwp, ipak_number,
    facture_address, city_facture, zip_code_facture, number_phone_facture, email_facture, fax_facture, pic_facture,
    item_address, city_item, zip_code_item, number_phone_item, email_item, fax_item, pic_item,
    contact_person, tax_code_id, top,
    created_at, created_by, updated_at, updated_by
) VALUES
      ('Company A', '123 Main St', '123 NPWP St', '1234567890', 'IPAK123',
       '123 Facture St', 'City A', '12345', '123-456-7890', 'emailA@company.com', '123-456-7891', 'PIC A',
       '123 Item St', 'City A', '12345', '123-456-7892', 'itemA@company.com', '123-456-7893', 'PIC Item A',
       'Contact A', 1, '30 days',
       NOW(), 'admin', NOW(), 'admin'),

      ('Company B', '456 Market St', '456 NPWP St', '0987654321', 'IPAK456',
       '456 Facture St', 'City B', '67890', '234-567-8901', 'emailB@company.com', '234-567-8902', 'PIC B',
       '456 Item St', 'City B', '67890', '234-567-8903', 'itemB@company.com', '234-567-8904', 'PIC Item B',
       'Contact B', 2, '45 days',
       NOW(), 'admin', NOW(), 'admin'),

      ('Company C', '789 Broadway', '789 NPWP St', '1122334455', 'IPAK789',
       '789 Facture St', 'City C', '54321', '345-678-9012', 'emailC@company.com', '345-678-9013', 'PIC C',
       '789 Item St', 'City C', '54321', '345-678-9014', 'itemC@company.com', '345-678-9015', 'PIC Item C',
       'Contact C', 3, '60 days',
       NOW(), 'admin', NOW(), 'admin'),

      ('Company D', '101 State St', '101 NPWP St', '2233445566', 'IPAK101',
       '101 Facture St', 'City D', '98765', '456-789-0123', 'emailD@company.com', '456-789-0124', 'PIC D',
       '101 Item St', 'City D', '98765', '456-789-0125', 'itemD@company.com', '456-789-0126', 'PIC Item D',
       'Contact D', 1, '30 days',
       NOW(), 'admin', NOW(), 'admin'),

      ('Company E', '202 Oak St', '202 NPWP St', '3344556677', 'IPAK202',
       '202 Facture St', 'City E', '87654', '567-890-1234', 'emailE@company.com', '567-890-1235', 'PIC E',
       '202 Item St', 'City E', '87654', '567-890-1236', 'itemE@company.com', '567-890-1237', 'PIC Item E',
       'Contact E', 2, '45 days',
       NOW(), 'admin', NOW(), 'admin');


select * from customer c