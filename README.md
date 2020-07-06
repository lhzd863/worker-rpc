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

# 查看作业状态
```
go run main.go --run="status" --conf=job-conf.yaml --ctl="XXDQC.XXDQC_OSUB_A01_ANALYSIS_REPORT_INVIEW.20200703.000000"
```

# 停止作业运行
```
go run main.go --run="stop" --conf=job-conf.yaml --ctl="XXDQC.XXDQC_OSUB_A01_ANALYSIS_REPORT_INVIEW.20200703.000000" --jobid="3f9e8625-8042-4a9a-b01f-832cc63c6924"
```

