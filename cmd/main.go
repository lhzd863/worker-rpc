package main

import (
	"flag"
	"log"
	"os"

	"github.com/lhzd863/worker-rpc/cli"
	"github.com/lhzd863/worker-rpc/module"
        "github.com/lhzd863/worker-rpc/worker"
)

var (
	cfg   = flag.String("conf", "conf.yaml", "basic config")
	ctl   = flag.String("ctl", "XXDQC.XXDQC_OSUB_A01_ANALYSIS_REPORT_INVIEW.YYYYMMDD.HHmmSS", "")
	conf  *module.MetaCliConf
	run   = flag.String("run", "invoke", "invoke job")
	jobid = flag.String("jobid", "jobid", "job id")
)

func main() {
	flag.Parse()

	mpara := make(map[string]interface{})
	mpara["cfg"] = *cfg
	if *run == "worker" {
		ws := worker.NewWorkerServer(mpara)
		ws.Main()
		return
	}

        mpara["ctl"] = *ctl
	s := cli.NewCli(mpara)
	if *run == "invoke" {
		retcd, err := s.InvokeJob()
		if err != nil {
			log.Printf("error: %s", err)
			os.Exit(int(retcd))
		}
		return
	}
	if *run == "status" {
		log.Printf(*run)
		retcd, err := s.StatusJob()
		if err != nil {
			log.Printf("error: %s", err)
			os.Exit(int(retcd))
		}
		return
	}
	if *run == "stop" {
		retcd, err := s.StopJob(*jobid)
		if err != nil {
			log.Printf("error: %s", err)
			os.Exit(int(retcd))
		}
		return
	}
}
