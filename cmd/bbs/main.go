package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/encse/altnet/lib/io"
	"github.com/hako/durafmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"golang.org/x/term"
)

func Banner(screenWidth int) string {
	var out strings.Builder

	arch := "-"
	platform := "-"
	cpus := "-"
	load1 := "-"
	load5 := "-"
	load15 := "-"
	uptime := "-"

	infoStat, err := host.Info()
	if err == nil {
		arch = infoStat.KernelArch
		platform = infoStat.OS
		uptime = fmt.Sprintf("%v", durafmt.Parse(time.Duration(infoStat.Uptime)*time.Second))
	}

	cpuCount, err := cpu.Counts(false)
	if err == nil {
		cpus = strconv.Itoa(cpuCount)
	}

	loadAvg, err := load.Avg()
	if err == nil {
		load1 = fmt.Sprintf("%.2f", loadAvg.Load1)
		load5 = fmt.Sprintf("%.2f", loadAvg.Load5)
		load15 = fmt.Sprintf("%.2f", loadAvg.Load15)
	}

	fmt.Println(io.Center("Connected to CSOKAVAR, Encse's home on the web. Happy surfing.", screenWidth))
	fmt.Println()
	fmt.Println(
		io.Center(
			fmt.Sprintf("Server: %v %v with %v cpu(s), load average: %v, %v, %v", arch, platform, cpus, load1, load5, load15),
			screenWidth,
		))

	fmt.Println(io.Center(fmt.Sprintf("uptime: %v", uptime), screenWidth))
	fmt.Println()
	fmt.Println(io.Center("SysOp: encse", screenWidth))
	fmt.Println()

	return out.String()
}

func main() {
	screenWidth, _, err := term.GetSize(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(Banner(screenWidth))

	fmt.Println("Enter your username or GUEST")
	fmt.Print("Username: ")

	username := ""
	for username == "" {
		username, err = io.Readline()
		if err != nil {
			log.Fatal(err)
		}
		username = strings.TrimSpace(strings.ToLower(username))
	}

	if username != "guest" {
		for i := 0; i < 3; i++ {
			fmt.Print("Password: ")
			_, err = io.ReadPassword()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("")
		}
		return
	}

	logo, err := ioutil.ReadFile("data/logo.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(logo))
	fmt.Println("Welcome", username)

loop:
	for {
		fmt.Println("BBS Menu")
		fmt.Println("------------")
		fmt.Println(": Latest [T]weets")
		fmt.Println(": [G]itHub skyline")
		fmt.Println(": [C]ontact sysop")
		fmt.Println(": play [I]dőrégész")
		fmt.Println(": e[X]it")

		option, err := io.ReadOption("Select an item", "tgcix")
		if err != nil {
			log.Fatal(err)
		}
		switch strings.ToLower(option) {
		case "t":
			cmd := exec.Command("./twitter")
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
			if err != nil {
				log.Fatal(err)
			}
		case "g":
			cmd := exec.Command("./githubskyline")
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
			if err != nil {
				log.Fatal(err)
			}
		case "c":
			cmd := exec.Command("/bin/bash")
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
			if err != nil {
				log.Fatal(err)
			}
		case "i":
			fmt.Println("idoregesz")
			fmt.Println("idoregesz")
			fmt.Println("idoregesz")
			fmt.Println("idoregesz")
			fmt.Println("idoregesz")
		case "x":
			break loop
		}
	}

	fmt.Println("Have a nice day!")

	footer, err := ioutil.ReadFile("data/footer.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(footer))
}
