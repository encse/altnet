package altnet

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/encse/altnet/lib/uumap"
)

type key string

const hostKey key = "ALTNET_HOST"
const userKey key = "ALTNET_USER"
const exeKey key = "ALTNET_EXE"
const sessionKey key = "ALTNET_SESSION"

type User string

func ContextToEnv(env []string, ctx context.Context) []string {
	user, err := GetUser(ctx)
	if err == nil {
		env = append(env, fmt.Sprintf("%v=%v", userKey, user))
	}

	host, err := GetHost(ctx)
	if err == nil {
		env = append(env, fmt.Sprintf("%v=%v", hostKey, host))
	}
	return env
}

func ContextFromEnv(ctx context.Context) context.Context {
	if user, ok := os.LookupEnv(string(userKey)); ok {
		ctx = SetUser(ctx, User(user))
	}
	if host, ok := os.LookupEnv(string(hostKey)); ok {
		ctx = SetHost(ctx, uumap.Host(host))
	}
	return ctx
}

func GetHost(ctx context.Context) (uumap.Host, error) {
	res := ctx.Value(hostKey)
	if res == nil {
		return "", errors.New("Host cannot be found")
	}
	return res.(uumap.Host), nil
}

func GetUser(ctx context.Context) (User, error) {
	res := ctx.Value(userKey)
	if res == nil {
		return "", errors.New("User cannot be found")
	}
	return res.(User), nil
}

func SetUser(ctx context.Context, user User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func SetHost(ctx context.Context, host uumap.Host) context.Context {
	return context.WithValue(ctx, hostKey, host)
}
