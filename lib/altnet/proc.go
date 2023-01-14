package altnet

import (
	"strings"
	"time"

	"github.com/encse/altnet/lib/log"
	"github.com/shirou/gopsutil/v3/process"
)

type Exe string
type Pid int32

type ProcInfo struct {
	Pid     Pid
	Exe     Exe
	User    User
	Started time.Time
}

func GetProcesses(host Host) ([]ProcInfo, error) {
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

		if Host(procHost) == host && procUser != "" && procExe != "" {
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
