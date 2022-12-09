package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
	"unicode/utf8"
)

type GithubActivity struct {
	Username      string                `json:"username"`
	Year          string                `json:"year"`
	Min           int                   `json:"min"`
	Max           int                   `json:"max"`
	Median        int                   `json:"median"`
	P80           float32               `json:"p80"`
	P90           float32               `json:"p90"`
	P99           float32               `json:"p99"`
	Contributions []GithubContributions `json:"contributions"`
}

type GithubContributions struct {
	Week int        `json:"week"`
	Days []DayCount `json:"days"`
}

type DayCount struct {
	Count int `json:"count"`
}

func GetSkyline(githubUser string, screenWidth int) (string, error) {
	year := time.Now().Year()
	rsp, err := http.Get(fmt.Sprintf("https://skyline.github.com/%v/%v.json", githubUser, year))
	if err != nil {
		return "", err
	}
	defer rsp.Body.Close()

	var githubActivity GithubActivity
	err = json.NewDecoder(rsp.Body).Decode(&githubActivity)
	if err != nil {
		return "", err
	}

	d := githubActivity.Max / 8

	msg := "\n"
	msg += center(fmt.Sprintf("Github SkyLine for %v\n", year), screenWidth)
	msg += "\n"
	msg += "\n"

	for j := 8; j >= 0; j-- {
		row := ""
		for _, contibution := range githubActivity.Contributions {
			maxPerWeek := 0
			for _, day := range contibution.Days {
				if day.Count > maxPerWeek {
					maxPerWeek = day.Count
				}
			}

			if maxPerWeek >= d*j {
				row += "#" // █
			} else {
				r := rand.Float32()
				if r <= 0.025 {
					row += "*"
				} else if r <= 0.050 {
					row += "*"
				} else if r <= 0.055 {
					row += "("
				} else {
					row += " "
				}
			}
		}
		row += "\n"
		msg += center(row, screenWidth)
	}
	msg += center(fmt.Sprintf("https://github.com/%v/", githubUser), screenWidth) + "\n"
	msg += "\n"
	msg += "\n"
	return msg, nil
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
	st, err := GetSkyline("encse", 120)
	if err == nil {
		fmt.Println(st)
	} else {
		fmt.Println(err)
	}
}
