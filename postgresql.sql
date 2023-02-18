CREATE TABLE IF NOT EXISTS Movies(
  id BIGSERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description VARCHAR(1000),
  rating FLOAT, 
  image VARCHAR(1000),
  created_at TIMESTAMP,
  update_at TIMESTAMP
);

INSERT INTO
  Movies (title, description, rating, image, created_at, update_at)
VALUES
  ('Pengabdi Setan 2 Comunion', 'Sebuah film horor Indonesia tahun 2022 yang disutradarai dan
    ditulis oleh Joko Anwar sebagai sekuel dari film tahun 2017, Pengabdi
    Setan.', 7, '', '2023-02-17 16:00:00', '2023-02-17 16:00:00'),
  ('Pengabdi Setan ', '', 8, '', '2023-02-17 16:00:00', '2023-02-17 16:00:00');