# Kruskal Count simulator

Monte-Carlo simulator of the [Kruskal Count](https://arxiv.org/pdf/math/0110143.pd) card trick.

How to run:

```
go build kruskal.go
./kruskal -n 1000000000 # This takes about a minute on my machine.
P(success) = 854659035/1000000000 (0.854659035)
```
