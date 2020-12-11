import numpy as np
import matplotlib.pyplot as plt
import matplotlib.cm as cm

a = np.loadtxt(open("input_py.txt", "rb"), delimiter=", ")
plt.imshow(a)
plt.show()
