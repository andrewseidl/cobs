package backend

import (
	"log"
	"net/http"
	"os"

	"github.com/garyburd/redigo/redis"

	"code.google.com/p/go-uuid/uuid"
	"github.com/codegangsta/negroni"
	"github.com/gophergala/cobs/instrumenter"
	"github.com/gorilla/mux"
)

var rc redis.Conn

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("hello"))
}

func BuildTarballHandler(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//	case "POST":
	//		//var buf []byte
	//		file, _, _ := r.FormFile("file")
	//		defer file.Close()
	//		data, _ := ioutil.ReadAll(file)
	//		//file.Read(buf)
	//		rc.Do("SET", mux.Vars(r)["imageid"], data)
	default:
		data, _ := redis.Bytes(rc.Do("GET", mux.Vars(r)["imageid"]))
		rw.Write(data)
	}
}

func BuildDockerfileHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte(mux.Vars(r)["imageid"]))
}

func BuildStatusHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte(mux.Vars(r)["imageid"]))
}

func BuildHandler(rw http.ResponseWriter, r *http.Request) {
	//imageid := mux.Vars(r)["imageid"]
	switch r.Method {
	case "POST":
		repository := r.FormValue("repository")
		imageid := uuid.New()
		go instrumenter.Run(repository)
		rw.Write([]byte(imageid))
	default:
		data, _ := redis.Bytes(rc.Do("GET", "tarball"))
		rw.Write(data)
	}
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
	r.HandleFunc("/api/v1/build/", BuildHandler)

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
