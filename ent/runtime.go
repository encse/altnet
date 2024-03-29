// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/encse/altnet/ent/host"
	"github.com/encse/altnet/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	hostFields := schema.Host{}.Fields()
	_ = hostFields
	// hostDescEntry is the schema descriptor for entry field.
	hostDescEntry := hostFields[2].Descriptor()
	// host.DefaultEntry holds the default value on creation for the entry field.
	host.DefaultEntry = hostDescEntry.Default.(string)
	// hostDescMachineType is the schema descriptor for machine_type field.
	hostDescMachineType := hostFields[3].Descriptor()
	// host.DefaultMachineType holds the default value on creation for the machine_type field.
	host.DefaultMachineType = hostDescMachineType.Default.(string)
	// hostDescOrganization is the schema descriptor for organization field.
	hostDescOrganization := hostFields[4].Descriptor()
	// host.DefaultOrganization holds the default value on creation for the organization field.
	host.DefaultOrganization = hostDescOrganization.Default.(string)
	// hostDescContact is the schema descriptor for contact field.
	hostDescContact := hostFields[5].Descriptor()
	// host.DefaultContact holds the default value on creation for the contact field.
	host.DefaultContact = hostDescContact.Default.(string)
	// hostDescContactAddress is the schema descriptor for contact_address field.
	hostDescContactAddress := hostFields[6].Descriptor()
	// host.DefaultContactAddress holds the default value on creation for the contact_address field.
	host.DefaultContactAddress = hostDescContactAddress.Default.(string)
	// hostDescCountry is the schema descriptor for country field.
	hostDescCountry := hostFields[7].Descriptor()
	// host.DefaultCountry holds the default value on creation for the country field.
	host.DefaultCountry = hostDescCountry.Default.(string)
	// hostDescLocation is the schema descriptor for location field.
	hostDescLocation := hostFields[8].Descriptor()
	// host.DefaultLocation holds the default value on creation for the location field.
	host.DefaultLocation = hostDescLocation.Default.(string)
	// hostDescGeoLocation is the schema descriptor for geo_location field.
	hostDescGeoLocation := hostFields[9].Descriptor()
	// host.DefaultGeoLocation holds the default value on creation for the geo_location field.
	host.DefaultGeoLocation = hostDescGeoLocation.Default.(string)
}
