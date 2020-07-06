package module

import (
	"os/exec"
)

type MetaParaJobWorkerBean struct {
	Id        string        `json:"id"`
	WorkerId  string        `json:"workerid"`
	Sys       string        `josn:"sys"`
	Job       string        `json:"job"`
	Context   string        `json:"context"`
	Cmd       []interface{} `json:"cmd"`
	Parameter []interface{} `json:"parameter"`
	TimeOut   int64         `json:"timeout"`
	Retry     int8          `json:"retry"`
}

type MetaJobWorkerBean struct {
	Id         string        `json:"id"`
	WorkerId   string        `json:"workerid"`
	Ip         string        `json:"ip"`
	Port       string        `json:"port"`
	Sys        string        `josn:"sys"`
	Job        string        `json:"job"`
	Context    string        `json:"context"`
	Cmd        []interface{} `json:"cmd"`
	Parameter  []interface{} `json:"parameter"`
	TimeOut    int64         `json:"timeout"`
	Retry      int8          `json:"retry"`
	StartTime  string        `json:"starttime"`
	EndTime    string        `json:"endtime"`
	Status     string        `json:"status"`
	Command    *exec.Cmd     `json:"command"`
	CmdRunning string        `json:"cmdrunning"`
	RetCode    string        `json:"retcode"`
}

type KVBean struct {
	K string `json:"k"`
	V string `json:"v"`
}
