CREATE TABLE "pizzas" (
                          "pizza_id" serial PRIMARY KEY,
                          "name" VARCHAR(100) NOT NULL,
                          "description" TEXT,
                          "price" DOUBLE PRECISION NOT NULL
);

CREATE TABLE "pizza_orders" (
                                "order_id" serial PRIMARY KEY,
                                "customer_name" VARCHAR(100) NOT NULL,
                                "customer_phone" VARCHAR(15) NOT NULL,
                                "pizza_id" INT NOT NULL,
                                "delivery_address" TEXT NOT NULL,
                                "order_date" TIMESTAMP DEFAULT 'now()'
);


ALTER TABLE "pizza_orders" ADD FOREIGN KEY ("pizza_id") REFERENCES "pizzas" ("pizza_id");
