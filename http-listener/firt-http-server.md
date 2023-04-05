# First HTTP Server



Creating an HTTP server using Go is pretty straightforward. Golang has built in support fpr creating HTTP servers using tht **`net/http`** package.



**`net/http`** package provides functions for requests and response. 



#### Steps to create the HTTP server :

1. Create a HTTP handler function which processes the incoming HTTP requests and generates responses.         
   
   - This takes an `http.ResponseWriter` and `http.Request` as arguments          

2. Register the HTTP handler function by using `http.HandleFunc` which would associate the handler function to a specified URL path. 

3. Start the HTTP server using `http.ListenAndServe()`
   
   - This listens for an incoming HTTP request on the specified port and dispatches that request to the appropriate handler function.



Code File : 

```go
package main

// defining main package is necessary because that tells the compiler that this should
// be compiled as a standalone executable program.

// without it the compiler won't make an executable binary

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// entry point of the executable !

	// problem statement: Create a web server!
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello World!")

		d, err := ioutil.ReadAll(r.Body)

		if err != nil {
			//w.WriteHeader(http.StatusBadRequest)
			//w.Write([]byte("Oops"))
			http.Error(w, "oops", http.StatusBadRequest)
			return
		}
		log.Printf("Data :%s", d)

		fmt.Fprintf(w, "Hello %s", d)

	})
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World!")
	})
	http.ListenAndServe(":9090", nil)

}

```



```
curl  -d 'Manas!' localhost:9090/hello
```

response : 

```
Hello Manas!
```

> hitting a HTTP request to localhost:9090/goodbye would print a log message "Goodbye World!"
