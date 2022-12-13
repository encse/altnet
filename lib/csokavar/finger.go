package csokavar

import (
	"strings"

	log "github.com/sirupsen/logrus"
)

func Finger(user string, screenWidth int) string {
	var out strings.Builder

	banner := Banner(screenWidth)
	out.WriteString(banner)

	logo, err := Logo(screenWidth)
	if err != nil {
		log.Error(err)
	} else {
		out.WriteString(logo)
	}
	out.WriteString("\n")

	tweets, err := GetTweets(user, screenWidth)
	if err != nil {
		log.Error(err)
	} else {
		out.WriteString(tweets)
	}

	skyline, err := GetSkyline(user, screenWidth)
	if err != nil {
		log.Error(err)
	} else {
		out.WriteString(skyline)
	}

	contact, err := GpgKey(screenWidth)
	if err != nil {
		log.Error(err)
	} else {
		out.WriteString(contact)
	}

	return out.String()
}
