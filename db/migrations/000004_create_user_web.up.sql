CREATE USER 'web'@'%' IDENTIFIED BY 'pass';

GRANT SELECT, INSERT, UPDATE, DELETE
ON snippetbox.*
TO 'web'@'%';
