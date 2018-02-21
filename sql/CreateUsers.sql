grant SELECT on Jtree.* to 'select'@'%' identified by 'passwords';
flush privileges;
grant INSERT, UPDATE on Jtree.* to 'update'@'%' identified by 'passwordu';
flush privileges;