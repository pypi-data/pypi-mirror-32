import argparse
import numpy as np 
import matplotlib.pyplot as plt 
from general.md import (make_bond_matrix, rdf) 

def main(fn, atoms_i, atoms_j, dr, rmax, start, stride):
    x, y = rdf(fn, atoms_i, atoms_j, dr, rmax, start, stride)  
    plt.xlabel('radius R (Angstrom)') 
    plt.ylabel('Pair distributon function (arb. un.)') 
    plt.plot(x, y)
    plt.show()

def read_cmd_line(parser):
    """
    Parse Command line options.
    """
    args = parser.parse_args()

    attributes = ['file', 'atom1', 'atom2', 'bin', 'rmax', 'start', 'stride']

    return [getattr(args, p) for p in attributes]

if __name__ == "__main__":
    msg = "rdf -file <path/to/filename> -atom1 <atom typer> -atom2 <atom type> -bin <size of the histogram bin> -rmax <cutoff> -start <starting frame> -stride <skipped frames >"

    parser = argparse.ArgumentParser(description=msg)
    parser.add_argument(
        '-file', required=True, help='name of the xyz trajectory file')
    parser.add_argument(
        '-atom1', required=True, help='first atom type to compute rdf')
    parser.add_argument(
        '-atom2', required=True, help='second atom type to compute rdf')
    parser.add_argument(
        '-bin', required=False, type=float, default=0.05, help='size of the bin')
    parser.add_argument(
        '-rmax', required=False, type=float, default=12.0, help='cutoff in Angstrom for computing the rdf')
    parser.add_argument(
        '-start', required=False, type=int, default=1, help='starting frame for computing rdf')
    parser.add_argument(
        '-stride', required=False, type=int, default=1, help='skip n frames')
    main(*read_cmd_line(parser))





