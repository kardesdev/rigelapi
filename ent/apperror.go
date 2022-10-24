// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/vmkevv/rigelapi/ent/apperror"
)

// AppError is the model entity for the AppError schema.
type AppError struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// Cause holds the value of the "cause" field.
	Cause string `json:"cause,omitempty"`
	// ErrorMsg holds the value of the "error_msg" field.
	ErrorMsg string `json:"error_msg,omitempty"`
	// ErrorStack holds the value of the "error_stack" field.
	ErrorStack string `json:"error_stack,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppError) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case apperror.FieldID, apperror.FieldUserID, apperror.FieldCause, apperror.FieldErrorMsg, apperror.FieldErrorStack:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AppError", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppError fields.
func (ae *AppError) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case apperror.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				ae.ID = value.String
			}
		case apperror.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ae.UserID = value.String
			}
		case apperror.FieldCause:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cause", values[i])
			} else if value.Valid {
				ae.Cause = value.String
			}
		case apperror.FieldErrorMsg:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field error_msg", values[i])
			} else if value.Valid {
				ae.ErrorMsg = value.String
			}
		case apperror.FieldErrorStack:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field error_stack", values[i])
			} else if value.Valid {
				ae.ErrorStack = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AppError.
// Note that you need to call AppError.Unwrap() before calling this method if this AppError
// was returned from a transaction, and the transaction was committed or rolled back.
func (ae *AppError) Update() *AppErrorUpdateOne {
	return (&AppErrorClient{config: ae.config}).UpdateOne(ae)
}

// Unwrap unwraps the AppError entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ae *AppError) Unwrap() *AppError {
	_tx, ok := ae.config.driver.(*txDriver)
	if !ok {
		panic("ent: AppError is not a transactional entity")
	}
	ae.config.driver = _tx.drv
	return ae
}

// String implements the fmt.Stringer.
func (ae *AppError) String() string {
	var builder strings.Builder
	builder.WriteString("AppError(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ae.ID))
	builder.WriteString("user_id=")
	builder.WriteString(ae.UserID)
	builder.WriteString(", ")
	builder.WriteString("cause=")
	builder.WriteString(ae.Cause)
	builder.WriteString(", ")
	builder.WriteString("error_msg=")
	builder.WriteString(ae.ErrorMsg)
	builder.WriteString(", ")
	builder.WriteString("error_stack=")
	builder.WriteString(ae.ErrorStack)
	builder.WriteByte(')')
	return builder.String()
}

// AppErrors is a parsable slice of AppError.
type AppErrors []*AppError

func (ae AppErrors) config(cfg config) {
	for _i := range ae {
		ae[_i].config = cfg
	}
}