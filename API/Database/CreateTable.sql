-------------------------------------------------------------------------------------

CREATE TYPE GENDER AS ENUM('MALE', 'FEMALE', 'OTHER');

-- Avoid signup for users with available email/username
CREATE TABLE Account (
    account_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    user_name VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    phone_number VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    gender GENDER NOT NULL,
    birth_date DATE NOT NULL,
    join_date TIMESTAMP NOT NULL
);

CREATE TABLE Address(
    address_id SERIAL PRIMARY KEY,
    account_id INT REFERENCES Account, -- FK
    country VARCHAR(25) NOT NULL,
    city VARCHAR(25) NOT NULL,
    street VARCHAR(25) NOT NULL,
    plaque VARCHAR(25) NOT NULL
);

CREATE TABLE Notification (
    notification_id SERIAL PRIMARY KEY,
    account_id INT REFERENCES Account, -- FK
    subject VARCHAR(100) NOT NULL,
    description TEXT NOT NULL
);

-------------------------------------------------------------------------------------

CREATE TYPE DELIVERY AS ENUM('PREMIUM', 'NORMAL', 'CHEAP');
CREATE TABLE OrderItem(
    order_id SERIAL PRIMARY KEY,
    account_id INT REFERENCES Account, -- FK
    description TEXT,
    address VARCHAR(100) NOT NULL,
    delivery_method DELIVERY NOT NULL ,
    order_date TIMESTAMP NOT NULL
);

CREATE TABLE PromotionCode(
    promotion_code_id SERIAL PRIMARY KEY,
    order_id INT REFERENCES OrderItem, -- FK
    value FLOAT CHECK ( value > 0 ) NOT NULL,
    expire_date TIMESTAMP NOT NULL
);

CREATE TABLE TicketTracking(
    ticket_id SERIAL PRIMARY KEY,
    order_id INT REFERENCES OrderItem, -- FK
    subject VARCHAR(50) NOT NULL,
    ticket_date TIMESTAMP NOT NULL
);

CREATE TABLE Message(
    message_id SERIAL PRIMARY KEY,
    ticket_id INT REFERENCES TicketTracking, -- FK
    message_text VARCHAR(200) NOT NULL,
    message_date TIMESTAMP NOT NULL
);

-------------------------------------------------------------------------------------

CREATE TYPE COLOR AS ENUM('RED', 'BLUE', 'ORANGE', 'GREEN', 'WHITE', 'BLACK');
CREATE TABLE Product(
    product_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    color COLOR,
    price FLOAT CHECK ( price > 0 ) NOT NULL,
    weight FLOAT CHECK ( weight > 0 ) NOT NULL,
    Quantity INT NOT NULL
);

CREATE TABLE Review(
    review_id SERIAL PRIMARY KEY,
    product_id INT REFERENCES Product, -- FK
    account_id INT REFERENCES Account, -- FK
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

CREATE TABLE Reaction (
    review_id INT,
    account_id INT,
    PRIMARY KEY (account_id, review_id),
    FOREIGN KEY (account_id) REFERENCES Account,
    FOREIGN KEY (review_id) REFERENCES Review,
    up_vote BOOLEAN DEFAULT FALSE,
    down_vote BOOLEAN DEFAULT FALSE
);

CREATE TABLE WatchList(
    account_id INT,
    product_id INT,
    PRIMARY KEY (account_id, product_id),
    FOREIGN KEY (account_id) REFERENCES Account,
    FOREIGN KEY (product_id) REFERENCES Product
);

CREATE TABLE Order_Product_Counter(
    order_id INT,
    product_id INT,
    PRIMARY KEY (product_id, order_id),
    FOREIGN KEY (product_id) REFERENCES Product,
    FOREIGN KEY (order_id) REFERENCES OrderItem,
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
