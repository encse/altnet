// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/encse/altnet/ent/host"
	"github.com/encse/altnet/ent/schema"
)

// Host is the model entity for the Host schema.
type Host struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name schema.HostName `json:"name,omitempty"`
	// Type holds the value of the "type" field.
	Type host.Type `json:"type,omitempty"`
	// Entry holds the value of the "entry" field.
	Entry string `json:"entry,omitempty"`
	// MachineType holds the value of the "machine_type" field.
	MachineType string `json:"machine_type,omitempty"`
	// Organization holds the value of the "organization" field.
	Organization string `json:"organization,omitempty"`
	// Contact holds the value of the "contact" field.
	Contact string `json:"contact,omitempty"`
	// ContactAddress holds the value of the "contact_address" field.
	ContactAddress string `json:"contact_address,omitempty"`
	// Country holds the value of the "country" field.
	Country string `json:"country,omitempty"`
	// Location holds the value of the "location" field.
	Location string `json:"location,omitempty"`
	// GeoLocation holds the value of the "geo_location" field.
	GeoLocation string `json:"geo_location,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone []schema.PhoneNumber `json:"phone,omitempty"`
	// Neighbours holds the value of the "neighbours" field.
	Neighbours []schema.HostName `json:"neighbours,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the HostQuery when eager-loading is set.
	Edges HostEdges `json:"edges"`
}

// HostEdges holds the relations/edges for other nodes in the graph.
type HostEdges struct {
	// Virtualusers holds the value of the virtualusers edge.
	Virtualusers []*VirtualUser `json:"virtualusers,omitempty"`
	// Hackers holds the value of the hackers edge.
	Hackers []*User `json:"hackers,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// VirtualusersOrErr returns the Virtualusers value or an error if the edge
// was not loaded in eager-loading.
func (e HostEdges) VirtualusersOrErr() ([]*VirtualUser, error) {
	if e.loadedTypes[0] {
		return e.Virtualusers, nil
	}
	return nil, &NotLoadedError{edge: "virtualusers"}
}

// HackersOrErr returns the Hackers value or an error if the edge
// was not loaded in eager-loading.
func (e HostEdges) HackersOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Hackers, nil
	}
	return nil, &NotLoadedError{edge: "hackers"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Host) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case host.FieldPhone, host.FieldNeighbours:
			values[i] = new([]byte)
		case host.FieldID:
			values[i] = new(sql.NullInt64)
		case host.FieldName, host.FieldType, host.FieldEntry, host.FieldMachineType, host.FieldOrganization, host.FieldContact, host.FieldContactAddress, host.FieldCountry, host.FieldLocation, host.FieldGeoLocation:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Host", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Host fields.
func (h *Host) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case host.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			h.ID = int(value.Int64)
		case host.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				h.Name = schema.HostName(value.String)
			}
		case host.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				h.Type = host.Type(value.String)
			}
		case host.FieldEntry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field entry", values[i])
			} else if value.Valid {
				h.Entry = value.String
			}
		case host.FieldMachineType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field machine_type", values[i])
			} else if value.Valid {
				h.MachineType = value.String
			}
		case host.FieldOrganization:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field organization", values[i])
			} else if value.Valid {
				h.Organization = value.String
			}
		case host.FieldContact:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field contact", values[i])
			} else if value.Valid {
				h.Contact = value.String
			}
		case host.FieldContactAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field contact_address", values[i])
			} else if value.Valid {
				h.ContactAddress = value.String
			}
		case host.FieldCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country", values[i])
			} else if value.Valid {
				h.Country = value.String
			}
		case host.FieldLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field location", values[i])
			} else if value.Valid {
				h.Location = value.String
			}
		case host.FieldGeoLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field geo_location", values[i])
			} else if value.Valid {
				h.GeoLocation = value.String
			}
		case host.FieldPhone:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &h.Phone); err != nil {
					return fmt.Errorf("unmarshal field phone: %w", err)
				}
			}
		case host.FieldNeighbours:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field neighbours", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &h.Neighbours); err != nil {
					return fmt.Errorf("unmarshal field neighbours: %w", err)
				}
			}
		}
	}
	return nil
}

// QueryVirtualusers queries the "virtualusers" edge of the Host entity.
func (h *Host) QueryVirtualusers() *VirtualUserQuery {
	return NewHostClient(h.config).QueryVirtualusers(h)
}

// QueryHackers queries the "hackers" edge of the Host entity.
func (h *Host) QueryHackers() *UserQuery {
	return NewHostClient(h.config).QueryHackers(h)
}

// Update returns a builder for updating this Host.
// Note that you need to call Host.Unwrap() before calling this method if this Host
// was returned from a transaction, and the transaction was committed or rolled back.
func (h *Host) Update() *HostUpdateOne {
	return NewHostClient(h.config).UpdateOne(h)
}

// Unwrap unwraps the Host entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (h *Host) Unwrap() *Host {
	_tx, ok := h.config.driver.(*txDriver)
	if !ok {
		panic("ent: Host is not a transactional entity")
	}
	h.config.driver = _tx.drv
	return h
}

// String implements the fmt.Stringer.
func (h *Host) String() string {
	var builder strings.Builder
	builder.WriteString("Host(")
	builder.WriteString(fmt.Sprintf("id=%v, ", h.ID))
	builder.WriteString("name=")
	builder.WriteString(fmt.Sprintf("%v", h.Name))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", h.Type))
	builder.WriteString(", ")
	builder.WriteString("entry=")
	builder.WriteString(h.Entry)
	builder.WriteString(", ")
	builder.WriteString("machine_type=")
	builder.WriteString(h.MachineType)
	builder.WriteString(", ")
	builder.WriteString("organization=")
	builder.WriteString(h.Organization)
	builder.WriteString(", ")
	builder.WriteString("contact=")
	builder.WriteString(h.Contact)
	builder.WriteString(", ")
	builder.WriteString("contact_address=")
	builder.WriteString(h.ContactAddress)
	builder.WriteString(", ")
	builder.WriteString("country=")
	builder.WriteString(h.Country)
	builder.WriteString(", ")
	builder.WriteString("location=")
	builder.WriteString(h.Location)
	builder.WriteString(", ")
	builder.WriteString("geo_location=")
	builder.WriteString(h.GeoLocation)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(fmt.Sprintf("%v", h.Phone))
	builder.WriteString(", ")
	builder.WriteString("neighbours=")
	builder.WriteString(fmt.Sprintf("%v", h.Neighbours))
	builder.WriteByte(')')
	return builder.String()
}

// Hosts is a parsable slice of Host.
type Hosts []*Host
