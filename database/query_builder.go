package database

import (
	"fmt"

	strings "strings"

	"github.com/Bio-core/jtree/models"
)

//BuildQuery takes a Query object and returns a string of the query
func BuildQuery(query models.Query) string {
	if len(query.SelectedFields) == 1 && query.SelectedFields[0] == "*" {
		query.SelectedFields = GetColumns(query.SelectedTables)
	}
	fields := printFields(query.SelectedFields)
	tables := printTables(query.SelectedTables)
	queryString := "SELECT " + fields + " FROM " + tables
	if len(query.SelectedCondition) != 0 {
		if len(query.SelectedCondition[0]) != 0 {
			conditions := printConditions(query.SelectedCondition)
			queryString += " WHERE (" + conditions + ")"
		}
	}
	return queryString
}

// Print comma separated selected fields
func printFields(selectedFields []string) string {
	var str = ""
	for i := 0; i < len(selectedFields); i++ {
		str += selectedFields[i] + " AS '" + selectedFields[i] + "', "
	}
	str = str[0 : len(str)-2]
	return str
}

func printTables(selectedTables []string) string {
	var str = ""
	for i := 0; i < len(selectedTables); i++ {
		if i == 0 {
			str += selectedTables[i]

			// } else {
			// 	str += " JOIN " + selectedTables[i] + " ON " + "patients.sample_id = samples.sample_id"
		} else {
			str += " JOIN " + selectedTables[i] + " ON " + selectedTables[i-1] + ".sample_id = " + selectedTables[i] + ".sample_id"
		}

	}
	return str
}

func printConditions(SelectedCondition [][]string) string {
	var str = ""
	for i := 0; i < len(SelectedCondition); i++ {
		SelectedCondition[i][3] = escapeChars(SelectedCondition[i][3])
		SelectedCondition[i] = formatCondition(SelectedCondition[i])
		if SelectedCondition[i] == nil {
			return "0=1"
		}
		str += SelectedCondition[i][0] + " " + SelectedCondition[i][1] + SelectedCondition[i][2] + "\"" + SelectedCondition[i][3] + "\" "
	}

	str = str[4 : len(str)-1]
	return str
}

//GetColumns returns colums based off of table names
func GetColumns(tables []string) []string {
	var columns []string
	for _, tableName := range tables {
		rows, err := DBSelect.Query("Select * from " + tableName + " where 0=1")
		defer rows.Close()
		if err != nil {
			fmt.Println(err)
			return nil
		}
		columnsSet, err := rows.Columns()
		if err != nil {
			fmt.Println(err)
			return nil
		}
		for _, j := range columnsSet {
			columns = append(columns, tableName+"."+j)
		}
	}
	return columns
}

//GetTables gets all of the tables in the db
func GetTables() []string {
	var tables []string
	rows, err := DBSelect.Query("Show Tables")
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		var tname string
		rows.Scan(&tname)
		tables = append(tables, strings.ToLower(tname))
	}
	return tables
}

func formatCondition(condition []string) []string {
	switch condition[2] {
	case "Equal to":
		condition[2] = "="
		break
	case "Not equal to":
		condition[2] = "<>"
		break
	case "Greater than":
		condition[2] = ">"
		break
	case "Less than":
		condition[2] = "<"
		break
	case "Greater or equal to":
		condition[2] = ">="
		break
	case "Less or equal to":
		condition[2] = "<="
		break
	case "Begins with":
		condition[2] = " LIKE "
		condition[3] += "%"
		break
	case "Not begins with":
		condition[0] += " NOT"
		condition[2] = " LIKE "
		condition[3] += "%"
		break
	case "Ends with":
		condition[2] = " LIKE "
		condition[3] = "%" + condition[3]
		break
	case "Not ends with":
		condition[0] += " NOT"
		condition[2] = " LIKE "
		condition[3] = "%" + condition[3]
		break
	case "Contains":
		condition[2] = " LIKE "
		condition[3] = "%" + condition[3] + "%"
		break
	case "Not contains":
		condition[0] += " NOT"
		condition[2] = " LIKE "
		condition[3] = "%" + condition[3] + "%"
		break
	default:
		return nil
	}
	return condition
}
