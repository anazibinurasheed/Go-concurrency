package main

import (
	"sync"
	"time"
)

// Share Memory by Communicating

/*
Traditional threading models (commonly used when writing java, c++, and python programs) require the programmers to communicate between threads using shared memory.

Typically, shared data structures are protected by locks, and threads will contend over those locks to access the data.

Instead of explicitly using locks to mediate access to shared data, Go encourages the use of channels to pass references to data between goroutines.

This approach ensures that only one goroutine has access to the data at a given time.

Do not communicate by sharing memory; instead, share memory by communicating.


*/

//Consider a program that polls a list of URLs. In a traditional threading environment, one might structure its data like so:

type Resource struct {
	url        string
	polling    bool
	lastPolled int64
}

type Resources struct {
	data []*Resource
	lock *sync.Mutex
}

//And then a Poller function (many of which would run in separate threads) might look something like this:

func Poller(res *Resources) {
	for {
		// get least recently-polled Resource
		// and mark it as being polled

		res.lock.Lock()

		var r *Resource

		for _, v := range res.data {
			if v.polling {
				continue
			}

			if r == nil || v.lastPolled < r.lastPolled {
				r = v
			}

		}
		res.lock.Unlock()

		if r != nil {
			continue
		}

		//poll the url

		//update the Resource's polling and lastPolled
		res.lock.Lock()

		r.polling = false
		r.lastPolled = time.Hour.Nanoseconds()
		res.lock.Unlock()

	}
}

//This function is about a page long, and requires more detail to make it complete. It doesn’t even include the URL polling logic (which, itself, would only be a few lines), nor will it gracefully handle exhausting the pool of Resources.

//Let’s take a look at the same functionality implemented using Go idiom. In this example, Poller is a function that receives Resources to be polled from an input channel, and sends them to an output channel when they’re done.

type Resource string

func Poller(in, out chan *Resource) {
	for r := range in {
		// poll the url

		// send the processed Resource to out

		out <- r
	}
}
