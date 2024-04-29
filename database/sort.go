package database

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

func applySort(tx *gorm.DB, req *PageRequest, sortMap ...SortMap) *gorm.DB {
	sortBy := req.SortBy
	if sortBy == "" {
		sortBy = defaultSortBy
	}

	if parsed := parseSort(sortBy, sortMap...); parsed != "" {
		tx = tx.Order(parsed)
	}

	return tx
}

func parseSort(s string, sortMap ...SortMap) string {
	order := Ascending
	if len(s) > 0 && s[0] == '-' {
		order = Descending
		s = s[1:]
	}

	sortingMap := defaultSortMap
	if len(sortMap) > 0 {
		sortingMap = sortMap[0]
	}

	var columnName string
	if cn, ok := sortingMap[s]; ok {
		columnName = cn
	}

	if columnName == "" {
		return ""
	}

	return fmt.Sprintf("%s %s", strings.ToLower(columnName), strings.ToUpper(string(order)))
}
