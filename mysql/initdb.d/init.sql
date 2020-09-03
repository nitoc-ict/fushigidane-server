DROP DATABASE IF EXISTS fushigidane;
CREATE DATABASE fushigidane;

USE fushigidane;

DROP TABLE IF EXISTS transitpoints;
CREATE TABLE transitpoints
(
   id SERIAL PRIMARY KEY,
   address TEXT NOT NULL,
   label TEXT NOT NULL,
   latitude FLOAT,
   longitude FLOAT
);

INSERT INTO transitpoints(address, label, latitude, longitude) VALUES ('TEST', 'TEST', 0.0, 0.0);
