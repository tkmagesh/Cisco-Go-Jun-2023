# Concurrency #
## Channels ##
- Declaration
    > var <chan_var> chan <datatype>
    > ex: var ch chan int
- Instantiation
    > ch = make(chan int)
- Operations (<- operator)
    - Send Operation
        > ch <- 100
    - Receive Operation
        > data := <-ch

### Channel Behaviors ###
    - A RECEIVE operation is ALWAYS a blocking operation
    - A SEND operation is BLOCKED until a RECEIVE operation is initiated

