#!/usr/bin/env python

import sys
import os.path
import argparse
from pylab import *

def parse_arguments():
    parser = argparse.ArgumentParser(description='benchmark result plotter')
    parser.add_argument('result_file', )
    return parser.parse_args()

opts = parse_arguments()

bname = os.path.splitext(os.path.basename(opts.result_file))[0]
out_name = 'plot-' + bname + '.png'
data = np.loadtxt(opts.result_file, delimiter=',', skiprows=1)

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


savefig(out_name)
