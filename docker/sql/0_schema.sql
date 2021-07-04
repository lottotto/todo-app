CREATE DATABASE sample;
\c sample;

CREATE TABLE task (
  id SERIAL NOT NULL,
  user_id int NOT NULL,
  type_id int NOT NULL,
  title varchar NOT NULL,
  detail text,
  deadline timestamp NOT NULL,
  PRIMARY KEY (id)
) ;