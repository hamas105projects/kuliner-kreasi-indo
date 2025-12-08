-- DDL query for creating table


CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    email VARCHAR(120) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    role VARCHAR(20) NOT NULL,
    
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(150) NOT NULL,
    price NUMERIC(12,2) NOT NULL CHECK (price >= 0),
    description TEXT,

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);


-- header transaction

CREATE TABLE sale_orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    cashier_id UUID NOT NULL,
    total_amount NUMERIC(12,2) NOT NULL DEFAULT 0,
    notes TEXT,

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP,

    CONSTRAINT fk_sale_order_cashier FOREIGN KEY (cashier_id)
        REFERENCES users (id)
        ON UPDATE CASCADE ON DELETE RESTRICT
);

--transaction details
CREATE TABLE sale_order_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    sale_order_id UUID NOT NULL,
    product_id UUID NOT NULL,
    qty INT NOT NULL CHECK (qty > 0),
    price_snapshot NUMERIC(12,2) NOT NULL CHECK (price_snapshot >= 0),

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP,

    CONSTRAINT fk_item_order FOREIGN KEY (sale_order_id)
        REFERENCES sale_orders (id)
        ON UPDATE CASCADE ON DELETE CASCADE,

    CONSTRAINT fk_item_product FOREIGN KEY (product_id)
        REFERENCES products (id)
        ON UPDATE CASCADE ON DELETE RESTRICT
);


-- next step DML for query insert data, in this time just dummy

INSERT INTO products (name, price, description)
VALUES
('Ayam Geprek Mozarella', 25000, 'Ayam geprek pedas dengan topping keju mozarella leleh'),
('Nasi Goreng Kampung', 20000, 'Nasi goreng khas kampung dengan telur dan sambal terasi'),
('Es Kopi Susu Gula Aren', 18000, 'Minuman kopi susu kekinian dengan gula aren asli'),
('Mie Pedas Level 10', 22000, 'Mie instan dengan bumbu pedas ekstrem, bisa pilih level'),
('Burger Sambal Matah', 30000, 'Burger daging sapi dengan sambal matah khas Bali'),
('Pisang Coklat Lumer', 15000, 'Dessert pisang goreng isi coklat yang meleleh');