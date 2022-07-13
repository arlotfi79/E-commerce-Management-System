-------------------------------------------------------------------------------------

CREATE TYPE GENDER AS ENUM('MALE', 'FEMALE', 'OTHER');

-- Avoid signup for users with available email/username
CREATE TABLE Account (
    account_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    user_name VARCHAR(100) UNIQUE NOT NULL,
    password TEXT NOT NULL,
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
    description TEXT NOT NULL,
    notify_time TIMESTAMP NOT NULL
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

CREATE TABLE Cart(
    account_id INT,
    product_id INT,
    PRIMARY KEY (account_id, product_id),
    FOREIGN KEY (account_id) REFERENCES Account,
    FOREIGN KEY (product_id) REFERENCES Product,
    product_count INT NOT NULL,
    is_sold BOOLEAN DEFAULT FALSE
);

CREATE TABLE Order_Product(
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

-------------------------------------------------------------------------------------
-- Procedures

CREATE PROCEDURE CreateOrderAndClearCart(
    AccountID INTEGER,
    OrderDescription TEXT,
    OrderAddress VARCHAR(100),
    DeliveryMethod DELIVERY,
    OrderDate TIMESTAMP
)
LANGUAGE plpgsql
AS $$
BEGIN
    -- Mark products as Sold if the store has enough of them
    UPDATE Cart AS c
    SET is_sold = TRUE
    FROM Product AS p
    WHERE p.quantity >= c.product_count AND c.account_id = AccountID AND p.product_id = c.product_id;

    -- Decrease Product's Quantity based on WantedQuantity in the Cart
    UPDATE Product AS p
    SET quantity = p.quantity - c.product_count
    FROM Cart AS c
    WHERE c.account_id = AccountID AND c.is_sold = TRUE AND p.product_id = c.product_id;

    -- Now CreateOrder and insert the Cart's Products with is_sold = TRUE to the created order
    WITH OrderID AS (
        INSERT INTO orderitem (account_id, description, address, delivery_method, order_date)
        VALUES (AccountID, OrderDescription, OrderAddress, DeliveryMethod, OrderDate)
        RETURNING order_id
    )
    INSERT INTO order_product (order_id, product_id, product_count)
    SELECT order_id, product_id, product_count
    FROM (
         SELECT *
         FROM cart
         CROSS JOIN OrderID
         WHERE account_id = AccountID AND is_sold = TRUE
    ) AS cartList;

    -- Delete all products with is_sold = TRUE
    DELETE FROM Cart
    WHERE account_id = AccountID AND is_sold = TRUE;

    -- If any product is left in the cart THEN raise an exception
    IF EXISTS (
            SELECT *
            FROM cart
            WHERE account_id = AccountID AND is_sold = FALSE
        )
        THEN
            RAISE EXCEPTION 'There is a product in your cart which is not available in the store';
    END IF;
END;
$$;

CREATE PROCEDURE RefreshNotificationsAndClearAvailableWatchListProducts(
    AccountID INTEGER
)
    LANGUAGE plpgsql
AS $$
BEGIN
    IF EXISTS (
        SELECT *
        FROM WatchList AS wl
        INNER JOIN Product AS p ON wl.product_id = p.product_id
        WHERE p.quantity > 0 AND wl.account_id = AccountID
        )
        THEN
            INSERT INTO Notification (account_id, description, notify_time)
            SELECT AccountID, CONCAT('Product "', name ,'" with ID = ', product_id, ' is now available'), now()
            FROM (
                SELECT p.name, wl.product_id
                FROM WatchList AS wl
                INNER JOIN Product AS p ON wl.product_id = p.product_id
                WHERE p.quantity > 0 AND wl.account_id = AccountID
                 ) AS WatchListWithProduct;

            DELETE FROM WatchList AS wl
                USING Product AS p
            WHERE p.quantity > 0 AND wl.account_id = AccountID;
    END IF;
END;
$$;