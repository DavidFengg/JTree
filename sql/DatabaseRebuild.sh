mysqldump -u root -p Jtree > jtree_backup.sql
mysql -u "root" "-pwaterloo" < "DropDatabase.sql"
mysql -u "root" "-pwaterloo" < "CreateDatabase.sql"
mysql -u "root" "-pwaterloo" < "CreateTables.sql"


