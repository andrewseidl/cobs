package backend

import (
	"log"
	"net/http"
	"os"

	"github.com/garyburd/redigo/redis"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

var rc redis.Conn

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
	var err error
	log.Println("Connecting to Redis")
	rc, err = redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatalf("Issues with redis: %s", err)
	}
	defer rc.Close()

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
