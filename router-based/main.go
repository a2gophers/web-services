// +build OMIT
package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/ttacon/chalk"
	"github.com/ttacon/go-tigertonic"
)

// START OMIT
var webPort = flag.String("p", "18090", "port to run the web server on")

func main() {
	flag.Parse()

	mux := tigertonic.NewTrieServeMux()

	mux.HandleFunc("GET", "/herro/{id}", idHerro)

	fmt.Printf(chalk.Cyan.Color("[server]")+" listening on :%s...\n", *webPort)
	http.ListenAndServe(":"+*webPort, mux)
}

// STOP OMIT
func idHerro(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`
<html>
  <head><title>Herro There!</title></head>
  <body>
    <div style="">
      <h1>This is a web page for ID: %s</h1>
      <hr/>
      <p>
        So much cool, right?
      </p>
    </div>
  </body>
</html>
`,
		r.URL.Query().Get("id"))))
}
