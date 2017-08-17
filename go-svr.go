package main

import (
	"fmt"
	"github.com/alecthomas/kingpin"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func cwd() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return cwd
}

var (
	app = kingpin.New("go-svr", `Super simple Golang HTTP server.
It hosts --js files under /js, --css files under /css, and --html files under /
It registers a /random route that returns data in JSON form:\
  {
    "int": [0,64-bit-max-int],
    "float": 0-1.0
  }
It registers a /time route that returns a JSON structure like the following:
  {
    "now": current time in RFC3339 / "2006-01-02T15:04:05Z07:00" format
    "runtime": duration server has been running
  }
Otherwise, it serves out files that exist in whatever --html points to.
`)
	js   = app.Flag("js", "hosted javascript folder to map to /js.  Default expects a folder named 'js' in the current working directory").Short('j').Default("js").String()
	css  = app.Flag("css", "hosted css folder to map to /css.  Default expects a folder named 'css' in the current working directory").Short('c').Default("css").String()
	html = app.Flag("html", "hosted set of HTML files to serve out of /. Default is the current working directory.").Short('h').Default(cwd()).String()
	port = app.Flag("port", "TCP port to host the HTTP traffic on.").Short('p').Default(":8080").String()
)

var startTime = time.Now()

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	router := mux.NewRouter()

	router.HandleFunc("/random", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"int":%d,"float":"%v"}`, rand.Int(), rand.Float64())
	})
	router.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"now":"%v","runtime":"%v"}`, time.Now().Format(time.RFC3339), time.Since(startTime))
	})
	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir(*js))))
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir(*css))))
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(*html)))
	// router.Handle("/", http.FileServer(http.Dir(*html)))
	fmt.Println("Starting server on", *port)
	http.ListenAndServe(*port, router)
}
