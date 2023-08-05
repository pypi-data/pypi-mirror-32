#! /usr/bin/python
import rowingdata
from sys import argv

def main():
    readFile=argv[1]

    try:
	rowerFile=argv[2]
    except IndexError:
	rowerFile="defaultrower.txt"

    rower=rowingdata.getrower(rowerFile)
    
    row=rowingdata.rowingdata(readFile,rowtype="Indoor Rower",
				    rower=rower)
    
    print(row.allstats())

    row.plottime_erg()




    print("done "+readFile)
