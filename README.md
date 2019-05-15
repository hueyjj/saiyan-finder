# saiyan-finder
Finds a file at super saiyan speed.

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
./saiyan-finder wheremyfileat
```

# Benchmarks
This should be obvious, but saiyan-finder is a super, simple file finder and probably is never going to outmatch the better tools out there. It literally finds all the file in a directory and does a substring comparison to find a file match. In the name of go's concurrency, saiyan-finder is supposed to be fast if this was done right.

Other tools like find, ripgrep, grep, etc. have pattern matching, optimizations in some cases, and more.

We'll use the FreeBSD kernel source code as a benchmark.
```bash
$ fetch ftp://ftp.freebsd.org/pub/FreeBSD/releases/amd64/12.0-RELEASE/src.txz
or
$ wget ftp://ftp.freebsd.org/pub/FreeBSD/releases/amd64/12.0-RELEASE/src.txz
```
```bash
cd ~/Downloads
tar xvf src.txz
cd usr/src
```
Run saiyan-finder
```bash
time saiyan-finder.exe pipe
```
Run find
```bash
time find . -name "*pipe*"
```
Run ripgrep
```bash
time rg --files -g '*pipe*'
```
