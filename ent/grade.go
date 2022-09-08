// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/vmkevv/rigelapi/ent/grade"
)

// Grade is the model entity for the Grade schema.
type Grade struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GradeQuery when eager-loading is set.
	Edges GradeEdges `json:"edges"`
}

// GradeEdges holds the relations/edges for other nodes in the graph.
type GradeEdges struct {
	// Classes holds the value of the classes edge.
	Classes []*Class `json:"classes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ClassesOrErr returns the Classes value or an error if the edge
// was not loaded in eager-loading.
func (e GradeEdges) ClassesOrErr() ([]*Class, error) {
	if e.loadedTypes[0] {
		return e.Classes, nil
	}
	return nil, &NotLoadedError{edge: "classes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Grade) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case grade.FieldID, grade.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Grade", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Grade fields.
func (gr *Grade) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case grade.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				gr.ID = value.String
			}
		case grade.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				gr.Name = value.String
			}
		}
	}
	return nil
}

// QueryClasses queries the "classes" edge of the Grade entity.
func (gr *Grade) QueryClasses() *ClassQuery {
	return (&GradeClient{config: gr.config}).QueryClasses(gr)
}

// Update returns a builder for updating this Grade.
// Note that you need to call Grade.Unwrap() before calling this method if this Grade
// was returned from a transaction, and the transaction was committed or rolled back.
func (gr *Grade) Update() *GradeUpdateOne {
	return (&GradeClient{config: gr.config}).UpdateOne(gr)
}

// Unwrap unwraps the Grade entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gr *Grade) Unwrap() *Grade {
	_tx, ok := gr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Grade is not a transactional entity")
	}
	gr.config.driver = _tx.drv
	return gr
}

// String implements the fmt.Stringer.
func (gr *Grade) String() string {
	var builder strings.Builder
	builder.WriteString("Grade(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gr.ID))
	builder.WriteString("name=")
	builder.WriteString(gr.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Grades is a parsable slice of Grade.
type Grades []*Grade

func (gr Grades) config(cfg config) {
	for _i := range gr {
		gr[_i].config = cfg
	}
}