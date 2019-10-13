package query

import (
	"strings"

	str "../str"
)

func PrepareInsertQuery(tableName string, data map[string]interface{}) (string, []interface{}) {
	keys := []string{}
	values := []interface{}{}
	for key, value := range data {
		keys = append(keys, key)
		values = append(values, value)
	}

	queryKeys := "(" + strings.Join(keys, ", ") + ")"
	queryValues := str.GetRepeated("?", ", ", len(values))

	query := "INSERT INTO " + tableName + " "
	query += queryKeys + " VALUES " + queryValues

	return query, values
}
