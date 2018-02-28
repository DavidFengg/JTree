grant SELECT on JTree.* to 'select'@'%' identified by 'passwords';
flush privileges;
grant INSERT, UPDATE on JTree.* to 'update'@'%' identified by 'passwordu';
flush privileges;