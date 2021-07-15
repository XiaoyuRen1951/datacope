f = open("./data.log", 'r', encoding='utf-8')
while True:
    block = f.readline()
    block = block[:-1]
    
    if block == "basic":
        dic = {'20' : 0, '32' : 1,  '44' : 2, '56' : 3, '68' : 4,  '80' : 5, '92' : 6, '110' : 7}
        
        filecpu = open("./cpubasic.log", 'r', encoding='utf-8')
        s = filecpu.read()
        cpu = s.split("\n")
        filecpu.close()
        for index in range(len(cpu)):
            cpu[index] = [ x for x in cpu[index].split(" ") ]
        
        filegpu = open("./gpubasic.log", 'r', encoding='utf-8')
        s = filegpu.read()
        gpu = s.split("\n")
        filegpu.close()
        for index in range(len(gpu)):
            gpu[index] = [ x for x in gpu[index].split(" ") ]
        
        filemem = open("./membasic.log", 'r', encoding='utf-8')
        s = filemem.read()
        mem = s.split("\n")
        filemem.close()
        for index in range(len(mem)):
            mem[index] = [ x for x in mem[index].split(" ") ]
        
        
        idv = f.readline()
        idv = idv[:-1]
        id = dic[idv]
        gpuv = f.readline()
        gpuv = gpuv[:-1]
        cpuv = f.readline()
        cpuv = cpuv[:-1]
        memv = f.readline()
        memv = memv[:-1]
        gpu[id].append( gpuv )
        cpu[id].append( cpuv )
        mem[id].append( memv )
        
        filecpu = open("./cpubasic.log", 'w', encoding='utf-8')     
        filegpu = open("./gpubasic.log", 'w', encoding='utf-8')
        filemem = open("./membasic.log", 'w', encoding='utf-8')
        filebasic = open("./basic.log", 'w', encoding='utf-8')
        
        filebasic.write("20 32 44 56 68 80 92 110\n")
        
        val = []
        for index in range(8):
            sumval = sum(float(x) for x in gpu[index])
            val.append(round(sumval/float(len(gpu[index])), 2))
            s = ' '.join(gpu[index])
            s.strip()
            filegpu.write(s+"\n")
        filebasic.write((' '.join([str(x) for x in val])).strip()+"\n")

        val = []              
        for index in range(8):
            sumval = sum(float(x) for x in cpu[index])
            val.append(round(sumval/float(len(cpu[index])), 2))
            s = ' '.join(cpu[index])
            s.strip()
            filecpu.write(s+"\n")
        filebasic.write((' '.join([str(x) for x in val])).strip()+"\n")
        
        val = []
        for index in range(8):
            sumval = sum(float(x) for x in mem[index])
            val.append(round(sumval/float(len(mem[index])), 2))
            s = ' '.join(mem[index])
            s.strip()
            filemem.write(s+"\n")
        filebasic.write((' '.join([str(x) for x in val])).strip()+"\n")
        
        filecpu.close()
        filegpu.close()
        filemem.close()
        filebasic.close()
        
    elif block == "bnk":
        dic = {'20' : 0, '29' : 1, '47' : 2, '56' : 3, '74' : 4, '92' : 5, '110' : 6}    

        filecpu = open("./cpubnk.log", 'r', encoding='utf-8')
        s = filecpu.read()
        cpu = s.split("\n")
        filecpu.close()
        for index in range(len(cpu)):
            cpu[index] = [ x for x in cpu[index].split(" ") ]
        
        filegpu = open("./gpubnk.log", 'r', encoding='utf-8')
        s = filegpu.read()
        gpu = s.split("\n")
        filegpu.close()
        for index in range(len(gpu)):
            gpu[index] = [ x for x in gpu[index].split(" ") ]
        
        filemem = open("./membnk.log", 'r', encoding='utf-8')
        s = filemem.read()
        mem = s.split("\n")
        filemem.close()
        for index in range(len(mem)):
            mem[index] = [ x for x in mem[index].split(" ") ]
        
        idv = f.readline()
        idv = idv[:-1]
        id = dic[idv]
        gpuv = f.readline()
        gpuv = gpuv[:-1]
        cpuv = f.readline()
        cpuv = cpuv[:-1]
        memv = f.readline()
        memv = memv[:-1]
        gpu[id].append( gpuv )
        cpu[id].append( cpuv )
        mem[id].append( memv )
        
        filecpu = open("./cpubnk.log", 'w', encoding='utf-8')     
        filegpu = open("./gpubnk.log", 'w', encoding='utf-8')
        filemem = open("./membnk.log", 'w', encoding='utf-8')
        filebnk = open("./bnk.log", 'w', encoding='utf-8')
        
        filebnk.write("20 29 47 56 74 92 110\n")
        
        val = []
        for index in range(7):
            sumval = sum(float(x) for x in gpu[index])
            val.append(round(sumval/float(len(gpu[index])), 2))
            s = ' '.join(gpu[index])
            s.strip()
            filegpu.write(s+"\n")
        filebnk.write((' '.join([str(x) for x in val])).strip()+"\n")

        val = []              
        for index in range(7):
            sumval = sum(float(x) for x in cpu[index])
            val.append(round(sumval/float(len(cpu[index])), 2))
            s = ' '.join(cpu[index])
            s.strip()
            filecpu.write(s+"\n")
        filebnk.write((' '.join([str(x) for x in val])).strip()+"\n")
        
        val = []
        for index in range(7):
            sumval = sum(float(x) for x in mem[index])
            val.append(round(sumval/float(len(mem[index])), 2))
            s = ' '.join(mem[index])
            s.strip()
            filemem.write(s+"\n")
        filebnk.write((' '.join([str(x) for x in val])).strip()+"\n")
        
        filecpu.close()
        filegpu.close()
        filemem.close()
        filebnk.close()
        
    else:
        break

f.close()    