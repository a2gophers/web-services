// +build OMIT
package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/ttacon/chalk"
)

// START OMIT
var webPort = flag.String("p", "18090", "port to run the web server on")

func main() {
	flag.Parse()

	http.HandleFunc("/herro", herro)

	fmt.Printf(chalk.Cyan.Color("[server]")+" listening on :%s...\n", *webPort)
	http.ListenAndServe(":"+*webPort, nil)
}

func herro(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(herroPage)
}

// STOP OMIT

var herroPage = []byte(`
<html>
  <head><title>Herro There!</title></head>
  <body>
    <div style="">
      <h1>This is a web page</h1>
      <hr/>
      <p>
        So much cool, right?
      </p>
    </div>
  </body>
</html>
`)
