# ssadump
Dumps the SSA representation of a function to stdout.

## Usage
```
ssadump filename function
```

## Example
```
$ cat > hello.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}

$ ssadump hello.go main
block 0:
*ssa.Alloc              t0  :=  new [1]interface{} (varargs)
*ssa.IndexAddr          t1  :=  &t0[0:int]
*ssa.MakeInterface      t2  :=  make interface{} <- string ("Hello, world!":string)
*ssa.Store             N/A      *t1 = t2
*ssa.Slice              t3  :=  slice t0[:]
*ssa.Call               t4  :=  fmt.Println(t3...)
*ssa.Return            N/A      return

```
