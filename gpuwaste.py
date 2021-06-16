import numpy as np
import matplotlib as mpl
import matplotlib.pyplot as plt

def tmp():
    file = open("./gpuwate.log", 'r', encoding='utf-8')
    str = file.read()
    src = str.split("\n")
    tot = []
    cnt = []
    mus = 0.0
    tasknum = 31660
    #tasknum = 42450
    data = src[0].split(" ")
    for v in data :
        mus = mus + float(v)
        #cnt.append(int(v))
        tot.append(mus*100/tasknum)
    _, ax = plt.subplots()
    
    ax.set_xlabel('PCIe Bandwidth Occupied/%')
    ax.set_ylabel('Task Occupied/%')
    #ax.set_ylim(bottom=0)
    
    xx = np.arange(len(data))

    plt.plot(xx, tot, label = "RX")


    #plt.xscale('log')
    plt.legend()
    plt.tight_layout()
    plt.show()
    #print(tot)
    #plt.savefig("./111111.jpg")
    file.close()

def cdf():
    file = open("./gpuwate.log", 'r', encoding='utf-8')
    str = file.read()
    src = str.split("\n")
    tot = []
    cnt = []
    mus = 0.0
    tasknum = 42450
    #tasknum = 31660
    data = src[1].split(" ")
    for v in data :
        mus = mus + float(v)
        cnt.append(int(v))
        tot.append(mus*100/tasknum)
    _, ax = plt.subplots()
    #plt.title("Greater 3h Single Card TASK CDF ")
    ax.set_xlabel('GPU Occupancy/%')
    ax.set_ylabel('Task Count')
    #ax.set_ylim(bottom=0)
    # cnt[0]=650
    # cnt[100] = 650
    #cnt[0]=6600
    cnt[100] = 6600

    xx = np.arange(len(data))
    ax.bar(xx,cnt)
    # plt.text(xx[0],cnt[0],12420,ha='center')
    # plt.text(xx[100],cnt[100],6297,ha='center')
    plt.text(xx[0],cnt[0],9085,ha='center')
    #plt.text(xx[100],cnt[100],11250,ha='center')
    plt.xticks(rotation=45)
    ax2=ax.twinx()
    ax2.set_ylabel('Occupy/%')
    ax2.set_ylim(bottom=0,top=110)
    #ax2.set_ylim(bottom=0)
    ax2.plot(xx,tot,c="orange")

    plt.tight_layout()
    #plt.savefig("./plot/cdf-ALL.jpg")
    plt.show()
    #print(tot)
    file.close()

#tmp()
cdf()

