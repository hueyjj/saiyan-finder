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
This should be obviously, but saiyan-finder is a super, simple file finder. It literally finds all the file in a directory and does a substring comparison to find a file match. In the name of go's concurrency, saiyan-finder is supposed to be fast if this was done right.


Other tools like find, ripgrep, grep, etc. have much pattern matching, optimizations in some scenarios, and more.