CREATE TABLE "users" (
                         "user_id" serial PRIMARY KEY,
                         "username" VARCHAR(50) UNIQUE NOT NULL,
                         "email" VARCHAR(100) UNIQUE NOT NULL,
                         "hashed_password" VARCHAR(100) NOT NULL,
                         "phone_number" VARCHAR(100) NOT NULL,
                         "role" VARCHAR(100) NOT NULL
);
