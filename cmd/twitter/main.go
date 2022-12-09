package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"unicode/utf8"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/encse/altnet/lib/config"
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
	for _, line := range strings.Split(linebreak(txt, width-4), "\n") {
		res += fmt.Sprintf("| %-*s |\n", width-4, line)
	}
	label = "--[" + label + "]--"
	res += "+" + strings.Repeat("-", width-len(label)-2) + label + "+\n"
	return res
}

func linebreak(text string, width int) string {
	lines := strings.Split(text, "\n")
	for i := 0; i < len(lines); i++ {
		line := []rune(lines[i])
		ichSpace := 0
		nonEscapedChars := 0
		for ich := 0; ich < len(line); ich++ {
			nonEscapedChars++
			if line[ich] == ' ' {
				ichSpace = ich
			}
			if nonEscapedChars > width {
				if ichSpace > 0 {
					lines = append(lines, "")
					copy(lines[i+1:], lines[i:])
					lines[i] = strings.TrimRight(string(line[:ichSpace]), " ")
					lines[i+1] = strings.TrimRight(string(line[ichSpace+1:]), " ")
				}
				break
			}
		}
	}
	return strings.Join(lines, "\n")
}

func center(st string, width int) string {
	lines := strings.Split(st, "\n")
	maxWidth := 0
	for _, line := range lines {
		if utf8.RuneCountInString(line) >= maxWidth {
			maxWidth = utf8.RuneCountInString(line)
		}
	}
	for i, line := range lines {
		pad := (width - maxWidth) / 2
		if line == "" || pad <= 0 {
			continue
		}
		lines[i] = strings.Repeat(" ", pad) + line
	}
	return strings.Join(lines, "\n")
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
