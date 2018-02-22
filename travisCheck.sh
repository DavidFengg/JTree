	if [ !$CI ]
	then
		mysql -u "root" "-pwaterloo" < "./tests/sql/DeleteData.sql" 
	fi