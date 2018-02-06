package database

import "github.com/Bio-core/jtree/models"

//BuildQuery takes a Query object and returns a string of the query
func BuildQuery(query models.Query) string {
	fields := printFields(query.SelectedFields)
	tables := printTables(query.SelectedTables)
	conditions := printConditions(query.SelectedCondition)
	queryString := "SELECT " + fields + " FROM " + tables + " WHERE (" + conditions + ")"
	return queryString
}

// Print comma seperated selected fields
func printFields(selectedFields []string) string {
	var str = ""
	for i := 0; i < len(selectedFields); i++ {
		str += selectedFields[i] + ", "
	}
	str = str[0 : len(str)-2]
	return str
}

func printTables(selectedTables []string) string {
	var str = ""
	for i := 0; i < len(selectedTables); i++ {
		str += selectedTables[i] + ", "
	}
	str = str[0 : len(str)-2]
	return str
}

func printConditions(SelectedCondition [][]string) string {
	var str = ""
	for i := 0; i < len(SelectedCondition); i++ {
		str += SelectedCondition[i][0] + " " + SelectedCondition[i][1] + SelectedCondition[i][2] + "\"" + SelectedCondition[i][3] + "\" "
	}

	str = str[4 : len(str)-1]
	return str
}
