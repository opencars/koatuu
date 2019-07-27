
CREATE TABLE `level1_territories` (
  `id`   SERIAL PRIMARY KEY,
  `code` VARCHAR (10) UNIQUE NOT NULL,
  `name` VARCHAR (1024)
  `type` VARCHAR (10)
);

CREATE TABLE `level2_territories` (
  `id`   SERIAL PRIMARY KEY,
  `code` VARCHAR (10) UNIQUE NOT NULL,
  `name` VARCHAR (1024)
  `type` VARCHAR (10)
  `level1_id` INTEGER NOT NULL,
);

CREATE TABLE `level3_territories` (
  `id`   SERIAL PRIMARY KEY,
  `code` VARCHAR (10) UNIQUE NOT NULL,
  `name` VARCHAR (1024)
  `type` VARCHAR (10)
  `level2_id` INTEGER NOT NULL,
);

CREATE TABLE `level4_territories` (
  `id`   SERIAL PRIMARY KEY,
  `code` VARCHAR (10) UNIQUE NOT NULL,
  `name` VARCHAR (1024)
  `type` VARCHAR (10)
  `level3_id` INTEGER NOT NULL,
);