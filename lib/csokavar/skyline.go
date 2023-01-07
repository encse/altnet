package csokavar

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/encse/altnet/lib/cache"
	"github.com/encse/altnet/lib/io"
)

type GithubActivity struct {
	Username      string                `json:"username"`
	Year          string                `json:"year"`
	Min           float32               `json:"min"`
	Max           float32               `json:"max"`
	Median        float32               `json:"median"`
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
	return cache.Cached("github-skyline-for-"+githubUser, 1*time.Hour, func() (string, error) {
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
		msg += io.Center(fmt.Sprintf("Github SkyLine for %v\n", year), screenWidth)
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
					row += "█"
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
			msg += io.Center(row, screenWidth)
		}
		msg += io.Center(fmt.Sprintf("https://github.com/%v/", githubUser), screenWidth) + "\n"
		msg += "\n"
		msg += "\n"
		return msg, nil
	})
}
