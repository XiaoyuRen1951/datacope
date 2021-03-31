package main

import (
	//"fmt"
)

func main() {
    Podmp = make(map[string]TaskLog)
	Readtaskinfo()
	Readcpuinfo()
	Readgpuinfomem()
	Readgpuinfoutil()
	Readmeminfo()
	//Readgpumemcpyutil()
	
	//Readnodegpuinfo(&tmp)
	//Readnodegpuuse(&tmp)
	//Readnodecpuuti(&tmp)
	//Readnodeio(tmp)
	
	AverageVal("./tmp.log")
	//DiffResourceTGPUUti("./tmp.log")
	//SCGPUUtui("./tmp.log")
	//MCGPUUti("./tmp.log")
	//Getgpucpu("./tmp.log")
	//PrintGPUCPUUtiRange("./tmp.log")
}
