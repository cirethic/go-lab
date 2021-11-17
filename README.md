# go-math
fun math - what could be done differently in Go

**1. cantor pair function**

![image](https://user-images.githubusercontent.com/69271523/142099519-3f138495-e962-40e3-b42d-7877594051b2.png)


cantor  pair function  reference

    https://en.wikipedia.org/wiki/Cantor_function
    
    https://en.wikipedia.org/wiki/Pairing_function
    
    https://gist.github.com/hannesl/8031402 
    
    
    
    go run main.go
    
    **func cantor_pair_calculate()**
    33, 1  =>  596
    6, 8  =>  113
    596, 113  =>  251808
    
    **func cantor_pair_reverse()**
    251808 =>  596, 113
    596 => 33, 1
    113 => 6, 8
    
