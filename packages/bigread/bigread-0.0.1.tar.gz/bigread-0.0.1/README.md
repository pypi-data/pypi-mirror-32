# BigRead

In text processing, it is sometimes necessary to read files that are too large to fit into a machine's main memory. In those situations, one must read only a few lines of the file into RAM, then read the next few lines, and so on, this way ensuring that only a small number of lines from the input file are in main memory at any given time. This package implements a simple file reader for exactly these situations, to allow machines to read files that are absolutely huge.

## Installation

To install bigread, just run:

```
pip install bigread
```

## Usage

Suppose you have a large file to read into RAM:

```python
with open('large.txt', 'w') as out:
  for i in range(10**10):
    out.write('this is line ' + str(i) + '\n')
```

To read this file with bigread, import the module and create a stream, then iterate over the stream to read the file lines one-by-one:

```python
from bigread import Reader

stream = Reader(file='large.txt', block_size=10)
for i in stream:
  print(i)
```

This will iterate over each line in the input file without loading all of the file into RAM concurrently.

If it's more convenient, you can also read lines from the file one-by-one "on demand" by calling the `.next()` method to fetch the next line from the file:

```
stream = Reader(file='large.txt', block_size=100)
print( stream.next() ) # prints the first line in the text
print( stream.next() ) # prints the second line in the text...
```
