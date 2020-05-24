package main

import "runtime"

//初始化程序使用服务器的线程数
func initEnv(){
	runtime.GOMAXPROCS(runtime.NumCPU())
}



func main() {
	//初始化程序
	initEnv()

	//加载全局配置

	//api http服务启动


}
