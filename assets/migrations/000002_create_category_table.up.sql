CREATE TABLE IF NOT EXISTS categories(
   category_id uuid PRIMARY KEY,
   category_name VARCHAR (50) UNIQUE NOT NULL,
   category_icon TEXT NOT NULL
);