package altnet

import (
	"fmt"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/encse/altnet/lib/log"
	"github.com/encse/altnet/lib/uumap"
	"github.com/shirou/gopsutil/v3/process"
)

type Exe string
type Terminal string
type Pid int32
type SessionId int32

type ProcInfo struct {
	Pid     Pid
	Exe     Exe
	User    User
	Started time.Time
}

func KillSession(sessionId SessionId, signal process.Signal) error {
	log.Info("Killing session", sessionId)

	any := true
	for any {

		any = false
		processes, err := process.Processes()
		if err != nil {
			return fmt.Errorf("killing session: error collecting process information, %w", err)
		}

		for _, process := range processes {
			environ, err := process.Environ()
			if err != nil {
				log.Error(err)
				continue
			}

			for _, variable := range environ {
				parts := strings.Split(string(variable), "=")
				if len(parts) == 2 {
					if parts[0] == string(sessionKey) {
						if parts[1] == strconv.Itoa(int(sessionId)) {

							any = true
							err := syscall.Kill(int(process.Pid), signal)
							if err != nil {
								log.Error(err)
							}
						}
					}
				}
			}
		}
	}
	log.Info("Killed session %v", sessionId)
	return nil
}

func GetProcesses(host uumap.Host) ([]ProcInfo, error) {
	processes, err := process.Processes()
	if err != nil {
		return nil, err
	}

	var res []ProcInfo
	for _, process := range processes {

		createTime, err := process.CreateTime()
		if err != nil {
			log.Error(err)
			continue
		}

		environ, err := process.Environ()
		if err != nil {
			log.Error(err)
			continue
		}

		env := map[string]string{}
		for _, variable := range environ {
			parts := strings.Split(string(variable), "=")
			if len(parts) == 2 {
				env[parts[0]] = parts[1]
			}
		}

		procHost := env[string(hostKey)]
		procUser := env[string(userKey)]
		procExe := env[string(exeKey)]

		if uumap.Host(procHost) == host && procUser != "" && procExe != "" {
			res = append(res, ProcInfo{
				User:    User(procUser),
				Exe:     Exe(procExe),
				Pid:     Pid(process.Pid),
				Started: time.Unix(createTime/1000, createTime%1000),
			})
		}
	}
	return res, nil
}
