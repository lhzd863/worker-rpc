package module

type MetaWorkerConf struct {
	Version    string `yaml:"version"`
	Name       string `yaml:"name"`
	Ip         string `yaml:"ip"`
	Port       string `yaml:"port"`
	HomeDir    string `yaml:"homedir"`
	ProcessNum string `yaml:"processnum"`
}

type MetaWorkerBean struct {
	Id         string `json:"id"`
	WorkerId   string `json:"workerid"`
	Ip         string `json:"ip"`
	Port       string `json:"port"`
	MaxCnt     string `json:"maxcnt"`
	RunningCnt string `json:"runningcnt"`
	StartTime  string `json:"starttime"`
	UpdateTime string `json:"updatetime"`
	Duration   string `json:"duration"`
}

type MetaWorkerStopJobBean struct {
        Id    string        `json:"id"`
}

