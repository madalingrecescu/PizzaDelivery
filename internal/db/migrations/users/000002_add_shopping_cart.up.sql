CREATE TABLE cart (
                      cart_id serial PRIMARY KEY,
                      user_id INTEGER REFERENCES users(user_id),
                      pizza_name VARCHAR(100) NOT NULL,
                      pizza_price DECIMAL(10, 2) NOT NULL,
                      pizza_quantity INTEGER NOT NULL
);