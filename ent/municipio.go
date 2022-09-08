// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/vmkevv/rigelapi/ent/municipio"
	"github.com/vmkevv/rigelapi/ent/provincia"
)

// Municipio is the model entity for the Municipio schema.
type Municipio struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MunicipioQuery when eager-loading is set.
	Edges                MunicipioEdges `json:"edges"`
	provincia_municipios *string
}

// MunicipioEdges holds the relations/edges for other nodes in the graph.
type MunicipioEdges struct {
	// Schools holds the value of the schools edge.
	Schools []*School `json:"schools,omitempty"`
	// Provincia holds the value of the provincia edge.
	Provincia *Provincia `json:"provincia,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// SchoolsOrErr returns the Schools value or an error if the edge
// was not loaded in eager-loading.
func (e MunicipioEdges) SchoolsOrErr() ([]*School, error) {
	if e.loadedTypes[0] {
		return e.Schools, nil
	}
	return nil, &NotLoadedError{edge: "schools"}
}

// ProvinciaOrErr returns the Provincia value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MunicipioEdges) ProvinciaOrErr() (*Provincia, error) {
	if e.loadedTypes[1] {
		if e.Provincia == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: provincia.Label}
		}
		return e.Provincia, nil
	}
	return nil, &NotLoadedError{edge: "provincia"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Municipio) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case municipio.FieldID, municipio.FieldName:
			values[i] = new(sql.NullString)
		case municipio.ForeignKeys[0]: // provincia_municipios
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Municipio", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Municipio fields.
func (m *Municipio) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case municipio.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				m.ID = value.String
			}
		case municipio.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				m.Name = value.String
			}
		case municipio.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field provincia_municipios", values[i])
			} else if value.Valid {
				m.provincia_municipios = new(string)
				*m.provincia_municipios = value.String
			}
		}
	}
	return nil
}

// QuerySchools queries the "schools" edge of the Municipio entity.
func (m *Municipio) QuerySchools() *SchoolQuery {
	return (&MunicipioClient{config: m.config}).QuerySchools(m)
}

// QueryProvincia queries the "provincia" edge of the Municipio entity.
func (m *Municipio) QueryProvincia() *ProvinciaQuery {
	return (&MunicipioClient{config: m.config}).QueryProvincia(m)
}

// Update returns a builder for updating this Municipio.
// Note that you need to call Municipio.Unwrap() before calling this method if this Municipio
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Municipio) Update() *MunicipioUpdateOne {
	return (&MunicipioClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the Municipio entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Municipio) Unwrap() *Municipio {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Municipio is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Municipio) String() string {
	var builder strings.Builder
	builder.WriteString("Municipio(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("name=")
	builder.WriteString(m.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Municipios is a parsable slice of Municipio.
type Municipios []*Municipio

func (m Municipios) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}