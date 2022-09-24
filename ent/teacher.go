// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/vmkevv/rigelapi/ent/teacher"
)

// Teacher is the model entity for the Teacher schema.
type Teacher struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TeacherQuery when eager-loading is set.
	Edges TeacherEdges `json:"edges"`
}

// TeacherEdges holds the relations/edges for other nodes in the graph.
type TeacherEdges struct {
	// Classes holds the value of the classes edge.
	Classes []*Class `json:"classes,omitempty"`
	// StudentSyncs holds the value of the studentSyncs edge.
	StudentSyncs []*StudentSync `json:"studentSyncs,omitempty"`
	// ActivitySyncs holds the value of the activitySyncs edge.
	ActivitySyncs []*ActivitySync `json:"activitySyncs,omitempty"`
	// ClassPeriodSyncs holds the value of the classPeriodSyncs edge.
	ClassPeriodSyncs []*ClassPeriodSync `json:"classPeriodSyncs,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// ClassesOrErr returns the Classes value or an error if the edge
// was not loaded in eager-loading.
func (e TeacherEdges) ClassesOrErr() ([]*Class, error) {
	if e.loadedTypes[0] {
		return e.Classes, nil
	}
	return nil, &NotLoadedError{edge: "classes"}
}

// StudentSyncsOrErr returns the StudentSyncs value or an error if the edge
// was not loaded in eager-loading.
func (e TeacherEdges) StudentSyncsOrErr() ([]*StudentSync, error) {
	if e.loadedTypes[1] {
		return e.StudentSyncs, nil
	}
	return nil, &NotLoadedError{edge: "studentSyncs"}
}

// ActivitySyncsOrErr returns the ActivitySyncs value or an error if the edge
// was not loaded in eager-loading.
func (e TeacherEdges) ActivitySyncsOrErr() ([]*ActivitySync, error) {
	if e.loadedTypes[2] {
		return e.ActivitySyncs, nil
	}
	return nil, &NotLoadedError{edge: "activitySyncs"}
}

// ClassPeriodSyncsOrErr returns the ClassPeriodSyncs value or an error if the edge
// was not loaded in eager-loading.
func (e TeacherEdges) ClassPeriodSyncsOrErr() ([]*ClassPeriodSync, error) {
	if e.loadedTypes[3] {
		return e.ClassPeriodSyncs, nil
	}
	return nil, &NotLoadedError{edge: "classPeriodSyncs"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Teacher) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case teacher.FieldID, teacher.FieldName, teacher.FieldLastName, teacher.FieldEmail, teacher.FieldPassword:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Teacher", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Teacher fields.
func (t *Teacher) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case teacher.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				t.ID = value.String
			}
		case teacher.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case teacher.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				t.LastName = value.String
			}
		case teacher.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				t.Email = value.String
			}
		case teacher.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				t.Password = value.String
			}
		}
	}
	return nil
}

// QueryClasses queries the "classes" edge of the Teacher entity.
func (t *Teacher) QueryClasses() *ClassQuery {
	return (&TeacherClient{config: t.config}).QueryClasses(t)
}

// QueryStudentSyncs queries the "studentSyncs" edge of the Teacher entity.
func (t *Teacher) QueryStudentSyncs() *StudentSyncQuery {
	return (&TeacherClient{config: t.config}).QueryStudentSyncs(t)
}

// QueryActivitySyncs queries the "activitySyncs" edge of the Teacher entity.
func (t *Teacher) QueryActivitySyncs() *ActivitySyncQuery {
	return (&TeacherClient{config: t.config}).QueryActivitySyncs(t)
}

// QueryClassPeriodSyncs queries the "classPeriodSyncs" edge of the Teacher entity.
func (t *Teacher) QueryClassPeriodSyncs() *ClassPeriodSyncQuery {
	return (&TeacherClient{config: t.config}).QueryClassPeriodSyncs(t)
}

// Update returns a builder for updating this Teacher.
// Note that you need to call Teacher.Unwrap() before calling this method if this Teacher
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Teacher) Update() *TeacherUpdateOne {
	return (&TeacherClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Teacher entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Teacher) Unwrap() *Teacher {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Teacher is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Teacher) String() string {
	var builder strings.Builder
	builder.WriteString("Teacher(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("name=")
	builder.WriteString(t.Name)
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(t.LastName)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(t.Email)
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(t.Password)
	builder.WriteByte(')')
	return builder.String()
}

// Teachers is a parsable slice of Teacher.
type Teachers []*Teacher

func (t Teachers) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
