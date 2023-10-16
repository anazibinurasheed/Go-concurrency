# Go-concurrency patterns and practices

This repository contains Go-concurrency patterns and practices.
It contains different implementations using channels, switch, go-routines .


## Why concurrency supported in go ?
Look around you. What do you see?

Do you see a single-stepping world doing one thing at a time?

Or do you see a complex world of interacting, independently behaving pieces?

That's why. Sequential processing on its own does not model the world's behavior.

## What is concurrency?


Concurrency is the composition of independently executing computations.

Concurrency is a way to structure software, particularly as a way to write clean code that interacts well with the real world.

## Concurrency is not parallelism

Concurrency is not parallelism, although it enables parallelism.

If you have only one processor, your program can still be concurrent but it cannot be parallel.

On the other hand, a well-written concurrent program might run efficiently in parallel on a multiprocessor. That property could be important...

## What is a goroutine?
 It's an independently executing function, launched by a go statement.

It has its own call stack, which grows and shrinks as required.

It's very cheap. It's practical to have thousands, even hundreds of thousands of goroutines.

It's not a thread.

There might be only one thread in a program with thousands of goroutines.

Instead, goroutines are multiplexed dynamically onto threads as needed to keep all the goroutines running.

But if you think of it as a very cheap thread, you won't be far off.


## Channels
A channel in Go provides a connection between two goroutines, allowing them to communicate.

     Receiving from a channel.
     The "arrow" indicates the direction of data flow.
    value = <-c
     c <- "msg"

- A sender and receiver must both be ready to play their part in the communication. Otherwise we wait until they are.
Thus channels both communicate and synchronize.     

- buffer: Buffering removes synchronization.


### The Go approach
Don't communicate by sharing memory, share memory by communicating.