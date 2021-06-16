import numpy as np
import matplotlib as mpl
import matplotlib.pyplot as plt

def fun1():
    file = open("./tmp.log", 'r', encoding='utf-8')
    ss = file.read()
    data = ss.split("\n")
    xx = data[0].split(" ")
    yy = data[1].split(" ")
    tot = []
    for v in yy:
        tot.append(float(v))
    # print(xx)
    # print("\n")
    # print(tot)
    plt.plot(xx, tot)
    file.close()
    plt.xlabel('Time/min')
    plt.ylabel('Task Cnt')

    plt.xscale('log')
    plt.tight_layout()
    
    #plt.legend(loc = "lower right")
    #plt.savefig("./taskcnt.jpg")
    plt.show()

def fun2():
    #plt.title("Single Card TASK Time CDF ")
    plt.xlabel('Time/min')
    plt.ylabel('Occupy/%')
    file = open("./task-1.log", 'r', encoding='utf-8')
    str1 = file.read()
    data1 = str1.split("\n")
    xx1 = data1[0].split(" ")
    yy1 = data1[1].split(" ")
    tot = []
    tasknum = 42450
    for v in yy1 :
        tot.append(float(v)*100/tasknum) 
    
    plt.plot(xx1,tot,c = "orange", label = "1 GPU")
    file.close()

    file = open("./task-2.log", 'r', encoding='utf-8')
    str2 = file.read()
    data2 = str2.split("\n")
    xx2 = data2[0].split(" ")
    yy2 = data2[1].split(" ")
    
    tot = []
    tasknum = 15312
    for v in yy2 :
        tot.append(float(v)*100/tasknum) 
    plt.plot(xx2,tot,linestyle=':', label = "2-4 GPU")
    file.close()

    file = open("./task-3.log", 'r', encoding='utf-8')
    str3 = file.read()
    data3 = str3.split("\n")
    xx3 = data3[0].split(" ")
    yy3 = data3[1].split(" ")
    tot = []
    tasknum = 3311
    for v in yy3 :
        tot.append(float(v)*100/tasknum) 
    plt.plot(xx3,tot,linestyle='-.', label = "5-8 GPU")
    file.close()

    file = open("./task-4.log", 'r', encoding='utf-8')
    str4 = file.read()
    data4 = str4.split("\n")
    xx4 = data4[0].split(" ")
    yy4 = data4[1].split(" ")
    tot = []
    tasknum = 2482
    for v in yy4 :
        tot.append(float(v)*100/tasknum) 
    plt.plot(xx4,tot,linestyle='--', label = "> 8 GPU")
    file.close()

    #plt.xticks([10,100,1000,10000,100000],["10", "100", "1000", "10000", "100000"])
    plt.xscale('log')
    plt.tight_layout()
    
    plt.legend(loc = "lower right")
    plt.vlines(60, 0, 100,  colors = "c",linestyles = "--")
    plt.vlines(1440, 0, 100,  colors = "c",linestyles = "--")
    plt.vlines(10080, 0, 100, colors = "c",linestyles = "--")
    #plt.savefig("./cdf-ALL.jpg")
    plt.show()
    #print(tot)

def fun3():
    xx = [0,1]
    yy = [23, 66]
    plt.ylabel('Node Cnt')
    plt.xticks(xx,[">= 75%", "<= 25%"])
    plt.bar(xx, yy)
    plt.savefig("./Nodecnt.jpg")

fun2()