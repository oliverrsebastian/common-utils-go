package database

const TransactionCtxKey = "DBTrx"

type Operation string

const (
	Equal    Operation = "="
	NotEqual Operation = "!="

	LessThan      Operation = "<"
	LessThanEqual Operation = "<="

	GreaterThan      Operation = ">"
	GreaterThanEqual Operation = ">="

	In   Operation = "IN"
	Like Operation = "LIKE"
)

const defaultSortBy = "-createdAt"

type SortDirection string

const (
	Ascending  SortDirection = "ASC"
	Descending SortDirection = "DESC"
)

type SortMap map[string]string

var defaultSortMap = SortMap{
	"createdAt": "created_at",
}
