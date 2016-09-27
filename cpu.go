package main

import "time"

//固定cpu曲线
func cpu() {
	//cpu 2.3 GHz
	var hz int = 2300 * 1000 * 1000
	//一个时钟周期大概执行4行代码,cpu%50,所以执行2行代码,分为10次执行,即
	num := hz * 4 / 2 / 100
	for {
		for i := 0; i < num; i++ {
			time.Sleep(10 * time.Millisecond)
		}
	}
}
