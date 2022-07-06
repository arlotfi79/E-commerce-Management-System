-------------------------------------------------------------------------------------

CREATE EXTENSION pgcrypto; -- To Hash Passwords
CREATE TYPE GENDER AS ENUM('MALE', 'FEMALE', 'OTHER');

CREATE TABLE User (
    national_code VARCHAR(100) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    user_name VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    phone_number VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    gender GENDER NOT NULL,
    birth_date DATE NOT NULL,
    join_date TIMESTAMP NOT NULL
);

CREATE TABLE Address(
    address_id SERIAL PRIMARY KEY,
    national_code VARCHAR(100) REFERENCES User, -- FK
    country VARCHAR(25) NOT NULL,
    city VARCHAR(25) NOT NULL,
    street VARCHAR(25) NOT NULL,
    plaque VARCHAR(25) NOT NULL
);

-------------------------------------------------------------------------------------

CREATE TYPE DELIVERY AS ENUM('PREMIUM', 'NORMAL', 'CHEAP');
CREATE TABLE Order(
    order_id SERIAL PRIMARY KEY,
    national_code VARCHAR(100) REFERENCES User, -- FK
    description TEXT,
    address VARCHAR(100) NOT NULL,
    delivery_method DELIVERY NOT NULL ,
    order_date TIMESTAMP NOT NULL
);

CREATE TABLE TicketTracking(
    ticket_id SERIAL PRIMARY KEY,
    order_id INT REFERENCES Order,
    subject VARCHAR(50) NOT NULL,
    ticket_date TIMESTAMP NOT NULL
);

CREATE TABLE Message(
    message_id SERIAL PRIMARY KEY,
    ticket_id INT REFERENCES TicketTracking, -- FK
    message_text VARCHAR(200) NOT NULL,
    attachment BYTEA,
    message_date TIMESTAMP NOT NULL
);

-------------------------------------------------------------------------------------

CREATE TYPE COLOR AS ENUM('RED', 'BLUE', 'ORANGE', 'GREEN', 'WHITE', 'BLACK');
CREATE TABLE Product(
    product_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    ProductImage BYTEA,
    color COLOR,
    price FLOAT CHECK ( price > 0 ) NOT NULL,
    weight FLOAT CHECK ( weight > 0 ) NOT NULL,
    Quantity INT NOT NULL
);

CREATE TABLE Review(
    review_id SERIAL PRIMARY KEY,
    product_id INT REFERENCES Product, -- FK
    national_code VARCHAR(100) REFERENCES User,
    rating INT CHECK ( 0 <= rating AND rating <= 5 )
);

CREATE TABLE Store(
    store_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    phone_number VARCHAR(100) NOT NULL
);

CREATE TABLE Category(
    category_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

-------------------------------------------------------------------------------------
-- Many to Many relations

CREATE TABLE Order_Product_Counter(
    order_id INT,
    product_id INT,
    PRIMARY KEY (product_id, order_id),
    FOREIGN KEY (product_id) REFERENCES Product,
    FOREIGN KEY (order_id) REFERENCES Order,
    product_count INT NOT NULL
);

CREATE TABLE Product_Category(
    product_id INT,
    category_id INT,
    PRIMARY KEY (product_id, category_id),
    FOREIGN KEY (category_id) REFERENCES Category,
    FOREIGN KEY (product_id) REFERENCES Product
);

CREATE TABLE Product_Store(
    product_id INT,
    store_id INT,
    PRIMARY KEY (product_id, store_id),
    FOREIGN KEY (product_id) REFERENCES Product,
    FOREIGN KEY (store_id) REFERENCES Store
);
