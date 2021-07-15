import numpy as np
import matplotlib as mpl
import matplotlib.pyplot as plt

def tmp():
    file = open("./tmp-1.log", 'r', encoding='utf-8')
    str = file.read()
    data = str.split(" ")
    tot = []
    cnt = []
    mus = 0.0
    tasknum = 7433
    for v in data :
        #mus = mus + float(v)
        #cnt.append(int(v))
        tot.append(int(v)*100/tasknum)
    _, ax = plt.subplots()
    #plt.title("Greater 2h Single Card TASK CDF ")
    ax.set_xlabel('Time/min')
    ax.set_ylabel('Task Occupied/%')
    #ax.set_ylim(bottom=0)
    #cnt[0]=550

    xx = np.arange(len(data))

    plt.plot(xx, tot)

    plt.xscale('log')
    plt.tight_layout()
    #plt.show()
    #print(tot)
    plt.savefig("./111111.jpg")
    file.close()

tmp()
