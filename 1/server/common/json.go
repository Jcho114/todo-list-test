package common

import (
	"database/sql"
	"encoding/json"
)

func RowsToJson(rows *sql.Rows) string {
	columns, err := rows.Columns()
	FailOnError(err, "Failed to retrieve query rows")
	count := len(columns)

	res := make([]interface{}, 0)

	values := make([]interface{}, count)
	valuesPointers := make([]interface{}, count)
	for i := 0; i < count; i++ {
		valuesPointers[i] = &values[i]
	}

	for rows.Next() {
		rows.Scan(valuesPointers...)
		entry := make(map[string]interface{})
		for i, column := range columns {
			entry[column] = values[i]
		}
		res = append(res, entry)
	}

	jsonEncoding, err := json.Marshal(res)
	FailOnError(err, "Failed to encode to JSON")

	return string(jsonEncoding)
}
