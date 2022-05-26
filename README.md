# Generate an EVM Address of your choice ! 

### replace `suffix` and `prefix` with your desired beginning and ending! 
Node: the Longer the longer it takes to find an address! So keep it short! 
For no desired `suffix` or `prefix` just give it an empty string `""`

then run:

``` 
go run main.go
```

for quicker results you can uncomment the `go gen()` to get more goroutines. (threads)

The counter only counts the loops not the generated addresses. 

So if you have 3 goroutines and the counter is at 1 million, you generated about 3 million addresses.

To stop the script, just use ctrl + C for mac and (windows+X for windows)

Setup Go: 
https://www.youtube.com/watch?v=dgIh-VYcWYw