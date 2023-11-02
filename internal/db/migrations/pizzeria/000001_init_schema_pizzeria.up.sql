CREATE TABLE "pizzas" (
                          "pizza_id" serial PRIMARY KEY,
                          "name" VARCHAR(100) NOT NULL,
                          "description" TEXT,
                          "price" DOUBLE PRECISION NOT NULL
);

ALTER TABLE "pizza_orders" ADD FOREIGN KEY ("pizza_id") REFERENCES "pizzas" ("pizza_id");
