# coding=utf-8
""" sigmoid 主要是将值映射到(0,1) 之间，作用于一个数值，可以用于二分类
"""

import matplotlib.pyplot as plt
import numpy as np

def sigmoid(x):
    return 1/(1+np.exp(-x))

def sigmoid_prime(x):
    """Derivative of the sigmoid function."""
    return sigmoid(x)*(1-sigmoid(x))

x = np.arange(-10., 10., 0.2)
sig = [sigmoid(i) for i in x]
plt.plot(x,sig)
plt.show()