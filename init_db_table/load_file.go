package init_db_table

import (
	"os"
	"path/filepath"
)

func LoadFile(filename string) (string, error){

	path := filepath.Join("init_db_table", filename)

	sqlFile, err := os.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(sqlFile), nil
}