package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type openCallback func(*sql.DB) error
type QueryCallback func(result map[string]interface{}) error

func openDB(filename string, callback openCallback) error {

	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return err
	}
	defer db.Close()

	return callback(db)
}

func Query(filename, q string, callback QueryCallback) error {
	return openDB(filename, func(db *sql.DB) error {
		rows, err := db.Query(q) // Note: Ignoring errors for brevity
		if err != nil {
			return err
		}

		cols, err := rows.Columns()
		if err != nil {
			return err
		}

		for rows.Next() {
			// Create a slice of interface{}'s to represent each column,
			// and a second slice to contain pointers to each item in the columns slice.
			columns := make([]interface{}, len(cols))
			columnPointers := make([]interface{}, len(cols))
			for i, _ := range columns {
				columnPointers[i] = &columns[i]
			}

			// Scan the result into the column pointers...
			if err := rows.Scan(columnPointers...); err != nil {
				return err
			}

			// Create our map, and retrieve the value for each column from the pointers slice,
			// storing it in the map with the name of the column as the key.
			m := make(map[string]interface{})
			for i, colName := range cols {
				val := columnPointers[i].(*interface{})
				m[colName] = *val
			}

			callback(m)

		}

		return nil
	})
}
