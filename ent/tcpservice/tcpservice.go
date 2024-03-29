// Code generated by ent, DO NOT EDIT.

package tcpservice

const (
	// Label holds the string label denoting the tcpservice type in the database.
	Label = "tcp_service"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPort holds the string denoting the port field in the database.
	FieldPort = "port"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeHosts holds the string denoting the hosts edge name in mutations.
	EdgeHosts = "hosts"
	// Table holds the table name of the tcpservice in the database.
	Table = "tcp_services"
	// HostsTable is the table that holds the hosts relation/edge. The primary key declared below.
	HostsTable = "host_services"
	// HostsInverseTable is the table name for the Host entity.
	// It exists in this package in order to avoid circular dependency with the "host" package.
	HostsInverseTable = "hosts"
)

// Columns holds all SQL columns for tcpservice fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldPort,
	FieldDescription,
}

var (
	// HostsPrimaryKey and HostsColumn2 are the table columns denoting the
	// primary key for the hosts relation (M2M).
	HostsPrimaryKey = []string{"host_id", "tcp_service_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
