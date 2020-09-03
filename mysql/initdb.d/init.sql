DROP DATABASE IF EXISTS fushigidane;
CREATE DATABASE fushigidane;

USE fushigidane;

DROP TABLE IF EXISTS transitpoint;
CREATE TABLE transitpoint
(
   id SERIAL PRIMARY KEY,
   address TEXT NOT NULL,
   label TEXT NOT NULL,
   latitude FLOAT,
   longitude FLOAT
);
