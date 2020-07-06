package cli

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/lhzd863/worker-rpc/gproto"
	"github.com/lhzd863/worker-rpc/module"
)

type Cli struct {
	Ctl string
	cfg string
}

func NewCli(paraMap map[string]interface{}) *Cli {
	ctl := paraMap["ctl"].(string)
	cfg := paraMap["cfg"].(string)
	m := &Cli{Ctl: ctl, cfg: cfg}
	return m
}

func (c *Cli) InvokeJob() (int32, error) {
	m, err := c.ParseJobYaml()
	if err != nil {
		log.Printf(fmt.Sprintf("parse yaml err: %v", err))
		return 1, err
	}
	// 建立连接到gRPC服务
	conn, err := grpc.Dial(m.Ip+":"+m.Port, grpc.WithInsecure())
	if err != nil {
		log.Printf(fmt.Sprintf("did not connect: %v:%v, %v", m.Ip, m.Port, err))
		return 1, err
	}
	// 函数结束时关闭连接
	defer conn.Close()

	// 创建Waiter服务的客户端
	t := gproto.NewWorkerServerClient(conn)

	jsonstr, err := json.Marshal(m)
	if err != nil {
		log.Printf(fmt.Sprintf("json marshal %v", err))
		return 1, err
	}
	// 调用gRPC接口
	tr, err := t.JobStart(context.Background(), &gproto.Req{JsonStr: string(jsonstr)})
	if err != nil {
		log.Printf(fmt.Sprintf("could not greet: %v", err))
		return 1, err
	}
	//change job status
	if tr.Status_Code != 0 {
		log.Printf(fmt.Sprint(tr.Status_Txt))
		return tr.Status_Code, fmt.Errorf("%v", tr.Status_Txt)
	}
	return 0, nil
}

func (c *Cli) StatusJob() (int32, error) {
	m, err := c.ParseJobYaml()
	if err != nil {
		log.Printf(fmt.Sprintf("parse yaml err: %v", err))
		return 1, err
	}
	// 建立连接到gRPC服务
	conn, err := grpc.Dial(m.Ip+":"+m.Port, grpc.WithInsecure())
	if err != nil {
		log.Printf(fmt.Sprintf("did not connect: %v:%v, %v", m.Ip, m.Port, err))
		return 1, err
	}
	// 函数结束时关闭连接
	defer conn.Close()

	// 创建Waiter服务的客户端
	t := gproto.NewWorkerServerClient(conn)

	jsonstr, err := json.Marshal(m)
	if err != nil {
		log.Printf(fmt.Sprintf("json marshal %v", err))
		return 1, err
	}
	// 调用gRPC接口
	tr, err := t.JobStatus(context.Background(), &gproto.Req{JsonStr: string(jsonstr)})
	if err != nil {
		log.Printf(fmt.Sprintf("could not greet: %v", err))
		return 1, err
	}
	//change job status
	if tr.Status_Code != 0 {
		log.Printf(fmt.Sprint(tr.Status_Txt))
		return tr.Status_Code, fmt.Errorf("%v", tr.Status_Txt)
	}
	log.Printf(fmt.Sprint(tr.Data))
	return 0, nil
}

func (c *Cli) StopJob(id string) (int32, error) {
	m, err := c.ParseJobYaml()
	if err != nil {
		log.Printf(fmt.Sprintf("parse yaml err: %v", err))
		return 1, err
	}
	// 建立连接到gRPC服务
	conn, err := grpc.Dial(m.Ip+":"+m.Port, grpc.WithInsecure())
	if err != nil {
		log.Printf(fmt.Sprintf("did not connect: %v:%v, %v", m.Ip, m.Port, err))
		return 1, err
	}
	// 函数结束时关闭连接
	defer conn.Close()

	// 创建Waiter服务的客户端
	t := gproto.NewWorkerServerClient(conn)
	p := new(module.MetaWorkerStopJobBean)
	p.Id = id
	jsonstr, err := json.Marshal(p)
	if err != nil {
		log.Printf(fmt.Sprintf("json marshal %v", err))
		return 1, err
	}
	// 调用gRPC接口
	tr, err := t.JobStop(context.Background(), &gproto.Req{JsonStr: string(jsonstr)})
	if err != nil {
		log.Printf(fmt.Sprintf("could not greet: %v", err))
		return 1, err
	}
	//change job status
	if tr.Status_Code != 0 {
		log.Printf(fmt.Sprint(tr.Status_Txt))
		return tr.Status_Code, fmt.Errorf("%v", tr.Status_Txt)
	}
	log.Printf(fmt.Sprint(tr.Data))
	return 0, nil
}

func (c *Cli) ParseJobYaml() (*module.MetaJobWorkerBean, error) {
	sys, job, ctx := ctlsparse(c.Ctl)
	cf := new(module.MetaJobConf)
	m := new(module.MetaJobWorkerBean)
	yamlFile, err := ioutil.ReadFile(c.cfg)
	if err != nil {
		log.Printf("error: %s", err)
		return m, err
	}
	err = yaml.UnmarshalStrict(yamlFile, cf)
	if err != nil {
		log.Printf("error: %s", err)
		return m, err
	}
	for _, a := range cf.Attr {
		if sys != a.Sys || job != a.Job {
			continue
		}
		if a.Enable != "1" {
			log.Printf("%v.%v enable %v is not equals 1.", a.Sys, a.Job, a.Enable)
			return m, fmt.Errorf("%v.%v enable %v is not equals 1.", a.Sys, a.Job, a.Enable)
		}
		m.Sys = a.Sys
		m.Job = a.Job
		m.Context = ctx
		m.Cmd = a.Cmd
		m.Parameter = a.Parameter
		m.TimeOut = a.TimeOut
		m.Retry = a.Retry
		m.Ip = a.Ip
		m.Port = a.Port
		break
	}
	return m, nil
}
func ctlsparse(ctl string) (string, string, string) {
	arr := strings.Split(ctl, ".")
	date_fomat := fmt.Sprintf("%v-%v-%v %v:%v:%v", arr[2][0:4], arr[2][4:6], arr[2][6:8], arr[3][0:2], arr[3][2:4], arr[3][4:6])
	return arr[0], arr[1], date_fomat
}
