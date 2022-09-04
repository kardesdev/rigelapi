// Code generated by ent, DO NOT EDIT.

package period

const (
	// Label holds the string label denoting the period type in the database.
	Label = "period"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeClassPeriods holds the string denoting the classperiods edge name in mutations.
	EdgeClassPeriods = "classPeriods"
	// EdgeYear holds the string denoting the year edge name in mutations.
	EdgeYear = "year"
	// Table holds the table name of the period in the database.
	Table = "periods"
	// ClassPeriodsTable is the table that holds the classPeriods relation/edge.
	ClassPeriodsTable = "class_periods"
	// ClassPeriodsInverseTable is the table name for the ClassPeriod entity.
	// It exists in this package in order to avoid circular dependency with the "classperiod" package.
	ClassPeriodsInverseTable = "class_periods"
	// ClassPeriodsColumn is the table column denoting the classPeriods relation/edge.
	ClassPeriodsColumn = "period_class_periods"
	// YearTable is the table that holds the year relation/edge.
	YearTable = "periods"
	// YearInverseTable is the table name for the Year entity.
	// It exists in this package in order to avoid circular dependency with the "year" package.
	YearInverseTable = "years"
	// YearColumn is the table column denoting the year relation/edge.
	YearColumn = "year_periods"
)

// Columns holds all SQL columns for period fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "periods"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"year_periods",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
