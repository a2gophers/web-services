// +build OMIT
package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/a2gophers/web-services/dbutils"
	"github.com/ttacon/chalk"
	"github.com/ttacon/go-tigertonic"
	"github.com/ttacon/gorp"
)

var (
	webPort = flag.String("p", "18090", "port to run the web server on")
	dbMap   *gorp.DbMap
)

func main() {
	flag.Parse()

	// set up db conn
	db, err := dbutils.DBConnFromFlags()
	if err != nil {
		errLog("failed to open db conn: %v", err)
		return
	}

	// START OMIT
	dbMap = &gorp.DbMap{
		Db: db,
		Dialect: gorp.MySQLDialect{
			Engine:   "InnoDB",
			Encoding: "UTF8",
		},
	}

	dbMap.AddTableWithName(Herro{}, "herro").SetKeys(false, "ID")

	// STOP OMIT

	// set up router
	mux := tigertonic.NewTrieServeMux()

	mux.HandleFunc("POST", "/herro/{id}", saveHerro)
	mux.HandleFunc("GET", "/herro/{id}", getHerro)

	fmt.Printf(chalk.Cyan.Color("[server]")+" listening on :%s...\n", *webPort)
	http.ListenAndServe(":"+*webPort, mux)
}

func getHerro(w http.ResponseWriter, r *http.Request) {
	// START2 OMIT
	tx, err := dbMap.Begin()
	if err != nil {
		errLog("failed to begin transaction: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var (
		h  Herro
		d  interface{}
		id = r.URL.Query().Get("id")
	)
	if d, err = tx.Get(&h, id); err != nil {
		errLog("failed to retrieve herro: %v", err)
		tx.Rollback()
		w.WriteHeader(http.StatusNotFound)
		return
	}

	tx.Commit()
	// STOP2 OMIT
	h = *d.(*Herro)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`
<html>
  <head><title>Herro There!</title></head>
  <body>
    <div style="">
      <h1>This is a web page for ID: %s (val: %s)</h1>
      <hr/>
      <p>
        So much cool, right?
      </p>
    </div>
  </body>
</html>
`,
		id, h.Val)))
}

func saveHerro(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	val := r.URL.Query().Get("val")
	if len(val) == 0 {
		val = "default"
	}

	// START1 OMIT
	tx, err := dbMap.Begin()
	if err != nil {
		errLog("failed to begin transaction, err: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var herro = Herro{
		ID:  id,
		Val: val,
	}

	if err = tx.Insert(&herro); err != nil {
		errLog("failed to save herro: %v", err)
		w.WriteHeader(http.StatusNotFound)
		tx.Rollback() // yes we're ignoring the error
		return
	}

	tx.Commit() // yes, still ignoring the error
	// STOP1 OMIT
	http.Redirect(w, r, fmt.Sprintf("/herro/%s", id), http.StatusFound)
}

func errLog(fmtString string, err error) {
	fmt.Println(
		chalk.Bold.NewStyle().WithForeground(chalk.Red).Style(
			fmt.Sprintf(fmtString, err),
		),
	)
}

type Herro struct {
	ID  string
	Val string
}
