package mirgation

type ColumnDescription struct {
	Name              string
	Type              string
	DefaultType       string
	DefaultExpression string
	Comment           string
	CodecExpression   string
	TTLExpression     string
}
