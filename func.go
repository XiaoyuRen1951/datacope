package main

import (
	"fmt"
	"os"
	//"math"
	"reflect"
	// "sort"
	"strconv"
	"strings"
	"encoding/json"
	"bufio"
	"io"
    "encoding/csv"
)

var Podmp map[string]TaskLog

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

func Cal_Average_int(src []int64) float64 {
	var sum int64 = 0
	if len(src) == 0 {
		return 0
	}
	for _,v := range src {
		sum = sum + v
	}

	return Decimal(float64(sum)/float64(len(src)))
}

func Cal_Average_float(src []float64) float64 {
	var sum float64 = 0
	if len(src) == 0 {
		return 0
	}
	for _,v := range src {
		sum = sum + v
	}

	return Decimal(float64(sum)/float64(len(src)))
}

func MIN(x, y interface{}) interface{} {
	if reflect.TypeOf(x).Name() == "int64" {
		if reflect.ValueOf(x).Int() < reflect.ValueOf(y).Int() {
			return x
		}
		return y
	} else if reflect.TypeOf(x).Name() == "float64" {
		if reflect.ValueOf(x).Float() < reflect.ValueOf(y).Float() {
			return x
		}
		return y
	}

	return x
}

func MAX(x, y interface{}) interface{} {
	if reflect.TypeOf(x).Name() == "int64" {
		if reflect.ValueOf(x).Int() > reflect.ValueOf(y).Int() {
			return x
		}
		return y
	} else if reflect.TypeOf(x).Name() == "float64" {
		if reflect.ValueOf(x).Float() > reflect.ValueOf(y).Float() {
			return x
		}
		return y
	}
	return x
}

func Readtaskinfo() error {
	File, err := os.Open("./data/taskinfo.csv")
	defer File.Close()

	if err != nil {
		fmt.Println(err)
		return err
	}

	rd := csv.NewReader(File)
    rd.Read()
	for {
		line, err := rd.Read()
		if err != nil || io.EOF == err {
			break
		}
		
		res := *new(TaskLog)
		// taskinfo.WriteString(v.Pod+","+v.Container+","+v.Namespace+","+v.User+","+v.ResourceT+","+v.Node.Name+",")
		// taskinfo.WriteString(fmt.Sprintf("%d,%d,%d,%d,%d\n",v.Starttime,v.Endtime,v.CPU.Limit, v.Memory.Limit, v.GPU.NumGPU))
		res.Pod = line[0]
		res.Container = line[1]
		res.Namespace = line[2]
		res.User = line[3]
		res.ResourceT = line[4]
		res.Node = *new(NodeInfo)
		res.Node.Name = line[5]
		res.Starttime, _ = strconv.ParseInt(line[6], 10, 64)
		res.Endtime, _ = strconv.ParseInt(line[7], 10, 64)
		res.CPU = *new(CPUInfo)
		res.CPU.Limit, _ = strconv.ParseInt(line[8], 10, 64)
		res.CPU.Node = *new(NodeInfo)
		res.CPU.Node.Name = res.Node.Name
		res.CPU.Pod = res.Pod
		res.CPU.History = make([]float64, 0)
		res.Memory = *new(MemoryInfo)
		res.Memory.Limit, _ = strconv.ParseInt(line[9], 10, 64)
		res.Memory.Node = *new(NodeInfo)
		res.Memory.Node.Name = res.Node.Name
		res.Memory.Pod = res.Pod
		res.Memory.History = make([]int64, 0)
		res.GPU = *new(GPUInfo)
		res.GPU.NumGPU, _ = strconv.ParseInt(line[10], 10, 64)
		res.GPU.Node = *new(NodeInfo)
		res.GPU.Node.Name = res.Node.Name
		res.GPU.Pod = res.Pod
		res.GPU.GPUUtil = make([]GPUHistory, 0)
		res.GPU.GPUMem = make([]GPUMemHistory, 0)

		
		Podmp[res.Pod] = res
	}
	return nil
}

func Readcpuinfo() error {
	File, err := os.Open("./data/cpuinfo.json")
	defer File.Close()

	if err != nil {
		fmt.Println(err)
		return err
	}

	rd := bufio.NewReader(File)

	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}

		if line[0] != '{' {
			continue
		}
		
		res := new(CPUInfo)
		prdec := json.NewDecoder(strings.NewReader(line))
		err = prdec.Decode(&res)
		if err != nil {
			fmt.Println(err)
			return err
		}
		
		if val, ok := Podmp[res.Pod]; ok {
			//Deal With CPU Utilization
			val.CPU.History = res.History
			Podmp[res.Pod] = val
			
		}

		
	}
	return nil
}

func Readgpuinfomem() error {
	File, err := os.Open("./data/gpuinfomem.json")
	defer File.Close()

	if err != nil {
		fmt.Println(err)
		return err
	}

	rd := bufio.NewReader(File)

	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}

		if line[0] != '{' {
			continue
		}
		
		res := new(GPUMemHistory)
		prdec := json.NewDecoder(strings.NewReader(line))
		err = prdec.Decode(&res)
		if err != nil {
			fmt.Println(err)
			return err
		}
		
		if val, ok := Podmp[res.Pod]; ok {
			//Deal With GPU Memory
			val.GPU.GPUMem = append(val.GPU.GPUMem, *res)
			Podmp[res.Pod] = val	
		}
	}
	return nil
}

func Readgpuinfoutil() error {
	File, err := os.Open("./data/gpuinfoutil.json")
	defer File.Close()

	if err != nil {
		fmt.Println(err)
		return err
	}

	rd := bufio.NewReader(File)

	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}

		if line[0] != '{' {
			continue
		}
		
		res := new(GPUHistory)
		prdec := json.NewDecoder(strings.NewReader(line))
		err = prdec.Decode(&res)
		if err != nil {
			fmt.Println(err)
			return err
		}
		
		if val, ok := Podmp[res.Pod]; ok {
			//Deal With GPU Utilization
			val.GPU.GPUUtil = append(val.GPU.GPUUtil, *res)
			Podmp[res.Pod] = val
			
		}
	}
	return nil
}

func Readgpumemcpyutil() error {
	File, err := os.Open("./data/gpumemcpyutil.json")
	defer File.Close()

	if err != nil {
		fmt.Println(err)
		return err
	}

	rd := bufio.NewReader(File)

	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}

		if line[0] != '{' {
			continue
		}
		
		res := new(GPUHistory)
		prdec := json.NewDecoder(strings.NewReader(line))
		err = prdec.Decode(&res)
		if err != nil {
			fmt.Println(err)
			return err
		}
		
		if val, ok := Podmp[res.Pod]; ok {
			//Deal With GPU Utilization
			val.GPU.GPUMemCopy = append(val.GPU.GPUMemCopy, *res)
			Podmp[res.Pod] = val
			
		}
	}
	return nil
}

func Readmeminfo() error {
	File, err := os.Open("./data/meminfo.json")
	defer File.Close()

	if err != nil {
		fmt.Println(err)
		return err
	}

	rd := bufio.NewReader(File)

	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}

		if line[0] != '{' {
			continue
		}
		
		res := new(MemoryInfo)
		prdec := json.NewDecoder(strings.NewReader(line))
		err = prdec.Decode(&res)
		if err != nil {
			fmt.Println(err)
			return err
		}
		
		if val, ok := Podmp[res.Pod]; ok {
			//Deal With Memory
			val.Memory.History = res.History
			Podmp[res.Pod] = val
			//fmt.Println(val.Memory)
		}
	}
	return nil
}

func Readnodegpuinfo(val *[]NodeGPUstate) error {
	File, err := os.Open("./data/nodegpuinfo.json")
	defer File.Close()

	if err != nil {
		fmt.Println(err)
		return err
	}

	rd := bufio.NewReader(File)

	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}

		if line[0] != '{' {
			continue
		}
		
		res := new(NodeGPUstate)
		prdec := json.NewDecoder(strings.NewReader(line))
		err = prdec.Decode(&res)
		if err != nil {
			fmt.Println(err)
			return err
		}
		
		*val = append(*val,*res)
	}
	return nil
}

func Readnodegpuuse(val *[]Nodegpuuse) error {
	File, err := os.Open("./data/nodegpuuse.csv")
	defer File.Close()

	if err != nil {
		fmt.Println(err)
		return err
	}

	rd := csv.NewReader(File)
    rd.Read()
	for {
		line, err := rd.Read()
		if err != nil || io.EOF == err {
			break
		}
		var tmp Nodegpuuse
		tmp.Pod = line[0]
		tmp.Uuid = line[1]
		tmp.Ratio, _ = strconv.ParseFloat(line[2], 64) 
		
		*val = append(*val, tmp)
	}
	return nil
}

func Readnodecpuuti(val *[]CPUCore) error {
	File, err := os.Open("./data/nodecpuuti.json")
	defer File.Close()

	if err != nil {
		fmt.Println(err)
		return err
	}

	rd := bufio.NewReader(File)

	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}

		if line[0] != '{' {
			continue
		}
		
		res := new(CPUCore)
		prdec := json.NewDecoder(strings.NewReader(line))
		err = prdec.Decode(&res)
		if err != nil {
			fmt.Println(err)
			return err
		}
		//Deal With CPU Utilization
		*val = append(*val, *res)
	}
	return nil
}

func Readnodeio(val map[string]NodeIO) error {
	File, err := os.Open("./data/nodereiverate.json")

	if err != nil {
		fmt.Println(err)
		return err
	}

	rd := bufio.NewReader(File)

	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}

		if line[0] != '{' {
			continue
		}
		
		res := new(NodeIO)
		prdec := json.NewDecoder(strings.NewReader(line))
		err = prdec.Decode(&res)
		if err != nil {
			fmt.Println(err)
			return err
		}
		
		val[res.Node] = *res
	}
	File.Close()

	File, err = os.Open("./data/nodetransrate.json")

	if err != nil {
		fmt.Println(err)
		return err
	}

	rd = bufio.NewReader(File)

	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}

		if line[0] != '{' {
			continue
		}
		
		res := new(NodeIO)
		prdec := json.NewDecoder(strings.NewReader(line))
		err = prdec.Decode(&res)
		if err != nil {
			fmt.Println(err)
			return err
		}
		
		if v, ok := val[res.Node]; ok {
			v.ORate = res.ORate
			val[res.Node] = v
		} else {
			val[res.Node] = *res
		}
	}
	File.Close()

	File, err = os.Open("./data/nodeibreiverate.json")

	if err != nil {
		fmt.Println(err)
		return err
	}

	rd = bufio.NewReader(File)

	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}

		if line[0] != '{' {
			continue
		}
		
		res := new(NodeIO)
		prdec := json.NewDecoder(strings.NewReader(line))
		err = prdec.Decode(&res)
		if err != nil {
			fmt.Println(err)
			return err
		}
		
		if v, ok := val[res.Node]; ok {
			v.IbIRate = res.IbIRate
			val[res.Node] = v
		} else {
			val[res.Node] = *res
		}
	}
	File.Close()

	File, err = os.Open("./data/nodeibtransrate.json")

	if err != nil {
		fmt.Println(err)
		return err
	}

	rd = bufio.NewReader(File)

	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}

		if line[0] != '{' {
			continue
		}
		
		res := new(NodeIO)
		prdec := json.NewDecoder(strings.NewReader(line))
		err = prdec.Decode(&res)
		if err != nil {
			fmt.Println(err)
			return err
		}
		
		if v, ok := val[res.Node]; ok {
			v.IbORate = res.IbORate
			val[res.Node] = v
		} else {
			val[res.Node] = *res
		}
	}
	File.Close()

	return nil
}

func AverageVal(filepath string) {
	tmp, err := os.Create(filepath)
	
	if err != nil {
		fmt.Println("renxy-task File creating error", err)
		return
	}

	for _,v := range Podmp {
		if v.ResourceT == "debug" {
			continue
		}
		if v.GPU.NumGPU != 1 {
			continue
		}

		if strings.Contains(v.Pod,"renxy") != true {
			continue
		}
		//filter Task Duration
		// if v.Endtime - v.Starttime <=3600 {
		// 	continue
		// }

		var res float64 = 0
		for _,vv := range v.GPU.GPUUtil {
			res += Cal_Average_int(vv.History)
		}
		if leng:=len(v.GPU.GPUUtil); leng!=0 {
			res = Decimal(res / float64(leng))
		} else {
			res = 0
		}

		var gpumemr float64 = 0
		var gpumemmxr float64 = 0

		for _, vv := range v.GPU.GPUMem {
			gpumemr += Cal_Average_int(vv.History) * 100 / float64(vv.Total)
			gpumemmxr = MAX(gpumemmxr, float64(vv.MaxR*100) / float64(vv.Total)).(float64)
		}
		
		gpumemmxr = Decimal(gpumemmxr)

		if leng:=len(v.GPU.GPUMem); leng!=0 {
			gpumemr = Decimal( gpumemr / float64(leng))
		} else {
			gpumemr = 0
		}
		
		if leng := len(v.CPU.History); leng == 0 {
			continue
		}
		cpuutil := Cal_Average_float(v.CPU.History) * 100 / float64(v.CPU.Limit)
		
		if cpuutil == 0 && res == 0 && gpumemr == 0 {
			continue
		}
		memutil := Decimal(Cal_Average_int(v.Memory.History) * 100 / float64(v.Memory.Limit) )
		tmp.WriteString(v.Pod+" ")
		tmp.WriteString(fmt.Sprintf("%v %v %v %v %v %v\n", res, cpuutil, gpumemr, v.GPU.GPUUtil[0].MaxR, gpumemmxr, memutil))
		
		
	}
	
	tmp.Close()
}

//不同ResourceType的利用率
func DiffResourceTGPUUti(filepath string) {
	tmp, err := os.Create(filepath)
	
	if err != nil {
		fmt.Println("tmp File creating error", err)
		return
	}
	resourcetymp := make(map[string]([]int))
	for _,v := range Podmp {
		if v.ResourceT == "debug" {
			continue
		}
		tmpr, ok := resourcetymp[v.ResourceT]
		if !ok {
			tmpr = make([]int, 0)
		}
		if v.GPU.NumGPU < 1 {
			continue
		}
		if len(v.GPU.GPUUtil) == 0 {
			tmpr = append(tmpr, 0)
			resourcetymp[v.ResourceT] = tmpr
			continue
		}
		var re float64 = 0
		for _, vv := range v.GPU.GPUUtil {
			re = re+Cal_Average_int(vv.History)
		}
		re = Decimal(re / float64(len(v.GPU.GPUUtil)))
		tmpr = append(tmpr, int(re))
		resourcetymp[v.ResourceT] = tmpr
	}
	
	for k,v := range resourcetymp {
		tmp.WriteString(k)
		for _,vv := range v {
			tmp.WriteString(fmt.Sprintf(" %d", vv))
		}
		tmp.WriteString("\n")
	}
	tmp.Close()
}

//单卡任务利用率
func SCGPUUtui(filepath string) {
	tmp, err := os.Create(filepath)

	if err != nil {
		fmt.Println("tmp File creating error", err)
		return
	}
	scutil := make([]int,102)
	scutig := make([]int,102)
	for _,v := range Podmp {
		if v.ResourceT == "debug" {
			continue
		}
		if v.GPU.NumGPU != 1 {
			continue
		}
		//if v.Endtime - v.Starttime >=7200 {
		//	continue
		//}
		if len(v.GPU.GPUUtil) == 0 {
			if v.Endtime - v.Starttime >= 3600 {
				scutig[0]++
			} else {
				scutil[0]++
			}
			continue
		}
		var res float64 = 0
		for _,vv := range v.GPU.GPUUtil {
	
			res += Cal_Average_int(vv.History)
	
		}
		res = Decimal(res / float64(len(v.GPU.GPUUtil)))
		if v.Endtime - v.Starttime >= 3600 {
			scutig[int(res)]++
		} else {
			scutil[int(res)]++
		}
	}
	sum := 0
	tmp.WriteString("scutil:")
	for i:=0;i<=100;i++ {
		tmp.WriteString(fmt.Sprintf(" %d", scutil[i]))
		sum += scutil[i]
	}
	tmp.WriteString("\n")
	fmt.Println(sum)
	
	sum = 0
	tmp.WriteString("scutig:")
	for i:=0;i<=100;i++ {
		tmp.WriteString(fmt.Sprintf(" %d", scutig[i]))
		sum += scutig[i]
	}
	tmp.WriteString("\n")
	fmt.Println(sum)

	tmp.Close()
}

//计算多卡最大最小利用率
func MCGPUUti(filepath string) {
	tmp, err := os.Create(filepath)

	if err != nil {
		fmt.Println("tmp File creating error", err)
		return
	}

	mcutil := make(map[int]int)
	mcutil[-1]=0
	mcutig := make(map[int]int)
	mcutig[-1]=0
	for _,v := range Podmp {
		if v.ResourceT == "debug" {
			continue
		}
		if v.GPU.NumGPU <= 1 {
			continue
		}
		//if v.Endtime - v.Starttime >=10800 {
		//	continue
		//}
		if len(v.GPU.GPUUtil) == 0 {
			if v.Endtime - v.Starttime >= 3600 {
				mcutig[-1]++
			} else {
				mcutil[-1]++
			}
			continue
		}
		var mn int64 = 102
		var mx int64 = 0
		for _,vv := range v.GPU.GPUUtil {
			mn = MIN(int64(Cal_Average_int(vv.History)),mn).(int64)
			mx = MAX(int64(Cal_Average_int(vv.History)),mx).(int64)
		}
		val :=-1
		if mx != 0 {
			val = int(Decimal(float64(mn)*100/float64(mx)))
		}
		if v.Endtime - v.Starttime >= 3600 {
			if _,ok := mcutig[val];ok {
				mcutig[val]++
			} else {
				mcutig[val]=1
			}
		} else {
			if _,ok := mcutil[val];ok {
				mcutil[val]++
			} else {
				mcutil[val]=1
			}
		}
	
	}
	sum := 0
	tmp.WriteString("mcutil:")
	for i:=-1;i<=100;i++ {
		if _,ok:=mcutil[i];ok {
			tmp.WriteString(fmt.Sprintf(" %d", mcutil[i]))
			sum += mcutil[i]
		} else {
			tmp.WriteString(fmt.Sprintf(" 0"))
		}
	}
	fmt.Println(sum)
	tmp.WriteString("\n")
	
	sum = 0
	tmp.WriteString("mcutig:")
	for i:=-1;i<=100;i++ {
		if _,ok:=mcutil[i];ok {
			tmp.WriteString(fmt.Sprintf(" %d", mcutig[i]))
			sum += mcutig[i]
		} else {
			tmp.WriteString(fmt.Sprintf(" 0"))
		}
	}
	fmt.Println(sum)
	tmp.WriteString("\n")

	tmp.Close()
}

//计算pod的GPU util与CPU util差值分布
func Getgpucpu(filepath string) {
	tmp, err := os.Create(filepath)
	
	if err != nil {
		fmt.Println("tmp File creating error", err)
		return
	}

	mcutil := make(map[int]int)
	mcutig := make(map[int]int)
	
	for _,v := range Podmp {
		if v.ResourceT == "debug" {
			continue
		}
		if v.GPU.NumGPU != 1 {
			continue
		}

		//filter Task Duration
		// if v.Endtime - v.Starttime <=3600 {
		// 	continue
		// }

		var res float64 = 0
		for _,vv := range v.GPU.GPUUtil {
			res += Cal_Average_int(vv.History)
		}
		res = Decimal(res / float64(len(v.GPU.GPUUtil)))

		cpuutil := Cal_Average_float(v.CPU.History)*100 / float64(v.CPU.Limit)
		
		res=res-float64(cpuutil)

		if v.Endtime - v.Starttime >=10800 {
			mcutig[int(res)]++
		} else {
			mcutil[int(res)]++
		}
		
	}
	
	sum := 0
	for i:=-100;i<=100;i++ {
		if _,ok:=mcutil[i];ok {
			tmp.WriteString(fmt.Sprintf(" %d", mcutil[i]))
			sum += mcutil[i]
		} else {
			tmp.WriteString(fmt.Sprintf("0 "))
		}
	}
	fmt.Println(sum)
	tmp.WriteString("\n")

	sum = 0
	for i:=-100;i<=100;i++ {
		if _,ok:=mcutig[i];ok {
			tmp.WriteString(fmt.Sprintf(" %d", mcutig[i]))
			sum += mcutig[i]
		} else {
			tmp.WriteString(fmt.Sprintf("0 "))
		}
	}
	fmt.Println(sum)
	tmp.WriteString("\n")

	tmp.Close()
}

//输出cpu，gpu利用率边界情况&CPU利用率分布
func PrintGPUCPUUtiRange(filepath string) {
	tmp, err := os.Create(filepath)
	
	if err != nil {
		fmt.Println("tmp File creating error", err)
		return
	}

	cpuarr := make(map[int]int)
	for _,v := range Podmp {
		if v.ResourceT == "debug" {
			continue
		}
		if v.GPU.NumGPU != 1 {
			// if v.GPU.NumGPU >= 8 {
			// 	fmt.Println(v.Pod)
			// }
			continue
		}
		if v.Endtime - v.Starttime <=3600 {
			continue
		}

		var res float64 = 0
		for _,vv := range v.GPU.GPUUtil {
			res += Cal_Average_int(vv.History)
		}
		res = Decimal(res / float64(len(v.GPU.GPUUtil)))

		if res < 10 {
			tmp.WriteString(v.Pod)
			cpuutil := Cal_Average_float(v.CPU.History)*100 / float64(v.CPU.Limit)

			tmp.WriteString(fmt.Sprintf(" %v %v", int(res),cpuutil))
			if cpuutil > 90 {
				//fmt.Println(v.Pod+"        01")
			} else if cpuutil < 10 {
				//fmt.Println(v.Pod+"        00")
			}
			cpuarr[int(cpuutil)]++
			
			tmp.WriteString("\n")
		} else if res > 90 {

			cpuutil := Cal_Average_float(v.CPU.History)*100 / float64(v.CPU.Limit)

			tmp.WriteString(fmt.Sprintf(" %v %v", int(res),cpuutil))
			if cpuutil > 90 {
				//fmt.Println(v.Pod+"        11")
			} else if cpuutil < 10 {
				//fmt.Println(v.Pod+"        10")
			}
			cpuarr[int(cpuutil)]++
			
			tmp.WriteString("\n")
		}
		
	}
	for i:=0;i<=100;i++ {
		if _,ok := cpuarr[i];ok {
			tmp.WriteString(fmt.Sprintf(" %d", cpuarr[i]))
		} else {
			tmp.WriteString(" 0")
		}
	}
	tmp.WriteString("\n")
	tmp.Close()
}
