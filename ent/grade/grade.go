// Code generated by ent, DO NOT EDIT.

package grade

const (
	// Label holds the string label denoting the grade type in the database.
	Label = "grade"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeClasses holds the string denoting the classes edge name in mutations.
	EdgeClasses = "classes"
	// Table holds the table name of the grade in the database.
	Table = "grades"
	// ClassesTable is the table that holds the classes relation/edge.
	ClassesTable = "classes"
	// ClassesInverseTable is the table name for the Class entity.
	// It exists in this package in order to avoid circular dependency with the "class" package.
	ClassesInverseTable = "classes"
	// ClassesColumn is the table column denoting the classes relation/edge.
	ClassesColumn = "grade_classes"
)

// Columns holds all SQL columns for grade fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
