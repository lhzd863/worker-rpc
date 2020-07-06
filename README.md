# worker-rpc
worker 节点远程执行作业脚本

# 功能
```
1.根据传入job在服务器上执行job
2.脚本执行失败支持重试，根据retry决定重试次数。retry<=0 默认脚本执行一次
3.脚本设置超时，根据timeout设置决定时长。timeout<=0 为不设置timeout
4.job设置多个环境变量，命令中环境变量被替换，脚本中可以使用环境变量
5.job设置多个脚本按照顺序执行


```

# 作业配置格式
```
version: v1.0
name: yangtze
attr:
  #
  t-10000:
    sys: "XXDQC"
    job: "XXDQC_OSUB_A01_ANALYSIS_REPORT_INVIEW"
    cmd:
      - "bash /home/k8s/workspace-go/worker-rpc/cmd/t.sh XXSDL_HISA_S01_TEST.${cal_date}.000000"
      - "bash /home/k8s/workspace-go/worker-rpc/cmd/t1.sh XXSDL_HISA_S01_TMP.${cal_date1d}.000000"
    parameter:
      - "{\"k\":\"cal_date\",\"v\":\"20200703\"}"
      - "{\"k\":\"cal_date1d\",\"v\":\"20200702\"}"
    timeout: 3600
    retry: 2
    ip: "192.168.1.189"
    port: "12525"
    enable: "1"

```

# 启动worker
```
go run main.go --run="worker" --conf=worker-conf.yaml
```

# 查看作业状态
```
go run main.go --run="status" --conf=job-conf.yaml --ctl="XXDQC.XXDQC_OSUB_A01_ANALYSIS_REPORT_INVIEW.20200703.000000"
```

# 停止作业运行
```
go run main.go --run="stop" --conf=job-conf.yaml --ctl="XXDQC.XXDQC_OSUB_A01_ANALYSIS_REPORT_INVIEW.20200703.000000" \
--jobid="3f9e8625-8042-4a9a-b01f-832cc63c6924"
```

# 默认环境变量
```
CTX_DATE=20200703
CTX_TIME=000000
CTX_TIMESTAMP=20200703000000
CTX_CTL=XXDQC.XXDQC_OSUB_A01_ANALYSIS_REPORT_INVIEW.20200703.000000
CTX_SYS=XXDQC
CTX_JOB=XXDQC_OSUB_A01_ANALYSIS_REPORT_INVIEW
CTX_STR=2020-07-03 00:00:00
```

# LOG路径格式
```
LOG                                                                          #LOG根目录
└── 20200703                                                                 #批量日期
    └── 000000                                                               #批量时间
        └── XXDQC                                                            #系统名
            ├── xxdqc_osub_a01_analysis_report_inview_0_0_20200706141444.log #脚本LOG内容#作业名(小写)_脚本序列_执行次数_时间戳.log
            ├── xxdqc_osub_a01_analysis_report_inview_0_1_20200706141516.log       
            └── XXDQC_OSUB_A01_ANALYSIS_REPORT_INVIEW_20200703000000.log     #作业执行log#作业名(大写)_批量时间戳.log
```
