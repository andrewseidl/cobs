package backend

import (
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("hello"))
}

func BuildTarballHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte(mux.Vars(r)["imageid"]))
}

func BuildDockerfileHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte(mux.Vars(r)["imageid"]))
}

func BuildStatusHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte(mux.Vars(r)["imageid"]))
}

func Run() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)

	r.HandleFunc("/api/v1/build/{imageid}/tarball", BuildTarballHandler)
	r.HandleFunc("/api/v1/build/{imageid}/dockerfile", BuildDockerfileHandler)
	r.HandleFunc("/api/v1/build/{imageid}", BuildStatusHandler)

	n := negroni.Classic()
	n.UseHandler(r)

	// shouldn't this be automatic?
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	port = ":" + port
	n.Run(port)
}
