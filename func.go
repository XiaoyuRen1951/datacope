package main

import (
	"fmt"
	"os"
	"math"
	"reflect"
	//"sort"
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

func Readtaskinfo(path string) error {
	File, err := os.Open(path+"taskinfo.csv")
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
		
		if val, ok := Podmp[line[0]]; ok {
			tmpStarttime, _ := strconv.ParseInt(line[6], 10, 64)
			tmpEndtime, _ := strconv.ParseInt(line[7], 10, 64)
			val.Starttime = MIN(val.Starttime, tmpStarttime).(int64)
			val.Endtime = MAX(val.Endtime, tmpEndtime).(int64)
			Podmp[line[0]] = val
			continue
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
		res.GPU.GPUMemCopy = make([]GPUHistory, 0)
		
		Podmp[res.Pod] = res
	}
	return nil
}

func Readcpuinfo(path string) error {
	File, err := os.Open(path+"cpuinfo.json")
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
			val.CPU.History = append(val.CPU.History , res.History...)
			Podmp[res.Pod] = val
			
		}
	}
	return nil
}

func Readgpuinfomem(path string) error {
	File, err := os.Open(path+"gpuinfomem.json")
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
			fg := true
			for k, _ := range val.GPU.GPUMem {
				if val.GPU.GPUMem[k].Uuid == res.Uuid {
					fg = false
					val.GPU.GPUMem[k].MaxR = MAX(val.GPU.GPUMem[k].MaxR, res.MaxR).(int64)
					val.GPU.GPUMem[k].History = append(val.GPU.GPUMem[k].History, res.History...)
					Podmp[res.Pod] = val
					break
				}
			}
			if fg {
				val.GPU.GPUMem = append(val.GPU.GPUMem, *res)
				Podmp[res.Pod] = val
			}	
		}
	}
	return nil
}

func Readgpuinfoutil(path string) error {
	File, err := os.Open(path + "gpuinfoutil.json")
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

			fg := true
			for k, _ := range val.GPU.GPUUtil {
				if val.GPU.GPUUtil[k].Uuid == res.Uuid {
					fg = false
					val.GPU.GPUUtil[k].MaxR = MAX(val.GPU.GPUUtil[k].MaxR, res.MaxR).(int64)
					val.GPU.GPUUtil[k].History = append(val.GPU.GPUUtil[k].History, res.History...)
					Podmp[res.Pod] = val
					break
				}
			}
			if fg {
				val.GPU.GPUUtil = append(val.GPU.GPUUtil, *res)
				Podmp[res.Pod] = val
			}
			
		}
	}
	return nil
}

func Readgpuinfopcie(path string) error {
	File, err := os.Open(path + "pcie.json")
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
		
		res := new(GPUPCIEHistory)
		prdec := json.NewDecoder(strings.NewReader(line))
		err = prdec.Decode(&res)
		if err != nil {
			fmt.Println(err)
			return err
		}
		
		if val, ok := Podmp[res.Pod]; ok {
			//Deal With GPU PCIE

			fg := true
			for k, _ := range val.GPU.GPUPCIE {
				if val.GPU.GPUPCIE[k].Uuid == res.Uuid {
					fg = false
					val.GPU.GPUPCIE[k].RXMaxR = MAX(val.GPU.GPUPCIE[k].RXMaxR, res.RXMaxR).(float64)
					val.GPU.GPUPCIE[k].RXHistory = append(val.GPU.GPUPCIE[k].RXHistory, res.RXHistory...)
					val.GPU.GPUPCIE[k].TXMaxR = MAX(val.GPU.GPUPCIE[k].TXMaxR, res.TXMaxR).(float64)
					val.GPU.GPUPCIE[k].TXHistory = append(val.GPU.GPUPCIE[k].TXHistory, res.TXHistory...)
					Podmp[res.Pod] = val
					break
				}
			}
			if fg {
				val.GPU.GPUPCIE = append(val.GPU.GPUPCIE, *res)
				Podmp[res.Pod] = val
			}
			
		}
	}
	return nil
}

func Readgpumemcpyutil(path string) error {
	File, err := os.Open(path + "gpumemcpyutil.json")
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
			//Deal With GPU Memory Utilization
			fg := true
			for k, _ := range val.GPU.GPUMemCopy {
				if val.GPU.GPUMemCopy[k].Uuid == res.Uuid {
					fg = false
					val.GPU.GPUMemCopy[k].MaxR = MAX(val.GPU.GPUMemCopy[k].MaxR, res.MaxR).(int64)
					val.GPU.GPUMemCopy[k].History = append(val.GPU.GPUMemCopy[k].History, res.History...)
					Podmp[res.Pod] = val
					break
				}
			}
			if fg {
				val.GPU.GPUMemCopy = append(val.GPU.GPUMemCopy, *res)
				Podmp[res.Pod] = val
			}
			
		}
	}
	return nil
}

func Readmeminfo(path string) error {
	File, err := os.Open(path + "meminfo.json")
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
			val.Memory.History = append(val.Memory.History, res.History...)
			Podmp[res.Pod] = val
			//fmt.Println(val.Memory)
		}
	}
	return nil
}

// func Readnodegpuinfo(path string, vval *[]NodeGPUstate) error {
// 	File, err := os.Open(path + "nodegpuinfo.json")
// 	defer File.Close()

// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	rd := bufio.NewReader(File)

// 	for {
// 		line, err := rd.ReadString('\n')
// 		if err != nil || io.EOF == err {
// 			break
// 		}

// 		if line[0] != '{' {
// 			continue
// 		}
		
// 		res := new(NodeGPUstate)
// 		prdec := json.NewDecoder(strings.NewReader(line))
// 		err = prdec.Decode(&res)
// 		if err != nil {
// 			fmt.Println(err)
// 			return err
// 		}
// 		fg := true
// 		val := *vval
// 		for k, _ := range val {
// 			if val[k].Node == res.Node {
// 				fg = false
// 				val[k].State = append(val[k].State, res.State...)
// 				break
// 			}
// 		}
// 		if fg {
// 			*vval = append(*vval,*res)
// 		}
// 	}
// 	return nil
// }

func Readnodegpuuse(path string, vval *[]Nodegpuuse) error {
	File, err := os.Open(path + "nodegpuuse.csv")
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
		fg := true
		val := *vval
		for k, _ := range val {
			if val[k].Pod == line[0] && val[k].Uuid == line[1] {
				fg = false
				tmp, _ := strconv.ParseFloat(line[2], 64) 
				val[k].Ratio = (val[k].Ratio + tmp ) / 2
				break
			}
		}
		if !fg {
			continue
		}

		var tmp Nodegpuuse
		tmp.Pod = line[0]
		tmp.Uuid = line[1]
		tmp.Ratio, _ = strconv.ParseFloat(line[2], 64) 
		
		*vval = append(*vval, tmp)
	}
	return nil
}

func Readnodecpuuti(path string, vval *[]CPUCore) error {
	File, err := os.Open(path + "nodecpuuti.json")
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
		fg := true
		val := *vval
		for k, _ := range val {
			if val[k].Pod == res.Pod {
				fg = false
				val[k].Utilization = append(val[k].Utilization, res.Utilization...)
				break
			}
		}
		if fg {
			*vval = append(*vval, *res)
		}
		
	}
	return nil
}

func Readnodeio(path string, val map[string]NodeIO) error {
	File, err := os.Open(path + "nodereiverate.json")

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
		
		if v, ok := val[res.Node]; ok {
			for kk,vv := range res.IRate {
				v.IRate[kk] = vv
			}
			
			val[res.Node] = v
		} else {
			val[res.Node] = *res
		}

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
			for kk,vv := range res.ORate {
				v.ORate[kk] = vv
			}
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
			for kk,vv := range res.IbIRate {
				v.IbIRate[kk] = vv
			}
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
			for kk,vv := range res.IbORate {
				v.IbORate[kk] = vv
			}
			val[res.Node] = v
		} else {
			val[res.Node] = *res
		}
	}
	File.Close()

	return nil
}

func CalGPUWaste(filepath string) {
	tmp, err := os.Create(filepath)
	
	if err != nil {
		fmt.Println("renxy-task File creating error", err)
		return
	}

	abc := make([]int, 105)
	cba := make([]int, 105)
	for _,v := range Podmp {
		if strings.Contains(v.ResourceT,"debug") == true {
			continue
		}
		if v.Endtime - v.Starttime < 60 {
			continue
		}

		var res float64 = 0

		for _,vv := range v.GPU.GPUUtil {
			val := 0
			for _, vvv := range vv.History {
				if vvv > 0 {
					val ++
				}
			}
			if leng := len(vv.History); leng != 0 {
				res += (float64(leng - val) * 100 / float64(leng))
			}
			
		}
		if leng:=len(v.GPU.GPUUtil); leng!=0 {
			res = Decimal(res / float64(leng))
		} else {
			res = 0
		}

		res = 100 - res
		if v.GPU.NumGPU == 1 {
			abc[int(res)]++
		} else {
			cba[int(res)]++
		}
		
	}
	sum := 0
	
	for i:=0;i<=100;i++ {
		tmp.WriteString(fmt.Sprintf(" %d", abc[i]))
		sum += abc[i]
	}
	fmt.Println(sum)
	tmp.WriteString("\n")

	sum = 0
	for i:=0;i<=100;i++ {
		tmp.WriteString(fmt.Sprintf(" %d", cba[i]))
		sum += cba[i]
	}
	fmt.Println(sum)
	tmp.WriteString("\n")
	
	tmp.Close()
}

func CalTime() {
	var time int64 = 0
	var mxgpu int64 = 0
	var gpucnt int64 = 0
	var mxt int64 = 0
	cput := 0
	gput := 0

	for _,v := range Podmp {
		if strings.Contains(v.ResourceT,"debug") == true {
			continue
		}
		mxgpu = MAX(mxgpu, v.GPU.NumGPU).(int64)
		if v.GPU.NumGPU == 0 {
			cput++
		} else {
			gput++
		}
		gpucnt += v.GPU.NumGPU
		time += v.Endtime - v.Starttime
		mxt = MAX(mxt, v.Endtime - v.Starttime).(int64)
	}
	fmt.Println(cput,gput, mxgpu)
	fmt.Println(Decimal(float64(mxt)/3600/24))
	fmt.Println(Decimal(float64(time)/float64(cput+gput)))
	fmt.Println(Decimal(float64(gpucnt)/float64(cput+gput)))

}

func AverageVal(filepath string) {
	tmp, err := os.Create(filepath)
	
	if err != nil {
		fmt.Println("renxy-task File creating error", err)
		return
	}
	
	abc := make([]int, 205)
	cba := make([]int, 205)
	cab := make([]int, 100)
	for _,v := range Podmp {
		if strings.Contains(v.ResourceT,"debug") == true {
			continue
		}
		
		if v.Endtime - v.Starttime < 60 {
			continue
		}

		t := v.Endtime - v.Starttime
		tnc := 0
		for t > 0 {
			tnc++
			t= t/10
		}
		cab[tnc]++
		// if strings.Contains(v.Pod,"renxy") != true {
		// 	continue
		// }
		
		//filter Task Duration
		// if v.Endtime - v.Starttime <=3600 {
		// 	continue
		// }
		
		// var res float64 = 0
		// for _,vv := range v.GPU.GPUUtil {
			
		// 	tmp := Cal_Average_int(vv.History)
		// 	res += tmp
		// }
		// if leng:=len(v.GPU.GPUUtil); leng!=0 {
		// 	res = Decimal(res / float64(leng))
		// } else {
		// 	res = 0
		// }



		var res float64 = 0
		for _,vv := range v.GPU.GPUPCIE {
			res += (Cal_Average_float(vv.RXHistory) + Cal_Average_float(vv.TXHistory)) * 200 / 1024 / 1024 / 15.754
		}
		if leng:=len(v.GPU.GPUPCIE); leng!=0 {
			res = Decimal(res / float64(leng))
		} else {
			res = 0
		}

		if v.GPU.NumGPU == 1 {
			abc[int(res)] ++
		} else {
			cba[int(res)] ++
		}
		

		// var gpumemr float64 = 0
		// var gpumemmxr float64 = 0

		// for _, vv := range v.GPU.GPUMem {
		// 	gpumemr += Cal_Average_int(vv.History) * 100 / float64(vv.Total)
		// 	gpumemmxr = MAX(gpumemmxr, float64(vv.MaxR*100) / float64(vv.Total)).(float64)
		// }
		
		// gpumemmxr = Decimal(gpumemmxr)

		// if leng:=len(v.GPU.GPUMem); leng!=0 {
		// 	gpumemr = Decimal( gpumemr / float64(leng))
		// 	gpumemr = MIN(gpumemr,100.0).(float64)
		// } else {
		// 	gpumemr = 0
		// }

		// abc[int(gpumemr)] ++
		// cba[int(gpumemr)] ++
		
		// if leng := len(v.CPU.History); leng == 0 {
		// 	continue
		// }
		// cpuutil := Cal_Average_float(v.CPU.History) * 100 / float64(v.CPU.Limit)
		
		// if cpuutil == 0 && res == 0 && gpumemr == 0 {
		// 	continue
		// }

		
		// if len(v.Memory.History) == 0 {
		// 	if v.GPU.NumGPU == 1 {
		// 		abc[0] ++
		// 	} else {
		// 		cba[0] ++
		// 	}
		// 	continue
		// }
		// var res float64 = Cal_Average_int(v.Memory.History) * 100 / float64(v.Memory.Limit)

		// if v.GPU.NumGPU == 1 {
		// 	abc[int(res)] ++
		// } else {
		// 	cba[int(res)] ++
		// }
		
		// tmp.WriteString(v.Pod+" ")
		// //tmp.WriteString(fmt.Sprintf("%v %v %v %v %v %v %v %v\n", res, cpuutil, gpumemr, v.GPU.GPUUtil[0].MaxR, gpumemmxr, memutil, rx, tx))
		// if len(v.GPU.GPUPCIE) > 0 {
		// 	tmp.WriteString(fmt.Sprintf("%v%% %v%% %v%% %v%%", rx, Decimal(v.GPU.GPUPCIE[0].RXMaxR * 100 / 1024 / 1024 / 15.754 ),tx, Decimal(v.GPU.GPUPCIE[0].TXMaxR * 100 / 1024 / 1024 / 15.754)))
		// 	//tmp.WriteString(fmt.Sprintf("%v%% %v%% ", rx,tx))
		// }
		// tmp.WriteString("\n")
		
	}
	
	sum := 0
	tmp.WriteString("rx:")
	for i:=0;i<=200;i++ {
		tmp.WriteString(fmt.Sprintf(" %d", abc[i]))
		sum += abc[i]
	}
	fmt.Println(sum)
	tmp.WriteString("\n")

	sum = 0
	tmp.WriteString("tx:")
	for i:=0;i<=200;i++ {
		tmp.WriteString(fmt.Sprintf(" %d", cba[i]))
		sum += cba[i]
	}
	fmt.Println(sum)
	tmp.WriteString("\n")
	for i:=0; i < 10 ;i ++ {
		fmt.Printf(" %v", cab[i])
	}
	fmt.Printf("\n")
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
	
	durmp := make(map[int]int)

	sum := 0
	for _,v := range Podmp {
		if strings.Contains(v.ResourceT,"debug") == true {
			continue
		}
		
		//if v.GPU.NumGPU != 1 || v.Endtime - v.Starttime < 60 {
		if v.GPU.NumGPU == 1 || v.Endtime - v.Starttime < 60 {
			continue
		}
		
		// if v.Endtime - v.Starttime >=7200 {
		// 	continue
		// }
		if len(v.GPU.GPUUtil) == 0 {
			scutil[0]++
			continue
		}
		var res float64 = 0
		for _,vv := range v.GPU.GPUUtil {
	
			res += Cal_Average_int(vv.History)
	
		}
		res = Decimal(res / float64(len(v.GPU.GPUUtil)))
		scutil[int(res)]++
		if int(res) == 0 {
			duration := int(math.Ceil(float64(v.Endtime - v.Starttime) / 60))
			if d, ok := durmp[duration]; ok {
				durmp[duration] = d+1
			} else {
				durmp[duration] = 1
			}
			sum = sum + 1
		}
		// if durmp[v.Endtime - v.Starttime] ;res == 0 {
		// 	tmp.WriteString(fmt.Sprintf(" %v", v.Endtime - v.Starttime))
		// }
	}
	k := 0
	for i := 0; i<100000; i++ {
		if k >= sum {
			break
		}
		if j,ok := durmp[i]; ok {
			k=k+j
			
		}
		tmp.WriteString(fmt.Sprintf(" %v", k))
	}
	// tmp.WriteString("scutil:")
	// for i:=0;i<=100;i++ {
	// 	tmp.WriteString(fmt.Sprintf(" %d", scutil[i]))
	// 	sum += scutil[i]
	// }
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
	
	for _,v := range Podmp {
		if strings.Contains(v.ResourceT,"debug") == true {
			continue
		}
		if v.GPU.NumGPU <= 1 || v.Endtime - v.Starttime < 60 {
			continue
		}
		//if v.Endtime - v.Starttime >=10800 {
		//	continue
		//}
		if len(v.GPU.GPUUtil) == 0 {
			mcutil[-1]++
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
		if _,ok := mcutil[val];ok {
			mcutil[val]++
		} else {
			mcutil[val]=1
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
	scutil := make(map[int]int)
	
	for _,v := range Podmp {
		if strings.Contains(v.ResourceT,"debug") == true {
			continue
		}
		//if v.GPU.NumGPU != 1 || v.Endtime - v.Starttime < 60 {
		if v.Endtime - v.Starttime < 60 {
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

		if v.GPU.NumGPU == 1 {
			scutil[int(res)]++
		} else {
			mcutil[int(res)]++
		}
		
		
	}
	
	sum := 0
	for i:=-100;i<=100;i++ {
		if _,ok:=scutil[i];ok {
			tmp.WriteString(fmt.Sprintf(" %d", scutil[i]))
			sum += scutil[i]
		} else {
			tmp.WriteString(fmt.Sprintf(" 0"))
		}
	}
	fmt.Println(sum)
	tmp.WriteString("\n")

	sum = 0
	for i:=-100;i<=100;i++ {
		if _,ok:=mcutil[i];ok {
			tmp.WriteString(fmt.Sprintf(" %d", mcutil[i]))
			sum += mcutil[i]
		} else {
			tmp.WriteString(fmt.Sprintf(" 0"))
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

	//cpuarr := make(map[int]int)

	glutil := make([]int,102)
	ggutil := make([]int,102)

	for _,v := range Podmp {
		if strings.Contains(v.ResourceT,"debug") == true {
			continue
		}
		if v.GPU.NumGPU == 1 || v.Endtime - v.Starttime < 60 {
		//if v.GPU.NumGPU == 1 || v.Endtime - v.Starttime < 60 {
			// if v.GPU.NumGPU >= 8 {
			// 	fmt.Println(v.Pod)
			// }
			continue
		}
		// if v.Endtime - v.Starttime <=3600 {
		// 	continue
		// }

		var res float64 = 0
		for _,vv := range v.GPU.GPUUtil {
			res += Cal_Average_int(vv.History)
		}
		res = Decimal(res / float64(len(v.GPU.GPUUtil)))
		cpuutil := MIN(Cal_Average_float(v.CPU.History)*100 / float64(v.CPU.Limit), 100.0).(float64)
		if res < 20 {
			glutil[int(cpuutil)]++
		} else if res > 80 {
			ggutil[int(cpuutil)]++
		}

		// if res < 10 {
		// 	//tmp.WriteString(v.Pod)
		// 	cpuutil := MIN(Cal_Average_float(v.CPU.History)*100 / float64(v.CPU.Limit), 100.0).(float64)

		// 	//tmp.WriteString(fmt.Sprintf(" %v %v\n", int(res),cpuutil))
		// 	if cpuutil > 90 {
		// 		//fmt.Println(v.Pod+"        01")
		// 	} else if cpuutil < 10 {
		// 		//fmt.Println(v.Pod+"        00")
		// 	}
		// 	// fmt.Println(fmt.Sprintf("%v %v %v\n",v.Pod, Cal_Average_float(v.CPU.History),cpuutil))
		// 	cpuarr[int(cpuutil)]++
			
			// if v.Endtime - v.Starttime <=7200 {
			// 	glutil[int(cpuutil)]++
			// } else {
			// 	glutig[int(cpuutil)]++
			// }
		// } else if res > 90 {

		// 	cpuutil := MIN(Cal_Average_float(v.CPU.History)*100 / float64(v.CPU.Limit), 100.0).(float64)

		// 	//tmp.WriteString(fmt.Sprintf(" %v %v\n", int(res),cpuutil))
		// 	if cpuutil > 90 {
		// 		//fmt.Println(v.Pod+"        11")
		// 	} else if cpuutil < 10 {
		// 		//fmt.Println(v.Pod+"        10")
		// 	}
		// 	cpuarr[int(cpuutil)]++
			
			// if v.Endtime - v.Starttime <=7200 {
			// 	ggutil[int(cpuutil)]++
			// } else {
			// 	ggutig[int(cpuutil)]++
			// }
		// }
		
	}
	// for i:=0;i<=100;i++ {
	// 	if _,ok := cpuarr[i];ok {
	// 		tmp.WriteString(fmt.Sprintf(" %d", cpuarr[i]))
	// 	} else {
	// 		tmp.WriteString(" 0")
	// 	}
	// }
	// tmp.WriteString("\n")
	
	sum := 0
	tmp.WriteString("glutil:")
	for i:=0;i<=100;i++ {
		tmp.WriteString(fmt.Sprintf(" %d", glutil[i]))
		sum += glutil[i]
	}
	tmp.WriteString("\n")
	fmt.Println(sum)
	
	sum = 0
	tmp.WriteString("ggutil:")
	for i:=0;i<=100;i++ {
		tmp.WriteString(fmt.Sprintf(" %d", ggutil[i]))
		sum += ggutil[i]
	}
	tmp.WriteString("\n")
	fmt.Println(sum)
	
	tmp.Close()
}

func CalTaskCnt(filepath string, ll, rr int64) {
	
	taskcnt := make(map[int]int64)
	x := 0
	gap := 1
	cnt := 0
	for _,v := range Podmp {
		if strings.Contains(v.ResourceT,"debug") == true {
			continue
		}
		//if !(v.GPU.NumGPU == 1) || v.Endtime - v.Starttime < 60 {
		if !(v.GPU.NumGPU >= ll && v.GPU.NumGPU <= rr) || v.Endtime - v.Starttime < 60 {
			cnt ++
			continue
		}
		
		duration := int(math.Ceil(float64(v.Endtime - v.Starttime) / 60))
		if _, ok := taskcnt[duration/gap]; ok {
			taskcnt[duration/gap]++
		} else {
			taskcnt[duration/gap] = 1
			if duration/gap > x {
				x = duration/gap
			}
		}
			
	}
	
	fmt.Println(cnt)
	tmp, err := os.Create(filepath)
	
	if err != nil {
		fmt.Println("tmp File creating error", err)
		return
	}

	for i:= 0; i <= x; i++ {
		tmp.WriteString(fmt.Sprintf(" %v", i*gap))
	}
	tmp.WriteString("\n")
	var sum int64 = 0
	for v:=0; v <= x; v++ {
		sum = 0
		if _,ok := taskcnt[v]; ok {
			//sum += taskcnt[v]
			sum = taskcnt[v]
		}
		tmp.WriteString(fmt.Sprintf(" %v", sum))
		
	}
	//fmt.Println(sum)
	tmp.WriteString("\n")
}

func CalNodeBusy() {
	x := -1
	for _, v := range NodeGPU {
		if x == -1 || x > len(v.State) {
			x = len(v.State)
		}
	}
	fmt.Println(x)
	ll := 0
	rr := 0
	pos := -1
	l0 := 0
	r0 := 0
	for i:=0; i < x; i++ {
		l := 0
		r := 0
		ll0 := 0
		rr0 := 0
		for _, v := range NodeGPU {
			uti := float64(v.State[i].Use) * 100.0 / float64(v.State[i].Total)
			if uti <= 25 {
				l ++
			} else if uti >= 75 {
				r ++
			}
			if int(uti) == 0 {
				ll0 ++
			} else if int(uti) == 100 {
				rr0 ++
			}
		}
		if pos == -1 || l > ll {
			ll = l
			rr = r
			pos = i
			l0 = ll0
			r0 = rr0
		} else if r > rr {
			ll = l
			rr = r
			pos = i
			l0 = ll0
			r0 = rr0
		}
	}
	fmt.Println(fmt.Sprintf("%v %v %v %v %v", pos, ll, rr, l0, r0))

	cn := make([][4]int, 30)
	for j:=0; j<24;j++ {
		//i := 0 + j * 2880
		i := 0 + j * 120

		for _, v := range NodeGPU {
			i = int(MIN(int64(i),int64(len(v.State)-1)).(int64))
			uti := float64(v.State[i].Use) * 100.0 / float64(v.State[i].Total)
			if int(uti) == 0 {
				cn[j][0]++
			}
			if uti <= 25 {
				cn[j][1]++
			}
			if uti <= 100 {
				cn[j][3]++
			}
			if uti <= 75 {
				cn[j][2]++
			}
			
		}
		
	}
	for i:= 0; i<4; i++ {
		for j := 0; j<24; j++ {
			fmt.Printf(", %v", cn[j][i])
		}
		fmt.Printf("\n")
	}
}

func Readnodegpuinfo(path string) {
	File, err := os.Open(path + "nodegpuinfo.json")
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

		if line[0] != '{' {
			continue
		}
		
		res := new(NodeGPUstate)
		prdec := json.NewDecoder(strings.NewReader(line))
		err = prdec.Decode(&res)
		if err != nil {
			fmt.Println(err)
			return
		}
		if v, ok := NodeGPU[res.Node]; ok {
			v.State = append(v.State, res.State...)
			NodeGPU[res.Node] = v
		} else {
			NodeGPU[res.Node] = *res
		}
		
	}

	return
}

func TaskMemUtui(filepath string) {
	tmp, err := os.Create(filepath)

	if err != nil {
		fmt.Println("tmp File creating error", err)
		return
	}
	scutil := make([]int,102)
	
	sum := 0
	for _,v := range Podmp {
		if strings.Contains(v.ResourceT,"debug") == true {
			continue
		}
		
		//if v.GPU.NumGPU != 1 || v.Endtime - v.Starttime < 60 {
		if v.GPU.NumGPU == 1 || v.Endtime - v.Starttime < 60 {
			continue
		}
		
		// if v.Endtime - v.Starttime >=7200 {
		// 	continue
		// }
		if len(v.Memory.History) == 0 {
			scutil[0]++
			continue
		}
		var res float64 = Cal_Average_int(v.Memory.History) * 100 / float64(v.Memory.Limit)

		scutil[int(res)]++
		
		// if durmp[v.Endtime - v.Starttime] ;res == 0 {
		// 	tmp.WriteString(fmt.Sprintf(" %v", v.Endtime - v.Starttime))
		// }
	}
	
	tmp.WriteString("scutil:")
	for i:=0;i<=100;i++ {
		tmp.WriteString(fmt.Sprintf(" %d", scutil[i]))
		sum += scutil[i]
	}
	tmp.WriteString("\n")
	fmt.Println(sum)

	tmp.Close()
}