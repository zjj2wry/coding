"""
将彩色的图片转换为黑白图片

单通道图: 俗称灰度图，每个像素点只能有有一个值表示颜色，它的像素值在0到255之间，0是黑色，255是白色，中间值是一些不同等级的灰色。
（也有3通道的灰度图，3通道灰度图只有一个通道有值，其他两个通道的值都是零）。
三通道图: 每个像素点都有3个值表示 ，所以就是3通道。也有4通道的图。例如RGB图片即为三通道图片，RGB色彩模式是工业界的一种颜色标准，
是通过对红(R)、绿(G)、蓝(B)三个颜色通道的变化以及它们相互之间的叠加来得到各式各样的颜色的，RGB即是代表红、绿、蓝三个通道的颜色，
这个标准几乎包括了人类视力所能感知的所有颜色，是目前运用最广的颜色系统之一。总之，每一个点由三个值表示。
"""

from scipy import ndimage
from matplotlib import pyplot as plt
import cv2

rgb = cv2.imread("rgb.jpg")
gray_image = cv2.cvtColor(rgb, cv2.COLOR_BGR2GRAY)

print(rgb)
print(rgb.shape) # (758, 600, 3)
print(gray_image)
print(gray_image.shape) # (758, 600)

plt.imshow(gray_image,cmap='Greys')
plt.show()

