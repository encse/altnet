package altnet

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/encse/altnet/schema"
)

type key string

const (
	userKey     key = "ALTNET_USER"
	hostKey     key = "ALTNET_HOST"
	fromHostKey key = "ALTNET_FROM_HOST"
	fromUserKey key = "ALTNET_FROM_USER"
	realUserKey key = "ALTNET_REAL_USER"
	exeKey      key = "ALTNET_EXE"
	sessionKey  key = "ALTNET_SESSION"
)

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

	fromUser, err := GetFromUser(ctx)
	if err == nil {
		env = append(env, fmt.Sprintf("%v=%v", fromUserKey, fromUser))
	}

	fromHost, err := GetFromHost(ctx)
	if err == nil {
		env = append(env, fmt.Sprintf("%v=%v", fromHostKey, fromHost))
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

	if fromUser, ok := os.LookupEnv(string(fromUserKey)); ok {
		ctx = SetFromUser(ctx, schema.Uname(fromUser))
	}
	if fromHost, ok := os.LookupEnv(string(fromHostKey)); ok {
		ctx = SetFromHost(ctx, schema.HostName(fromHost))
	}
	return ctx
}

func GetFromHost(ctx context.Context) (schema.HostName, error) {
	res := ctx.Value(fromHostKey)
	if res == nil {
		return "", errors.New("fromhost cannot be found")
	}
	return res.(schema.HostName), nil
}

func GetFromUser(ctx context.Context) (schema.Uname, error) {
	res := ctx.Value(fromUserKey)
	if res == nil {
		return "", errors.New("fromuser cannot be found")
	}
	return res.(schema.Uname), nil
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

func SetFromUser(ctx context.Context, user schema.Uname) context.Context {
	return context.WithValue(ctx, fromUserKey, user)
}

func SetUser(ctx context.Context, user schema.Uname) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func SetFromHost(ctx context.Context, host schema.HostName) context.Context {
	return context.WithValue(ctx, fromHostKey, host)
}

func SetHost(ctx context.Context, host schema.HostName) context.Context {
	return context.WithValue(ctx, hostKey, host)
}

func EnterHost(ctx context.Context, host schema.HostName, uname schema.Uname) context.Context {
	ctx = context.WithValue(ctx, fromHostKey, ctx.Value(hostKey))
	ctx = context.WithValue(ctx, fromUserKey, ctx.Value(userKey))
	ctx = context.WithValue(ctx, hostKey, host)
	ctx = context.WithValue(ctx, userKey, uname)
	return ctx
}

func ReverseConnection(ctx context.Context) context.Context {
	fromHost := ctx.Value(fromHostKey)
	fromUser := ctx.Value(fromUserKey)

	ctx = context.WithValue(ctx, fromHostKey, ctx.Value(hostKey))
	ctx = context.WithValue(ctx, fromUserKey, ctx.Value(userKey))
	ctx = context.WithValue(ctx, hostKey, fromHost)
	ctx = context.WithValue(ctx, userKey, fromUser)
	return ctx
}
