package database

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type FilterArgs struct {
	Column    string
	Operation Operation
	Value     interface{}
}

func AddFilterToQuery(tx *gorm.DB, params []FilterArgs) *gorm.DB {
	if len(params) > 0 {
		var query string
		for _, param := range params {
			column := param.Column
			if !strings.Contains(column, ".") {
				column = fmt.Sprintf("%s.%s", tx.Statement.Table, column)
			}
			query = fmt.Sprintf("%s %s ?", column, param.Operation)
			tx.Where(query, param.Value)
		}
	}
	return tx
}
