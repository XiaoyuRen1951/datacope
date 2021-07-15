import numpy as np
import matplotlib as mpl
import matplotlib.pyplot as plt

file = open("./basic.log", 'r', encoding='utf-8')
str = file.read()
basic = str.split("\n")
basicx = [float(x) for x in basic[0].split(" ")]
basicgpu = [float(x) for x in basic[1].split(" ")]
basiccpu = [float(x) for x in basic[2].split(" ")]
basicmem = [float(x) for x in basic[3].split(" ")] 
file.close()

file = open("./bnk.log", 'r', encoding='utf-8')
str = file.read()
bnk = str.split("\n")
bnkx = [float(x) for x in bnk[0].split(" ")]
bnkgpu = [float(x) for x in bnk[1].split(" ")]
bnkcpu = [float(x) for x in bnk[2].split(" ")]
bnkmem = [float(x) for x in bnk[3].split(" ")] 
file.close()

def plotbasicbnk( ylabel, basic, bnk, basicval, bnkval, loc):
    _, ax = plt.subplots()

    ax.set_xlabel("Convolutional layer")
    ax.set_ylabel(ylabel)

    ax.set_ylim(0,100)
    #ax.set_ylim(0,12000)
    plt.xticks([20,29,32,44,47,56,68,74,80,92,110],["20", "29", "32", "44", "47", "56", "68", "74", "80", "92", "110"])
    
    plt.plot(basicx, basicval, c='r', marker = "x", markersize=8, label="basicblock")
    for index in range(8):
        tmp = basic[index].split(" ")
        val = [(float(x)) for x in tmp]
        valx = []
        for _ in range(len(val)):
            valx.append(basicx[index])
        ax.scatter(valx, val, c='r', marker='x', s=8)    

    ax.plot(bnkx, bnkval, c='g', marker = "2", markersize=10, label="bottleneck")
    for index in range(7):
        tmp = bnk[index].split(" ")
        val = [(float(x)) for x in tmp]
        valx = []
        for _ in range(len(val)):
            valx.append(bnkx[index])
        ax.scatter(valx, val, c='g', marker='+', s=10)
  
    plt.legend(loc=loc)
    plt.tight_layout()
    plt.show()
    plt.close()


# file = open("./gpubasic.log", 'r', encoding='utf-8')
# str = file.read()
# gpubasic = str.split("\n")
# file.close()

# file = open("./gpubnk.log", 'r', encoding='utf-8')
# str = file.read()
# gpubnk = str.split("\n")
# file.close()

#plotbasicbnk("GPU Utilization/%", gpubasic, gpubnk, basicgpu, bnkgpu, "upper left")

file = open("./cpubasic.log", 'r', encoding='utf-8')
str = file.read()
cpubasic = str.split("\n")
file.close()

file = open("./cpubnk.log", 'r', encoding='utf-8')
str = file.read()
cpubnk = str.split("\n")
file.close()

plotbasicbnk("CPU Utilization/%", cpubasic, cpubnk, basiccpu, bnkcpu, "upper right")


# file = open("./membasic.log", 'r', encoding='utf-8')
# str = file.read()
# membasic = str.split("\n")
# file.close()

# file = open("./membnk.log", 'r', encoding='utf-8')
# str = file.read()
# membnk = str.split("\n")
# file.close()

# for id in range(len(basicmem)):
#     basicmem[id] = float(basicmem[id]) * 32480 / 100
# print(basicmem)    
# for id in range(len(bnkmem)):
#     bnkmem[id] = float(bnkmem[id]) * 32480 / 100


# plotbasicbnk("GPU Memory Occupied/MB", membasic, membnk, basicmem, bnkmem, "upper left")

# file = open("./maxbnk.log", 'r', encoding='utf-8')
# str = file.read()
# maxbnk = str.split("\n")
# valbnk = []
# for i in [0,2,4,6,8,10,12]:
#     valbnk.append([ float(x) for x in maxbnk[i].split(" ") ])
# averagebnk = []
# for i in range(len(valbnk)):
#     averagebnk.append(round(sum(valbnk[i])/float(len(valbnk[i])), 2) )
# file.close()

# file = open("./maxbasic.log", 'r', encoding='utf-8')
# str = file.read()
# maxbasic = str.split("\n")
# valbas = []
# for i in [0,2,4,6,8,10,12,14]:
#     valbas.append([ float(x) for x in maxbasic[i].split(" ") ])
# averagebas = []
# for i in range(len(valbas)):
#     averagebas.append(round(sum(valbas[i])/float(len(valbas[i])), 2) )
# file.close()

# bnkx = [20,29,47,56,74,92,110]
# basx = [20,32,44,56,68,80,92,110]

# _, ax = plt.subplots()
# ax.set_xlabel("Convolutional layer")
# ax.set_ylabel('GPU Utilization/%')

# ax.set_ylim(0,100)
# plt.xticks([20,29,32,44,47,56,68,74,80,92,110],["20", "29", "32", "44", "47", "56", "68", "74", "80", "92", "110"])
# ax.plot(bnkx, averagebnk, c='g', marker = "2", markersize=8, label="bottleneck")
# for index in range(7):
#     valx = []
#     for _ in range(len(valbnk[index])):
#         valx.append(bnkx[index])
#     ax.scatter(valx, valbnk[index], c='g', marker='+', s=8)

# ax.plot(basx, averagebas, c='r', marker = "x", markersize=8, label="basicblock")
# for index in range(8):
#     valx = []
#     for _ in range(len(valbas[index])):
#         valx.append(basx[index])
#     ax.scatter(valx, valbas[index], c='r', marker='x', s=8) 

       
# plt.legend(loc = "lower left")
# plt.tight_layout()
# plt.show()
# plt.close()


# ax.set_title("gpu-cpu")
# ax.set_ylabel('GPU Utilization/%')
# x = ["resnet20", "resnet29", "resnet32", "resnet44", "resnet47", "resnet56", "resnet68", "resnet74", "resnet80", "resnet92", "resnet110"]
# lns1 = ax.plot(x, gpubasic, c='r')
# lns3 = ax.plot(x, gpubottle, c='y')

# ax2 = ax.twinx()
# ax2.set_ylabel('CPU Utilization/%')
# lns2 = ax2.plot(x, cpubasic, c='g')
# lns4 = ax2.plot(x, cpubottle, c='b')

# lns = lns1 + lns2 + lns3 + lns4
# plt.legend(lns, ["BasicBlock GPU Util", "BasicBlock CPU Util", "Bottleneck GPU Util", "Bottleneck CPU Util"], loc = 0)
# plt.tight_layout()
# plt.show()
# #plt.savefig("./gpucpu.jpg")
# plt.close()
