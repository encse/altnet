package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/encse/altnet/lib/config"
	"github.com/encse/altnet/lib/text"
)

type TimelineEntry struct {
	Id                string `json:"id_str"`
	CreatedAt         string `json:"created_at"`
	Text              string `json:"full_text"`
	InReplyToStatusId string `json:"in_reply_to_status_id_str"`
}

type GithubContributions struct {
	Week int        `json:"week"`
	Days []DayCount `json:"days"`
}

type DayCount struct {
	Count int `json:"count"`
}

func GetTweets(twitterAccessToken string, twitterUser string, screenWidth int) (string, error) {

	client := &http.Client{}

	request, err := http.NewRequest("GET", fmt.Sprintf("https://api.twitter.com/1.1/statuses/user_timeline.json?screen_name=%v&tweet_mode=extended", "encse"), nil)
	if err != nil {
		return "", err
	}

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %v", twitterAccessToken))
	rsp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer rsp.Body.Close()
	var timeline []TimelineEntry
	err = json.NewDecoder(rsp.Body).Decode(&timeline)
	if err != nil {
		return "", err
	}

	ids := mapset.NewSet[string]()
	for _, timelineEntry := range timeline {
		ids.Add(timelineEntry.Id)
	}

	res := fmt.Sprintf("Latest tweets https://twitter.com/%v\n", twitterUser)
	res += "\n"
	for _, entry := range timeline {
		if entry.InReplyToStatusId == "" {
			texts := []string{}
			for _, threadEntry := range createThread(timeline, entry) {
				texts = append(texts, threadEntry.Text+"\n")
			}
			res += box(
				strings.Join(texts, "\n"),
				entry.CreatedAt,
				screenWidth,
			)
			res += "\n"
		}
	}

	return res, nil
}

func createThread(timeline []TimelineEntry, currentEntry TimelineEntry) []TimelineEntry {
	res := []TimelineEntry{currentEntry}
	for _, entry := range timeline {
		if entry.InReplyToStatusId == currentEntry.Id {
			for _, threadEntry := range createThread(timeline, entry) {
				res = append(res, threadEntry)
			}
		}
	}
	return res
}

func box(txt, label string, width int) string {
	res := "+" + strings.Repeat("-", width-2) + "+\n"
	for _, line := range strings.Split(text.Linebreak(txt, width-4), "\n") {
		res += fmt.Sprintf("| %-*s |\n", width-4, line)
	}
	label = "--[" + label + "]--"
	res += "+" + strings.Repeat("-", width-len(label)-2) + label + "+\n"
	return res
}

func main() {
	fmt.Println(os.Getenv("TWITTER_ACCESS_TOKEN"))
	config, err := config.Read("../../config.yml")
	if err != nil {
		fmt.Println(err)
		return
	}

	st, err := GetTweets(config.Twitter.AccessToken, "encse", 80)
	if err == nil {
		fmt.Println(st)
	} else {
		fmt.Println(err)
	}
}
