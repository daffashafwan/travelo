CREATE TABLE IF NOT EXISTS destinations(
   destination_id uuid PRIMARY KEY,
   destination_name VARCHAR (50) NOT NULL,
   destination_location VARCHAR (50) NOT NULL,
   destination_description TEXT NOT NULL,
   destination_reviews INT NOT NULL,
   destination_price FLOAT NOT NULL,
   destination_gimmick_price FLOAT,
   destination_category_id uuid,
   CONSTRAINT fk_destination
      FOREIGN KEY(destination_category_id) 
	  REFERENCES categories(category_id)
);