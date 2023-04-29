package uman

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/encse/altnet/ent"
	"github.com/encse/altnet/ent/user"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/log"
	"github.com/encse/altnet/lib/uumap"
	"github.com/encse/altnet/schema"
	"golang.org/x/crypto/bcrypt"
)

type LoginRes struct {
	User             schema.Uname
	LastLogin        *time.Time
	LastLoginFailure *time.Time
}

func ValidatePassword(
	ctx context.Context,
	network uumap.Network,
	uname schema.Uname,
	password schema.Password,
) (bool, error) {
	u, err := network.Client.User.Query().
		Where(
			user.UserEQ(uname),
		).First(ctx)

	if ent.IsNotFound(err) {
		log.Info("User not found", uname)
		return false, nil
	}

	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil, nil
}

func LoginAttempt(
	ctx context.Context,
	network uumap.Network,
	uname schema.Uname,
	password schema.Password,
) (*LoginRes, error) {

	u, err := network.Client.User.Query().
		Where(
			user.UserEQ(uname),
		).First(ctx)

	if ent.IsNotFound(err) {
		log.Info("User not found", uname)
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	now := time.Now()
	lastLogin := u.LastLogin
	lastLoginAttempt := u.LastLoginAttempt

	err = u.Update().SetLastLoginAttempt(now).Exec(ctx)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		log.Info("Invalid password for", uname)
		return nil, nil // return nil error here
	}

	err = u.Update().SetLastLogin(now).Exec(ctx)
	if err != nil {
		return nil, err
	}

	if lastLoginAttempt != nil && lastLogin != nil && *lastLoginAttempt == *lastLogin {
		lastLoginAttempt = nil
	}

	log.Info("Successful login for", uname)
	return &LoginRes{
		User:             u.User,
		LastLogin:        lastLogin,
		LastLoginFailure: lastLoginAttempt,
	}, nil
}

func RegisterUser(ctx context.Context, network uumap.Network) (*LoginRes, error) {
	rx, err := regexp.Compile("^[a-z0-9_-]+$")
	if err != nil {
		return nil, err
	}
	var uname string
	var password schema.Password

	for {
		uname, err = io.ReadNotEmpty[string]("Enter your prefered username: ")
		if err != nil {
			return nil, err
		}

		if !rx.MatchString(uname) {
			fmt.Println("You can use lowercase letters, digits, '_' and '-'.")
			continue
		}

		if uname == "guest" || uname == "sys" {
			fmt.Println("Username is taken")
			continue
		}

		c, err := network.Client.User.Query().Where(user.UserEQ(schema.Uname(uname))).Count(ctx)
		if err != nil {
			return nil, err
		}
		if c > 0 {
			fmt.Println("Username is taken")
			continue
		}
		break
	}

	for {
		password, err = io.ReadPassword("Enter password: ")
		if err != nil {
			return nil, err
		}
		if len(password) < 6 {
			fmt.Println("Try a more complex one.")
			continue
		}
		break
	}

	hash, err := password.Hash()
	if err != nil {
		return nil, err
	}

	err = network.Client.User.Create().
		SetUser(schema.Uname(uname)).
		SetPassword(hash).
		SetStatus("").
		SetLastLogin(time.Now()).
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	log.Info("Registered", uname)

	return &LoginRes{User: schema.Uname(uname)}, nil
}
