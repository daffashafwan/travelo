CREATE TABLE IF NOT EXISTS reviews(
   review_id uuid PRIMARY KEY,
   review_user_name VARCHAR (50) NOT NULL,
   review_user_icon TEXT NOT NULL,
   review_user_location VARCHAR (50) NOT NULL,
   review_description TEXT NOT NULL,
   review_location VARCHAR (50) NOT NULL,
   review_date VARCHAR (50) NOT NULL,
   review_star INT NOT NULL
);