
# Performance Testing

##  Background 


This folder contains a performance impact of changing one line of Golang code. During our kata, we noticed rather than assigning a value to an array item within an if/else logic, we could assign a default value at first, and overwrite with a single if statement. 

<table>
<tr><th>With Else</th><th>Without Else</th></tr>
<tr><td> 
<pre lang=golang>
func ... {
    if |condition| {
        array[x]=z  
    } else {
        array[x]=y
    }
}
</pre>
</td>
<td><pre lang=golang>
func ... {
    array[x] = y
    if |condition| {
        array[x]=z  
    }
}
</td>
</tr>
</table>

## Actions Taken

Two different packages created under two folders with a `main.go` at root calling these functions. 
Then, following `Dave Cheney's` <a href="https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go"> blogpost </a> benchmarking tests added. The code to execute these tests with the results: 
``` golang
go test -run=XXX -bench=. 
```
Results from `WithElse`:  
```
24882            107140 ns/op
PASS
ok      github.com/go-kata-starter/trialBigOClient/trialBigOWithElse    3.975s
```
```
10000            104536 ns/op
PASS
ok      github.com/go-kata-starter/trialBigOClient/trialBigOWithElse    1.113s
```
And results from `WithoutElse`: 
```
10000            161712 ns/op
PASS
ok      github.com/go-kata-starter/trialBigOClient/trialBigOWithoutElse 1.799s
```
```
10000            144646 ns/op
PASS
ok      github.com/go-kata-starter/trialBigOClient/trialBigOWithoutElse 1.620s
```
