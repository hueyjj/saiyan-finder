# saiyan-finder
Finds a file at super saiyan speed, or not.

# Install
```bash
$ go get github.com/hueyjj/saiyan-finder
```

# Build
```bash
$ go build
```

# Running
```bash
$ ./saiyan-finder wheremyfileat
```

# Benchmarks
This should be obvious, but saiyan-finder is a super, simple file finder and probably is never going to outmatch the better tools out there. It literally finds all the file in a directory and does a substring comparison to find a file match. In the name of go's concurrency, saiyan-finder is supposed to be fast if this was done right.

Other tools like find, ripgrep, grep, etc. have pattern matching, optimizations in some cases, and more.

We'll use the FreeBSD kernel source code as a benchmark, indexing 73666 files and directories.
```bash
$ fetch ftp://ftp.freebsd.org/pub/FreeBSD/releases/amd64/12.0-RELEASE/src.txz
or
$ wget ftp://ftp.freebsd.org/pub/FreeBSD/releases/amd64/12.0-RELEASE/src.txz
$ find . -type f | wc -l # Count how many files there are
```
```bash
$ cd ~/Downloads
$ tar xvf src.txz
$ cd usr/src
```
Run saiyan-finder
```bash
$ time saiyan-finder.exe pipe
```
```
real	0m0.276s
user	0m0.621s
sys	0m0.517s
```
Run find
```bash
$ time find . -name "*pipe*"
```
```
real	0m0.098s
user	0m0.034s
sys	0m0.064s
```
Run ripgrep
```bash
$ time rg --files -g '*pipe*'
```
```
real	0m0.058s
user	0m0.191s
sys	0m0.139s
```
Clearly, we can see that saiyan-finder is **hot garbage** even when improperly using ripgrep and find.