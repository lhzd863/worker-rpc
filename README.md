# worker-rpc
worker 节点远程执行作业脚本

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
└── LOG                                                                          #LOG根目录
    └── 20200703                                                                 #批量日期
        └── 000000                                                               #批量时间
            └── XXDQC                                                            #系统名
                ├── xxdqc_osub_a01_analysis_report_inview_0_0_20200706141444.log #脚本LOG内容#作业名(小写)_脚本序列_执行次数_时间戳.log
                ├── xxdqc_osub_a01_analysis_report_inview_0_1_20200706141516.log       
                └── XXDQC_OSUB_A01_ANALYSIS_REPORT_INVIEW_20200703000000.log     #作业名(大写)_批量时间戳.log
```
