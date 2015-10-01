#!/usr/bin/env python

from pylab import *

data = np.loadtxt('benchresults-cache.csv', delimiter=',', skiprows=1)

dpi = 96
w, h = 640, 480
fig = figure(figsize=(w/dpi, h/dpi), dpi=dpi)
plt.plot(data[:,0], data[:,1], 'bo-', label='Map')
plt.plot(data[:,0], data[:,2], 'ro-', label='Slice')
plt.xscale('log')
plt.yscale('log')
plt.ylabel('Cache misses (#)')
plt.xlabel('Elements count')
plt.grid(True)
legend(loc='upper left')


savefig('plot-cache.png')
