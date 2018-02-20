mysqldump -u root -p Jtree > ./sql/jtree_backup.sql
mysql -u "root" "-pwaterloo" < "./sql/DropDatabase.sql"
mysql -u "root" "-pwaterloo" < "./sql/CreateDatabase.sql"
mysql -u "root" "-pwaterloo" < "./sql/CreateTables.sql"


