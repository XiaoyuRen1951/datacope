import numpy as np
import matplotlib as mpl
import matplotlib.pyplot as plt

def gpuutili():
    #GPU Util
    a1 = [88.26, 88.77, 88.58]
    a2 = [95.89, 96.65]
    a3 = [97.24, 96.04, 97.02, 97.19]

    aa = []
    aa.append(a1)
    aa.append(a2)
    aa.append(a3)

    b1 = [80.25, 77.46, 77.21]
    b2 = [92.86, 92.73]
    b3 = [95.42, 95.12, 95.48, 95.36]

    bb = []
    bb.append(b1)
    bb.append(b2)
    bb.append(b3)

    c1 = [71.58, 71.32, 70.96]
    c2 = [88.64, 89.8, 89.67]
    c3 = [93.96, 93.96, 93.8]

    cc = []
    cc.append(c1)
    cc.append(c2)
    cc.append(c3)
    
    _, ax = plt.subplots()

    ax.set_xlabel('Hidden Units')
    ax.set_ylabel('GPU Utilization/%')
    ax.set_ylim(0,100)
    xx = [650, 1800, 2500]
    for index in range(3):
        valx = []
        for _ in range(len(aa[index])):
            valx.append(xx[index])
        ax.scatter(valx, aa[index], c='g', marker='+', s=8)

    for index in range(3):
        valx = []
        for _ in range(len(bb[index])):
            valx.append(xx[index])
        ax.scatter(valx, bb[index], c='r', marker='+', s=8)
        
    for index in range(3):
        valx = []
        for _ in range(len(cc[index])):
            valx.append(xx[index])
        ax.scatter(valx, cc[index], c='y', marker='+', s=8)

    xx = [650, 900, 1200, 1500, 1800, 2200, 2500]
    #plt.xticks(xx,["650", "1800", "2500"])

    avga = []
    for idx in range(3):
        avga.append(np.average(aa[idx]))


    avgb = []
    for idx in range(3):
        avgb.append(np.average(bb[idx]))

    avgc = []
    for idx in range(3):
        avgc.append(np.average(cc[idx]))

    avga.insert(2, 96.74) #2200
    avga.insert(1, 94.53) #1500
    avga.insert(1, 92.67) #1200
    avga.insert(1, 91.93) #900
    avgb.insert(2, 94.35)
    avgb.insert(1, 90.04)
    avgb.insert(1, 86.48)
    avgb.insert(1, 83.4)
    avgc.insert(2, 91.78)
    avgc.insert(1, 85.15)
    avgc.insert(1, 80.77)
    avgc.insert(1, 77.8)

    ax.plot(xx, avga, c='g', marker = "2", markersize=8, label="2 layers")
    ax.plot(xx, avgb, c='r', marker = "x", markersize=8, label="4 layers")
    ax.plot(xx, avgc, c='y', marker = "4", markersize=8, label="8 layers")
        
    plt.legend(loc = "lower left")
    plt.tight_layout()
    plt.show()
    plt.close()

def cpuutil():
    #CPU Util
    a1 = [23.75, 23.75,24]
    a2 = [24.0, 24.0]
    a3 = [24.25, 24.25, 24.25]

    aa = []
    aa.append(a1)
    aa.append(a2)
    aa.append(a3)

    b1 = [24.0, 24.5, 24.25]
    b2 = [24.0, 24.25]
    b3 = [24.0, 24.0, 23.75, 24.25]

    bb = []
    bb.append(b1)
    bb.append(b2)
    bb.append(b3)

    c1 = [24.5, 23.5, 24.25]
    c2 = [24.0, 24.5, 23.75]
    c3 = [23.75, 24.0, 24.5]

    cc = []
    cc.append(c1)
    cc.append(c2)
    cc.append(c3)
    
    _, ax = plt.subplots()

    ax.set_xlabel('Hidden Units')
    ax.set_ylabel('CPU Utilization/%')
    ax.set_ylim(0,100)
    xx = [650, 1800, 2500]
    for index in range(3):
        valx = []
        for _ in range(len(aa[index])):
            valx.append(xx[index])
        ax.scatter(valx, aa[index], c='g', marker='+', s=8)

    for index in range(3):
        valx = []
        for _ in range(len(bb[index])):
            valx.append(xx[index])
        ax.scatter(valx, bb[index], c='r', marker='+', s=8)
        
    for index in range(3):
        valx = []
        for _ in range(len(cc[index])):
            valx.append(xx[index])
        ax.scatter(valx, cc[index], c='y', marker='+', s=8)

    xx = [650, 900, 1200, 1500, 1800, 2200, 2500]
    #plt.xticks(xx,["650", "1800", "2500"])

    avga = []
    for idx in range(3):
        avga.append(np.average(aa[idx]))
        
    avgb = []
    for idx in range(3):
        avgb.append(np.average(bb[idx]))

    avgc = []
    for idx in range(3):
        avgc.append(np.average(cc[idx]))

    avga.insert(2, 23.75) #2200
    avga.insert(1, 23.25) #1500
    avga.insert(1, 24.5) #1200
    avga.insert(1, 24.25) #900
    avgb.insert(2, 24)
    avgb.insert(1, 23.75)
    avgb.insert(1, 24)
    avgb.insert(1, 23.75)
    avgc.insert(2, 23.75)
    avgc.insert(1, 23.75)
    avgc.insert(1, 24)
    avgc.insert(1, 23.75)

    ax.plot(xx, avga, c='g', marker = "2", markersize=8, label="2 layers")
    ax.plot(xx, avgb, c='r', marker = "x", markersize=8, label="4 layers")
    ax.plot(xx, avgc, c='y', marker = "4", markersize=8, label="8 layers")
        
    plt.legend(loc = "upper left")
    plt.tight_layout()
    plt.show()
    plt.close()

def gpumem():
    #GPU Mem Utilization
    a1 = [6.24*324.8, 6.27*324.8, 5.71*324.8]
    a2 = [11.34*324.8, 11.34*324.8]
    a3 = [14.44*324.8, 14.27*324.8, 13.19*324.8]

    aa = []
    aa.append(a1)
    aa.append(a2)
    aa.append(a3)

    b1 = [6.84*324.8, 6.9*324.8, 6.29*324.8]
    b2 = [15.34*324.8, 15.31*324.8]
    b3 = [20.87*324.8, 20.83*324.8, 20.89*324.8, 19.97*324.8]

    bb = []
    bb.append(b1)
    bb.append(b2)
    bb.append(b3)

    c1 = [7.78*324.8, 7.8*324.8, 7.32*324.8]
    c2 = [21.17*324.8, 21.14*324.8, 20.46*324.8]
    c3 = [35.76*324.8, 35.84*324.8, 34.79*324.8]

    cc = []
    cc.append(c1)
    cc.append(c2)
    cc.append(c3)
    
    _, ax = plt.subplots()

    ax.set_xlabel('Hidden Units')
    ax.set_ylabel('GPU Memory Occupied/MB')
    ax.set_ylim(0,12000)
    xx = [650, 1800, 2500]
    for index in range(3):
        valx = []
        for _ in range(len(aa[index])):
            valx.append(xx[index])
        ax.scatter(valx, aa[index], c='g', marker='+', s=8)

    for index in range(3):
        valx = []
        for _ in range(len(bb[index])):
            valx.append(xx[index])
        ax.scatter(valx, bb[index], c='r', marker='+', s=8)
        
    for index in range(3):
        valx = []
        for _ in range(len(cc[index])):
            valx.append(xx[index])
        ax.scatter(valx, cc[index], c='y', marker='+', s=8)

    xx = [650, 900, 1200, 1500, 1800, 2200, 2500]
    #plt.xticks(xx,["650", "1800", "2500"])

    avga = []
    for idx in range(3):
        avga.append(np.average(aa[idx]))

    avgb = []
    for idx in range(3):
        avgb.append(np.average(bb[idx]))

    avgc = []
    for idx in range(3):
        avgc.append(np.average(cc[idx]))

    avga.insert(2, 12.47*324.8) #2200
    avga.insert(1, 9.6*324.8) #1500
    avga.insert(1, 8*324.8) #1200
    avga.insert(1, 7.33*324.8) #900
    avgb.insert(2, 19.32*324.8)
    avgb.insert(1, 11.6*324.8)
    avgb.insert(1, 9.56*324.8)
    avgb.insert(1, 7.97*324.8)
    avgc.insert(2, 28.07*324.8)
    avgc.insert(1, 17.54*324.8)
    avgc.insert(1, 13.19*324.8)
    avgc.insert(1, 9.84*324.8)

    ax.plot(xx, avga, c='g', marker = "2", markersize=8, label="2 layers")
    ax.plot(xx, avgb, c='r', marker = "x", markersize=8, label="4 layers")
    ax.plot(xx, avgc, c='y', marker = "4", markersize=8, label="8 layers")
        
    plt.legend(loc = "upper left")
    plt.tight_layout()
    plt.show()
    plt.close()

def gpumxutil():
    # GPU MAX Utilization
    a1 = [91.0, 92.0, 92.0]
    a2 = [98.0, 98.0]
    a3 = [98.0, 99.0, 98.0]

    aa = []
    aa.append(a1)
    aa.append(a2)
    aa.append(a3)

    b1 = [87.0, 86.0, 81.0]
    b2 = [99.0, 98.0]
    b3 = [99, 99, 99, 99]

    bb = []
    bb.append(b1)
    bb.append(b2)
    bb.append(b3)

    c1 = [86.0, 87.0, 87.0]
    c2 = [99, 99, 99]
    c3 = [100, 100, 100]

    cc = []
    cc.append(c1)
    cc.append(c2)
    cc.append(c3)

    _, ax = plt.subplots()

    ax.set_xlabel('Hidden Units')
    ax.set_ylabel('GPU Utilization/%')
    ax.set_ylim(0,100)
    xx = [650, 1800, 2500]
    for index in range(3):
        valx = []
        for _ in range(len(aa[index])):
            valx.append(xx[index])
        ax.scatter(valx, aa[index], c='g', marker='+', s=8)

    for index in range(3):
        valx = []
        for _ in range(len(bb[index])):
            valx.append(xx[index])
        ax.scatter(valx, bb[index], c='r', marker='+', s=8)
        
    for index in range(3):
        valx = []
        for _ in range(len(cc[index])):
            valx.append(xx[index])
        ax.scatter(valx, cc[index], c='y', marker='+', s=8)

    xx = [650, 900, 1200, 1500, 1800, 2200, 2500]
    #plt.xticks(xx,["650", "1800", "2500"])

    avga = []
    for idx in range(3):
        avga.append(np.average(aa[idx]))


    avgb = []
    for idx in range(3):
        avgb.append(np.average(bb[idx]))

    avgc = []
    for idx in range(3):
        avgc.append(np.average(cc[idx]))

    avga.insert(2, 98) #2200
    avga.insert(1, 96) #1500
    avga.insert(1, 95) #1200
    avga.insert(1, 93)
    avgb.insert(2, 99)
    avgb.insert(1, 98)
    avgb.insert(1, 95)
    avgb.insert(1, 92)
    avgc.insert(2, 99)
    avgc.insert(1, 98)
    avgc.insert(1, 96)
    avgc.insert(1, 92)

    ax.plot(xx, avga, c='g', marker = "2", markersize=8, label="2 layers")
    ax.plot(xx, avgb, c='r', marker = "x", markersize=8, label="4 layers")
    ax.plot(xx, avgc, c='y', marker = "4", markersize=8, label="8 layers")
        
    plt.legend(loc = "lower left")
    plt.tight_layout()
    plt.show()
    plt.close()

def gpupcie():
    #GPU rx
    # a1 = [,,]
    # a2 = [,]
    # a3 = [, ]
    
    # a1x = [31.37,35.37,32.98]
    # a2x = [38.84,40.57]
    # a3x = [41.73,41.46]
    
    
    #GPU tx
    a1 = [2.52+19.19, 2.39+18.25, 2.69+20.71]
    a2 = [2.97+22.86, 3.04+23.53]
    a3 = [3.17+24.67, 3.18+24.6]
    
    # a1x = [4.02,4.53,4.23]
    # a2x = [4.95,5.17]
    # a3x = [5.31,5.25]

    aa = []
    aa.append(a1)
    aa.append(a2)
    aa.append(a3)
    
    # aax = []
    # aax.append(a1x)
    # aax.append(a2x)
    # aax.append(a3x)

    # b1 = [, ]
    # b2 = [, , ]
    # b3 = [, ]
    
    # b1x = [37.18,37.47]
    # b2x = [44.68,44.32, 44.9]
    # b3x = [47.18, 47.63]
    
    b1 = [2.86+22.09, 2.83+21.85]
    b2 = [3.37+26.02, 3.41+26.51, 3.45+26.53]
    b3 = [3.59+27.79, 3.57+27.67]
    
    # b1x = [4.73,4.78]
    # b2x = [5.67, 5.62, 5.73]
    # b3x = [5.98, 6.04]

    bb = []
    bb.append(b1)
    bb.append(b2)
    bb.append(b3)
    
    # bbx = []
    # bbx.append(b1x)
    # bbx.append(b2x)
    # bbx.append(b3x)

    # c1 = [, ,  ]
    # c2 = [, , ,  ]
    # c3 = []
    
    # c1x = [37.54, 38.8, 38.22]
    # c2x = [46.18, 45.05, 46.33, 46.08]
    # c3x = [48.03, 47.64]
    
    c1 = [2.91+22.48, 2.99+23.21, 2.78+21.45]
    c2 = [3.48+26.92, 3.26+26.82, 3.46+26.79, 3.39+26.31 ]
    c3 = [3.55+27.66, 3.62+28.03]
    
    # c1x = [4.87, 4.94, 4.86]
    # c2x = [5.84, 5.72, 5.87, 5.84]
    # c3x = [6.07, 6.05]

    cc = []
    cc.append(c1)
    cc.append(c2)
    cc.append(c3)
    
    # ccx = []
    # ccx.append(c1x)
    # ccx.append(c2x)
    # ccx.append(c3x)
    
    _, ax = plt.subplots()

    ax.set_xlabel('Hidden Units')
    ax.set_ylabel('PCIE BandWidth Occupied/%')
    ax.set_ylim(0,100)
    xx = [650, 1800, 2500]
    for index in range(3):
        valx = []
        for _ in range(len(aa[index])):
            valx.append(xx[index])
        ax.scatter(valx, aa[index], c='g', marker='+', s=8)

    for index in range(3):
        valx = []
        for _ in range(len(bb[index])):
            valx.append(xx[index])
        ax.scatter(valx, bb[index], c='r', marker='+', s=8)
        
    for index in range(3):
        valx = []
        for _ in range(len(cc[index])):
            valx.append(xx[index])
        ax.scatter(valx, cc[index], c='y', marker='+', s=8)

    xx = [650, 900, 1200, 1500, 1800, 2200, 2500]
    #plt.xticks(xx,["650", "1800", "2500"])

    avga = []
    for idx in range(3):
        avga.append(np.average(aa[idx]))
    # avgax = []
    # for idx in range(3):
    #     avgax.append(np.average(aax[idx]))
    


    avgb = []
    for idx in range(3):
        avgb.append(np.average(bb[idx]))
    #avgbx = []
    # for idx in range(3):
    #     avgbx.append(np.average(bbx[idx]))

    avgc = []
    for idx in range(3):
        avgc.append(np.average(cc[idx]))
    # avgcx = []
    # for idx in range(3):
    #     avgcx.append(np.average(ccx[idx]))

    # avga.insert(2, ) #2200
    # avga.insert(1, ) #1500
    # avga.insert(1, ) #1200
    # avga.insert(1, ) #900
    
    # avgax.insert(2, 39.41) #2200
    # avgax.insert(1, 36.46) #1500
    # avgax.insert(1, 35.42) #1200
    # avgax.insert(1, 35.12) #900
    
    avga.insert(2, 23.15+3) #2200
    avga.insert(1, 21.58+2.81) #1500
    avga.insert(1, 19.6+2.56) #1200
    avga.insert(1, 20.65+2.68) #900
    
    # avgax.insert(2, 5) #2200
    # avgax.insert(1, 4.73) #1500
    # avgax.insert(1, 4.53) #1200
    # avgax.insert(1, 4.49) #900
    
    
    
    # avgb.insert(2, )
    # avgb.insert(1, )
    # avgb.insert(1, )
    # avgb.insert(1, )
    # avgbx.insert(2, 45.66)
    # avgbx.insert(1, 45.77)
    # avgbx.insert(1, 43.94)
    # avgbx.insert(1, 44.55)
    
    avgb.insert(2, 26.79+3.44)
    avgb.insert(1, 26.87+3.46)
    avgb.insert(1, 25.97+3.36)
    avgb.insert(1, 26.43+3.43)
    # avgbx.insert(2, 5.76)
    # avgbx.insert(1, 5.81)
    # avgbx.insert(1, 5.59)
    # avgbx.insert(1, 5.7)
    
    # avgc.insert(2, )
    # avgc.insert(1, )
    # avgc.insert(1, )
    # avgc.insert(1, )
    # avgcx.insert(2, 47.15)
    # avgcx.insert(1, 46.43)
    # avgcx.insert(1, 45.26)
    # avgcx.insert(1, 44.82)
    
    avgc.insert(2, 27.4+3.52)
    avgc.insert(1, 26.13+3.37)
    avgc.insert(1, 26.53+3.43)
    avgc.insert(1, 26.59+3.45)
    # avgcx.insert(2, 5.94)
    # avgcx.insert(1, 5.88)
    # avgcx.insert(1, 5.75)
    # avgcx.insert(1, 5.76)

    ax.plot(xx, avga, c='g', marker = "2", markersize=8, label="2 layers")
    ax.plot(xx, avgb, c='r', marker = "x", markersize=8, label="4 layers")
    ax.plot(xx, avgc, c='y', marker = "4", markersize=8, label="8 layers")
        
    plt.legend(loc = "upper left")
    plt.tight_layout()
    plt.show()
    plt.close()
    
    # _, ax = plt.subplots()

    # ax.set_ylabel('GPU PCIE TX MAX/%')
    # ax.plot(xx, avgax, c='g', marker = "2", markersize=8, label="2 layers")
    # ax.plot(xx, avgbx, c='r', marker = "x", markersize=8, label="4 layers")
    # ax.plot(xx, avgcx, c='y', marker = "4", markersize=8, label="8 layers")
        
    # plt.legend(loc = "upper left")
    # plt.tight_layout()
    # plt.show()
    # plt.close()

def resnetpcie():
    bnkrx = [0.75, 0.73, 0.72, 0.74, 0.71, 0.76, 0.76]
    bnktx = [0.12, 0.12, 0.12, 0.13, 0.13, 0.13, 0.14]
    basrx = [0.65, 0.7, 0.7, 0.59, 0.66, 0.65, 0.81, 0.67]
    bastx = [0.1, 0.11, 0.12, 0.1, 0.12, 0.12, 0.14, 0.12]
    bnkx = [20, 29, 47, 56, 74, 92, 110]
    basx = [20, 32, 44, 56, 68, 80, 92, 110]
    
    for i in range(len(bnkrx)):
        bnkrx[i] = bnkrx[i]+bnktx[i]
    for i in range(len(basrx)):
        basrx[i] = bastx[i]+basrx[i]
    
    plt.xlabel("Convolutional layer")
    plt.ylabel('PCIE BandWidth Occupied/%')
    plt.ylim(0,10)
    plt.plot(bnkx, bnkrx, linestyle=':', label="bottleneck")
    
    plt.plot(basx, basrx, linestyle='--', label="basicblock")
    
    
    #x = ["resnet20", "resnet29", "resnet32", "resnet44", "resnet47", "resnet56", "resnet68", "resnet74", "resnet80", "resnet92", "resnet110"]
    plt.xticks([20,29,32,44,47,56,68,74,80,92,110], [20,29,32,44,47,56,68,74,80,92,110])
    plt.legend()
    plt.tight_layout()
    plt.show()
    plt.close()
# gpuutili()
# cpuutil()
gpumem()
# gpumxutil()
# gpupcie()
# resnetpcie()

# #bnk
# [20, 29, 47, 56, 74, 92, 110]
# #[enxy402, renxy177, renxy61, renxy828, renxy939, renxy102, renxy872]
# [0.75, 0.73, 0.72, 0.74, 0.71, 0.76, 0.76]
# [0.12, 0.12, 0.12, 0.13, 0.13, 0.13, 0.14]
# #bas
# [20,32, 44, 56, 68, 80, 92, 110]
# #[renxy887, renxy836, renxy747, renxy823, renxy672, renxy169, renxy137, renxy344]
# [0.65, 0.7, 0.7, 0.59, 0.66, 0.65, 0.81, 0.67]
# [0.1, 0.11, 0.12, 0.1, 0.12, 0.12, 0.14, 0.12]
