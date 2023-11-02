CREATE TABLE "pizzas" (
                          "pizza_id" serial PRIMARY KEY,
                          "name" VARCHAR(100) NOT NULL,
                          "description" TEXT,
                          "price" DOUBLE PRECISION NOT NULL
);