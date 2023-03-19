package uumap

import (
	"context"
	"strings"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"github.com/encse/altnet/ent"
	"github.com/encse/altnet/ent/host"
	"github.com/encse/altnet/schema"
)

func (n Network) FindPhoneNumbersWithPrefix(ctx context.Context, prefix string) []schema.PhoneNumber {

	var vs []struct {
		Phone []schema.PhoneNumber
	}

	err := n.Client.Host.
		Query().
		Select(host.FieldPhone).
		Scan(ctx, &vs)

	if err != nil {
		return nil
	}

	res := make([]schema.PhoneNumber, 0)
	for _, v := range vs {
		for _, phone := range v.Phone {
			if strings.HasPrefix(string(phone), prefix) {
				res = append(res, schema.PhoneNumber(phone))
			}
		}
	}
	return res
}

// Lookup checks the number in the phonebook and returns with a hostname if found.
// If the host doesn't need an extension but an extension is provided in the phone number
// we return with the host regardless. However if the host requires an extension and no
// extension is provided in the phone number we return with failure.
// This is analogous to dialing a number: if the extension is not needed, it is simply
// ignored.
func (n Network) LookupHostByPhone(ctx context.Context, phoneNumber schema.PhoneNumber) (*ent.Host, error) {
	host, err := n.lookupHostByPhoneI(ctx, phoneNumber)
	if err != nil {
		return nil, err
	}
	if host != nil {
		return host, nil
	}

	withoutExt, err := schema.ParsePhoneNumberSkipExtension(string(phoneNumber))
	if err != nil {
		return nil, nil
	}

	return n.lookupHostByPhoneI(ctx, schema.PhoneNumber(withoutExt))
}

func (n Network) lookupHostByPhoneI(ctx context.Context, phoneNumber schema.PhoneNumber) (*ent.Host, error) {
	hosts, err := n.Client.Host.
		Query().
		Where(func(s *sql.Selector) {
			s.Where(sqljson.ValueContains(host.FieldPhone, phoneNumber))
		}).All(ctx)

	if err != nil {
		return nil, err
	}

	if len(hosts) == 0 {
		return nil, nil
	}

	return hosts[0], nil
}
