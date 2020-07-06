
#nohup go run cli-main.go --conf=job-conf.yaml --ctl="XXDQC.XXDQC_OSUB_A01_ANALYSIS_REPORT_INVIEW.20200703.000000"> run-cli.log 2>&1 &

go run cli-main.go --run="invoke" --conf=job-conf.yaml --ctl="XXDQC.XXDQC_OSUB_A01_ANALYSIS_REPORT_INVIEW.20200703.000000"
retcd=$?
echo "$retcd"


