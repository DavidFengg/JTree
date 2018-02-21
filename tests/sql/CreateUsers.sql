grant SELECT on TestJtree.* to 'select'@'%' identified by 'passwords';
flush privileges;
grant INSERT, UPDATE on TestJtree.* to 'update'@'%' identified by 'passwordu';
flush privileges;