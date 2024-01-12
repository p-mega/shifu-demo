package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/robfig/cron/v3"
)

func main() {
	// 新建一个定时任务对象,根据cron表达式进行时间调度,cron精确到秒
	cronTab := cron.New(cron.WithSeconds())
	// 定义定时器调用的任务函数
	task := GetMeasurement
	// 定时任务,cron表达式,每五秒一次
	spec := "*/5 * * * * ?"
	// 添加定时任务
	cronTab.AddFunc(spec, task)
	// 启动定时器
	cronTab.Start()
	// 阻塞主线程停止
	select {}
}

func GetMeasurement() {
	// 获取数据
	res, err := http.Get("http://deviceshifu-plate-reader.deviceshifu.svc.cluster.local/get_measurement")
	if err != nil {
		log.Printf("err: %s", err)
		return
	}
	// 数据转化
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("err: %s", err)
		return
	}
	// 数据拆分
	data := strings.Fields(string(body))
	// 定义数据的平均数和长度
	avg, l := 0.0, len(data)
	// 计算平均数
	for _, num := range data {
		n, err := strconv.ParseFloat(num, 64)
		if err != nil {
			log.Printf("err: %s", err)
			return
		}
		avg += n / float64(l)
	}
	log.Printf("avg: %f", avg)
}
