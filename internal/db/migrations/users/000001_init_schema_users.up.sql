CREATE TABLE "users" (
                         "user_id" serial PRIMARY KEY,
                         "username" VARCHAR(50) UNIQUE NOT NULL,
                         "email" VARCHAR(100) UNIQUE NOT NULL,
                         "password_hash" VARCHAR(100) NOT NULL
);
