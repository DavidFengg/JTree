grant SELECT on JTree.* to 'select'@'localhost' identified by 'passwords';
flush privileges;
grant SELECT,INSERT, UPDATE on JTree.* to 'update'@'localhost' identified by 'passwordu';
flush privileges;