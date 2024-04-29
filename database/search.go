package database

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
)

type SearchArgs struct {
	Column string
	Value  string
}

func AddSearchToQuery(tx *gorm.DB, args []SearchArgs) *gorm.DB {
	if args != nil && len(args) > 0 {
		expressions := make([]clause.Expression, len(args))
		for idx, argument := range args {
			column := argument.Column
			if !strings.Contains(column, ".") {
				column = fmt.Sprintf("%s.%s", tx.Statement.Table, column)
			}
			searchQuery := fmt.Sprintf("LOWER(%s) LIKE ?", column)
			condition := tx.Statement.BuildCondition(searchQuery, fmt.Sprintf("%%%s%%", argument.Value))[0]
			expressions[idx] = clause.Where{Exprs: []clause.Expression{clause.Or(condition)}}
		}

		tx.Clauses(clause.Where{Exprs: []clause.Expression{clause.And(clause.Or(expressions...))}})
	}
	return tx
}
