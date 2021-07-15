package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
)

var NodeGPU map[string]NodeGPUstate = make(map[string]NodeGPUstate)

func main() {
    Podmp = make(map[string]TaskLog)

	File, err := os.Open("./date.log")
	defer File.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	rd := bufio.NewReader(File)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		line = line[:len(line)-1]
		path := "./"+line+"/"
		Readtaskinfo(path)
		Readcpuinfo(path)
		Readgpuinfomem(path)
		Readgpuinfoutil(path)
		Readgpuinfopcie(path)
		Readmeminfo(path)
		//Readgpumemcpyutil(path)
		Readnodegpuinfo(path)

		// Readnodegpuinfo(path, &tmp)
		// Readnodegpuuse(path, &tmp)
		// Readnodecpuuti(path, &tmp)
		// Readnodeio(path, tmp)
	}
	//CalTime()
	//CalNodeBusy()
	//CalGPUWaste("./gpuwaste.log")
	AverageVal("./pcie.log")
	//TaskMemUtui("./tmp.log")
	//DiffResourceTGPUUti("./tmp.log")
	//SCGPUUtui("./tmp-1.log")
	//MCGPUUti("./mgpuload.log")
	//Getgpucpu("./tmp.log")
	//PrintGPUCPUUtiRange("./tmp-1.log")
	//CalTaskCnt("./tmp.log", 1, 100)
	
}
