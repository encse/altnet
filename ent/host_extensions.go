// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/encse/altnet/ent/user"
	"github.com/encse/altnet/schema"
)

func (h *Host) IsHacked(ctx context.Context, userName schema.Uname) (bool, error) {
	res, err := h.QueryHackers().Where(user.UserEQ(userName)).Count(ctx)
	return res > 0, err
}
