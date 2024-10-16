CREATE TABLE "users" (
                         "user_id" SERIAL PRIMARY KEY,
                         "username" VARCHAR(50) UNIQUE NOT NULL,
                         "password_hash" TEXT NOT NULL,
                         "email" VARCHAR(100) UNIQUE,
                         "role" varchar NOT NULL,
                         "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "orders" (
                          "order_id" SERIAL PRIMARY KEY,
                          "user_id" INT NOT NULL,
                          "order_description" TEXT,
                          "order_type" VARCHAR(50) NOT NULL,
                          "state_id" INT NOT NULL,
                          "priority_id" INT NOT NULL,
                          "goods_ids" varchar(512) NOT NULL,
                          "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
                          "updated_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "order_priorites" (
                                   "priority_id" SERIAL PRIMARY KEY,
                                   "priority" varchar NOT NULL
);

CREATE TABLE "order_states" (
                                "state_id" SERIAL PRIMARY KEY,
                                "state_name" VARCHAR(50) NOT NULL
);

CREATE TABLE "payments" (
                            "payment_id" SERIAL PRIMARY KEY,
                            "order_id" INT NOT NULL,
                            "payment_system" VARCHAR(50),
                            "amount" DECIMAL(10,2) NOT NULL,
                            "status" VARCHAR(50) NOT NULL,
                            "paid_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "goods" (
                         "id" SERIAL PRIMARY KEY,
                         "name" varchar NOT NULL,
                         "category_id" int NOT NULL,
                         "price" decimal(10,2) NOT NULL,
                         "description" varchar NOT NULL,
                         "weight" decimal(10,2)
);

CREATE TABLE "goods_orders" (
                                "good_id" int NOT NULL,
                                "order_id" int NOT NULL
);

CREATE TABLE "category" (
                            "id" SERIAL PRIMARY KEY,
                            "name" varchar NOT NULL
);

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "orders" ADD FOREIGN KEY ("state_id") REFERENCES "order_states" ("state_id");

ALTER TABLE "payments" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("order_id");

ALTER TABLE "orders" ADD FOREIGN KEY ("priority_id") REFERENCES "order_priorites" ("priority_id");

ALTER TABLE "goods_orders" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("order_id");

ALTER TABLE "goods_orders" ADD FOREIGN KEY ("good_id") REFERENCES "goods" ("id");

ALTER TABLE "goods" ADD FOREIGN KEY ("category_id") REFERENCES "category" ("id");
