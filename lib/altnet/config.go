package altnet

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/encse/altnet/ent/schema"
)

type key string

const hostKey key = "ALTNET_HOST"
const userKey key = "ALTNET_USER"
const realUserKey key = "ALTNET_REAL_USER"
const exeKey key = "ALTNET_EXE"
const sessionKey key = "ALTNET_SESSION"

func ContextToEnv(env []string, ctx context.Context) []string {
	user, err := GetUser(ctx)
	if err == nil {
		env = append(env, fmt.Sprintf("%v=%v", userKey, user))
	}

	realUser, err := GetRealUser(ctx)
	if err == nil {
		env = append(env, fmt.Sprintf("%v=%v", realUserKey, realUser))
	}

	host, err := GetHost(ctx)
	if err == nil {
		env = append(env, fmt.Sprintf("%v=%v", hostKey, host))
	}

	return env
}

func ContextFromEnv(ctx context.Context) context.Context {
	if user, ok := os.LookupEnv(string(userKey)); ok {
		ctx = SetUser(ctx, schema.Uname(user))
	}
	if host, ok := os.LookupEnv(string(hostKey)); ok {
		ctx = SetHost(ctx, schema.HostName(host))
	}
	if realUser, ok := os.LookupEnv(string(realUserKey)); ok {
		ctx = SetRealUser(ctx, schema.Uname(realUser))
	}
	return ctx
}

func GetHost(ctx context.Context) (schema.HostName, error) {
	res := ctx.Value(hostKey)
	if res == nil {
		return "", errors.New("host cannot be found")
	}
	return res.(schema.HostName), nil
}

func GetUser(ctx context.Context) (schema.Uname, error) {
	res := ctx.Value(userKey)
	if res == nil {
		return "", errors.New("user cannot be found")
	}
	return res.(schema.Uname), nil
}

func GetRealUser(ctx context.Context) (schema.Uname, error) {
	res := ctx.Value(realUserKey)
	if res == nil {
		return "", errors.New("user cannot be found")
	}
	return res.(schema.Uname), nil
}

func SetRealUser(ctx context.Context, user schema.Uname) context.Context {
	return context.WithValue(ctx, realUserKey, user)
}

func SetUser(ctx context.Context, user schema.Uname) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func SetHost(ctx context.Context, host schema.HostName) context.Context {
	return context.WithValue(ctx, hostKey, host)
}
