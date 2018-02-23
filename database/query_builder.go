package database

import (
	"fmt"
	"reflect"

	strings "strings"

	"github.com/Bio-core/jtree/models"
)

//Map for types
var Map map[string]string

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
		str += SelectedCondition[i][0] + " " + SelectedCondition[i][1] + SelectedCondition[i][2]
		if Map[SelectedCondition[i][1]] == "*string" {
			str += "\"" + SelectedCondition[i][3] + "\" "
		} else if Map[SelectedCondition[i][1]] == "*float32" {
			str += SelectedCondition[i][3] + " "

		}
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
		if Map[condition[1]] != "*string" && Map[condition[1]] != "*float32" {
			condition[2] = ""
			return nil
		}
		condition[2] = "="
		break
	case "Not equal to":
		if Map[condition[1]] != "*string" && Map[condition[1]] != "*float32" {
			condition[2] = ""
			return nil
		}
		condition[2] = "<>"

		break
	case "Greater than":
		if Map[condition[1]] != "*float32" {
			condition[2] = ""
			return nil
		}
		condition[2] = ">"
		break
	case "Less than":
		if Map[condition[1]] != "*float32" {
			condition[2] = ""
			return nil
		}
		condition[2] = "<"
		break
	case "Greater or equal to":
		if Map[condition[1]] != "*float32" {
			condition[2] = ""
			return nil
		}
		condition[2] = ">="
		break
	case "Less or equal to":
		if Map[condition[1]] != "*float32" {
			condition[2] = ""
			return nil
		}
		condition[2] = "<="
		break
	case "Begins with":
		if Map[condition[1]] != "*string" {
			condition[2] = ""
			return nil
		}
		condition[2] = " LIKE "
		condition[3] += "%"
		break
	case "Not begins with":
		if Map[condition[1]] != "*string" {
			condition[2] = ""
			return nil
		}
		condition[0] += " NOT"
		condition[2] = " LIKE "
		condition[3] += "%"
		break
	case "Ends with":
		if Map[condition[1]] != "*string" {
			condition[2] = ""
			return nil
		}
		condition[2] = " LIKE "
		condition[3] = "%" + condition[3]
		break
	case "Not ends with":
		if Map[condition[1]] != "*string" {
			condition[2] = ""
			return nil
		}
		condition[0] += " NOT"
		condition[2] = " LIKE "
		condition[3] = "%" + condition[3]
		break
	case "Contains":
		if Map[condition[1]] != "*string" {
			condition[2] = ""
			return nil
		}
		condition[2] = " LIKE "
		condition[3] = "%" + condition[3] + "%"
		break
	case "Not contains":
		if Map[condition[1]] != "*string" {
			condition[2] = ""
			return nil
		}
		condition[0] += " NOT"
		condition[2] = " LIKE "
		condition[3] = "%" + condition[3] + "%"
		break
	default:
		return nil
	}
	return condition
}

//MapSuper makes a map
func MapSuper() map[string]string {
	m := make(map[string]string)
	v := reflect.ValueOf(models.Patient{})

	for i := 0; i < v.NumField(); i++ {
		tag := string(reflect.TypeOf(models.Patient{}).Field(i).Tag)
		runes := []rune(tag)
		j := strings.Index(tag, ":")
		k := strings.Index(tag, "omit")
		tag = string(runes[j+2 : k-1])
		varType := reflect.TypeOf(models.Patient{}).Field(i).Type.String()
		m[tag] = varType
	}
	v = reflect.ValueOf(models.Sample{})

	for i := 0; i < v.NumField(); i++ {
		tag := string(reflect.TypeOf(models.Sample{}).Field(i).Tag)
		runes := []rune(tag)
		j := strings.Index(tag, ":")
		k := strings.Index(tag, "omit")
		tag = string(runes[j+2 : k-1])
		varType := reflect.TypeOf(models.Sample{}).Field(i).Type.String()
		m[tag] = varType
	}
	v = reflect.ValueOf(models.Experiment{})

	for i := 0; i < v.NumField(); i++ {
		tag := string(reflect.TypeOf(models.Experiment{}).Field(i).Tag)
		runes := []rune(tag)
		j := strings.Index(tag, ":")
		k := strings.Index(tag, "omit")
		tag = string(runes[j+2 : k-1])
		varType := reflect.TypeOf(models.Experiment{}).Field(i).Type.String()
		m[tag] = varType
	}

	return m
}
