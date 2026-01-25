# Simple Lang

Uses antlr to create ast for a really simple programming language.
## build.bash
Python is needed to install antlr (I used python 3.11)
```
pip install antlr4-tools
```

build.bash generates the parser based on the .g4 file. Antlr needs Java.

## Transbiling

Generate go code from the input simple lang.

```
go run main.go
```