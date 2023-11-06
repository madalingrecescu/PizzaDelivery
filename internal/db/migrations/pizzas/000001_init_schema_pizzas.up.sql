CREATE TABLE "pizzas" (
                          "pizza_id" serial PRIMARY KEY,
                          "name" VARCHAR(100) NOT NULL UNIQUE,
                          "description" VARCHAR(100) NOT NULL,
                          "price" DOUBLE PRECISION NOT NULL
);