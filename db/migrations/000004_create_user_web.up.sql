-- If you want to allow the web user to connect from any host,
-- you can create an account with a wildcard host:
CREATE USER 'web'@'%' IDENTIFIED BY 'pass';

GRANT SELECT, INSERT, UPDATE, DELETE
ON snippetbox.*
TO 'web'@'%';
