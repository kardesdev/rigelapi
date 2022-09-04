// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/vmkevv/rigelapi/ent/class"
	"github.com/vmkevv/rigelapi/ent/classperiod"
	"github.com/vmkevv/rigelapi/ent/period"
)

// ClassPeriod is the model entity for the ClassPeriod schema.
type ClassPeriod struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Start holds the value of the "start" field.
	Start time.Time `json:"start,omitempty"`
	// End holds the value of the "end" field.
	End time.Time `json:"end,omitempty"`
	// Finished holds the value of the "finished" field.
	Finished bool `json:"finished,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ClassPeriodQuery when eager-loading is set.
	Edges                ClassPeriodEdges `json:"edges"`
	class_class_periods  *string
	period_class_periods *string
}

// ClassPeriodEdges holds the relations/edges for other nodes in the graph.
type ClassPeriodEdges struct {
	// Attendances holds the value of the attendances edge.
	Attendances []*Attendance `json:"attendances,omitempty"`
	// AttendanceSyncs holds the value of the attendanceSyncs edge.
	AttendanceSyncs []*AttendanceSync `json:"attendanceSyncs,omitempty"`
	// Activities holds the value of the activities edge.
	Activities []*Activity `json:"activities,omitempty"`
	// ActivitySyncs holds the value of the activitySyncs edge.
	ActivitySyncs []*ActivitySync `json:"activitySyncs,omitempty"`
	// Class holds the value of the class edge.
	Class *Class `json:"class,omitempty"`
	// Period holds the value of the period edge.
	Period *Period `json:"period,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// AttendancesOrErr returns the Attendances value or an error if the edge
// was not loaded in eager-loading.
func (e ClassPeriodEdges) AttendancesOrErr() ([]*Attendance, error) {
	if e.loadedTypes[0] {
		return e.Attendances, nil
	}
	return nil, &NotLoadedError{edge: "attendances"}
}

// AttendanceSyncsOrErr returns the AttendanceSyncs value or an error if the edge
// was not loaded in eager-loading.
func (e ClassPeriodEdges) AttendanceSyncsOrErr() ([]*AttendanceSync, error) {
	if e.loadedTypes[1] {
		return e.AttendanceSyncs, nil
	}
	return nil, &NotLoadedError{edge: "attendanceSyncs"}
}

// ActivitiesOrErr returns the Activities value or an error if the edge
// was not loaded in eager-loading.
func (e ClassPeriodEdges) ActivitiesOrErr() ([]*Activity, error) {
	if e.loadedTypes[2] {
		return e.Activities, nil
	}
	return nil, &NotLoadedError{edge: "activities"}
}

// ActivitySyncsOrErr returns the ActivitySyncs value or an error if the edge
// was not loaded in eager-loading.
func (e ClassPeriodEdges) ActivitySyncsOrErr() ([]*ActivitySync, error) {
	if e.loadedTypes[3] {
		return e.ActivitySyncs, nil
	}
	return nil, &NotLoadedError{edge: "activitySyncs"}
}

// ClassOrErr returns the Class value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ClassPeriodEdges) ClassOrErr() (*Class, error) {
	if e.loadedTypes[4] {
		if e.Class == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: class.Label}
		}
		return e.Class, nil
	}
	return nil, &NotLoadedError{edge: "class"}
}

// PeriodOrErr returns the Period value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ClassPeriodEdges) PeriodOrErr() (*Period, error) {
	if e.loadedTypes[5] {
		if e.Period == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: period.Label}
		}
		return e.Period, nil
	}
	return nil, &NotLoadedError{edge: "period"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ClassPeriod) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case classperiod.FieldFinished:
			values[i] = new(sql.NullBool)
		case classperiod.FieldID:
			values[i] = new(sql.NullString)
		case classperiod.FieldStart, classperiod.FieldEnd:
			values[i] = new(sql.NullTime)
		case classperiod.ForeignKeys[0]: // class_class_periods
			values[i] = new(sql.NullString)
		case classperiod.ForeignKeys[1]: // period_class_periods
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ClassPeriod", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ClassPeriod fields.
func (cp *ClassPeriod) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case classperiod.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				cp.ID = value.String
			}
		case classperiod.FieldStart:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start", values[i])
			} else if value.Valid {
				cp.Start = value.Time
			}
		case classperiod.FieldEnd:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end", values[i])
			} else if value.Valid {
				cp.End = value.Time
			}
		case classperiod.FieldFinished:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field finished", values[i])
			} else if value.Valid {
				cp.Finished = value.Bool
			}
		case classperiod.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field class_class_periods", values[i])
			} else if value.Valid {
				cp.class_class_periods = new(string)
				*cp.class_class_periods = value.String
			}
		case classperiod.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field period_class_periods", values[i])
			} else if value.Valid {
				cp.period_class_periods = new(string)
				*cp.period_class_periods = value.String
			}
		}
	}
	return nil
}

// QueryAttendances queries the "attendances" edge of the ClassPeriod entity.
func (cp *ClassPeriod) QueryAttendances() *AttendanceQuery {
	return (&ClassPeriodClient{config: cp.config}).QueryAttendances(cp)
}

// QueryAttendanceSyncs queries the "attendanceSyncs" edge of the ClassPeriod entity.
func (cp *ClassPeriod) QueryAttendanceSyncs() *AttendanceSyncQuery {
	return (&ClassPeriodClient{config: cp.config}).QueryAttendanceSyncs(cp)
}

// QueryActivities queries the "activities" edge of the ClassPeriod entity.
func (cp *ClassPeriod) QueryActivities() *ActivityQuery {
	return (&ClassPeriodClient{config: cp.config}).QueryActivities(cp)
}

// QueryActivitySyncs queries the "activitySyncs" edge of the ClassPeriod entity.
func (cp *ClassPeriod) QueryActivitySyncs() *ActivitySyncQuery {
	return (&ClassPeriodClient{config: cp.config}).QueryActivitySyncs(cp)
}

// QueryClass queries the "class" edge of the ClassPeriod entity.
func (cp *ClassPeriod) QueryClass() *ClassQuery {
	return (&ClassPeriodClient{config: cp.config}).QueryClass(cp)
}

// QueryPeriod queries the "period" edge of the ClassPeriod entity.
func (cp *ClassPeriod) QueryPeriod() *PeriodQuery {
	return (&ClassPeriodClient{config: cp.config}).QueryPeriod(cp)
}

// Update returns a builder for updating this ClassPeriod.
// Note that you need to call ClassPeriod.Unwrap() before calling this method if this ClassPeriod
// was returned from a transaction, and the transaction was committed or rolled back.
func (cp *ClassPeriod) Update() *ClassPeriodUpdateOne {
	return (&ClassPeriodClient{config: cp.config}).UpdateOne(cp)
}

// Unwrap unwraps the ClassPeriod entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cp *ClassPeriod) Unwrap() *ClassPeriod {
	_tx, ok := cp.config.driver.(*txDriver)
	if !ok {
		panic("ent: ClassPeriod is not a transactional entity")
	}
	cp.config.driver = _tx.drv
	return cp
}

// String implements the fmt.Stringer.
func (cp *ClassPeriod) String() string {
	var builder strings.Builder
	builder.WriteString("ClassPeriod(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cp.ID))
	builder.WriteString("start=")
	builder.WriteString(cp.Start.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("end=")
	builder.WriteString(cp.End.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("finished=")
	builder.WriteString(fmt.Sprintf("%v", cp.Finished))
	builder.WriteByte(')')
	return builder.String()
}

// ClassPeriods is a parsable slice of ClassPeriod.
type ClassPeriods []*ClassPeriod

func (cp ClassPeriods) config(cfg config) {
	for _i := range cp {
		cp[_i].config = cfg
	}
}
