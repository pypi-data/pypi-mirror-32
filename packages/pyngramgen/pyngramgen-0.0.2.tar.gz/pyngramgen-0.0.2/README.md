Overview
========
This lib generate a ngram iterator

Usage
=====
```
import pyngramgen
if __name__ == '__main__':
    if len(sys.argv) != 2:
        print "python %s n" % (sys.argv[0])
        quit()
    while True:
        iput = raw_input()
        obj = pyngramgen.PyNgramGen(int(sys.argv[1]), 'abcdefghijklmnopqrstuvwxyz')
        for i in obj.generator(iput):
            print i
```

