from mpl_toolkits.mplot3d import Axes3D
import numpy as np
from matplotlib import pyplot as plt

fig = plt.figure()
ax = Axes3D(fig)
x = np.arange(0, 10 * np.pi, 0.1)
y = np.arange(0, 10 * np.pi, 0.1)
X, Y = np.meshgrid(x, y)
# Z = np.sin(X) * np.cos(Y)
Z = Y - X * 2
# Z=-X2-Y2
# Z=2X+2Y
# Z=4*np.sin(X)+Y**2
plt.xlabel('x')
plt.ylabel('y')
ax.plot_surface(X, Y, Z, rstride=1, cstride=1, cmap='rainbow')
ax.plot_surface(X, Y, 0, rstride=1, cstride=1, cmap='rainbow')
ax.plot_surface(-4, 2, 10, rstride=1, cstride=1, cmap='rainbow')
plt.show()
