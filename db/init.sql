CREATE DATABASE IF NOT EXISTS shiny;

-- create the users for each database
CREATE USER 'admin'@'%' IDENTIFIED BY 'abc123';
GRANT ALL PRIVILEGES ON *shiny.* TO 'admin'@'%';

CREATE USER 'admin'@'localhost' IDENTIFIED BY 'abc123';
GRANT ALL PRIVILEGES ON *shiny.* TO 'admin'@'localhost';

FLUSH PRIVILEGES;