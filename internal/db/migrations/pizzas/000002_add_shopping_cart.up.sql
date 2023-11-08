CREATE TABLE "shopping_cart" (
                               "shopping_cart_id" SERIAL PRIMARY KEY,
                               "username" VARCHAR(100) NOT NULL
);

CREATE TABLE "pizza_order" (
                             "pizza_order_id" SERIAL PRIMARY KEY,
                             "shopping_cart_id" INTEGER NOT NULL REFERENCES "shopping_cart"("shopping_cart_id"),
                             "pizza_name" VARCHAR(100) NOT NULL,
                             "pizza_price" DOUBLE PRECISION NOT NULL,
                             "quantity" INTEGER NOT NULL
);