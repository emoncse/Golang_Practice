package utils

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

// ConvertRowsToJSON converts SQL rows into a slice of maps for easy JSON conversion
func ConvertRowsToJSON(rows *sql.Rows) ([]map[string]interface{}, error) {
	columns, err := rows.Columns() // Get the column names
	if err != nil {
		return nil, err
	}

	results := []map[string]interface{}{} // Prepare to hold the result set

	// Loop through rows
	for rows.Next() {
		// Prepare a slice to hold each column value
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))

		// Set up pointers to each value (to store the row values)
		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		// Scan the row into the value pointers
		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, err
		}

		// Prepare a map to hold the column-value pair for each row
		rowMap := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]

			// Handle type assertion of column values
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}

			rowMap[col] = v
		}

		// Append the row to the result set
		results = append(results, rowMap)
	}

	return results, nil
}

// JSONResponse sends a JSON response with the provided status code and data
func JSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
