package main

import (
	"encoding/json"
	"flag"
	"fmt"
	st "godis_server/storage"
	"log"
	"net/http"
	"strings"
)

var modeFlag string
var portFlag string
var storageRep st.Storage

func init() {
	readStartupFlags()
	configureMode()
}

func main() {
	startHTTPServer()
}

func startHTTPServer() {
	http.HandleFunc("/storage", handleRequest)
	http.HandleFunc("/storage/keys", handleKeysRequest)
	err := http.ListenAndServe(":"+portFlag, nil)
	if err != nil {
		log.Fatal("Error! Couldn't start the server: ", err)
	}
}

func configureMode() {
	if modeFlag == "disk" {
		storageRep = st.NewDiskStorage()
	} else {
		storageRep = st.NewInmemoryStorage()
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		doGet(w, r)
	case "POST":
		doPost(w, r)
	case "DELETE":
		doDelete(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "This method is not allowed")
	}
}

func handleKeysRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		doGetKeys(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "This method is not allowed")
	}
}

func doGet(w http.ResponseWriter, r *http.Request) {
	keyParams := r.URL.Query()["key"]
	if len(keyParams) > 1 {
		http.Error(w, "Please, pass only one key as param or don't pass it at all (for getting all data)", http.StatusBadRequest)
	} else if len(keyParams) == 1 {
		key := keyParams[0]
		tuple, ok := storageRep.Get(key)
		if ok {
			jsonTuple, _ := json.Marshal(tuple)
			fmt.Fprintf(w, string(jsonTuple))
		} else {
			http.Error(w, "", http.StatusNotFound)
		}
	} else {
		jsonTuples, _ := json.Marshal(storageRep.GetAll())
		fmt.Fprintf(w, string(jsonTuples))
	}
}

func doPost(w http.ResponseWriter, r *http.Request) {
	body := json.NewDecoder(r.Body)
	t := &st.Tuple{}
	err := body.Decode(t)
	if err != nil {
		http.Error(w, "Pass the data in next format:", http.StatusBadRequest)
	}
	jsonTuple, _ := json.Marshal(storageRep.Set(*t))
	fmt.Fprintf(w, string(jsonTuple))
}

func doDelete(w http.ResponseWriter, r *http.Request) {
	keyParams := r.URL.Query()["key"]
	if len(keyParams) != 1 {
		http.Error(w, "Please, pass only one key as param", http.StatusBadRequest)
	} else {
		key := keyParams[0]
		storageRep.Del(key)
	}
}

func doGetKeys(w http.ResponseWriter, r *http.Request) {
	matchParams := r.URL.Query()["match"]
	if len(matchParams) != 1 || !strings.Contains(matchParams[0], "*") {
		http.Error(w, "Please, pass only one match param and use '*' for searching keys", http.StatusBadRequest)
	} else {
		keys, ok := storageRep.FindKeys(matchParams[0])
		if ok {
			jsonKeys, _ := json.Marshal(keys)
			fmt.Fprintf(w, string(jsonKeys))
		} else {
			http.Error(w, "", http.StatusNotFound)
		}
	}
}

func readStartupFlags() {
	const (
		shorthandDesc   = "shorthand flag; "
		modeFlagDefault = "memory"
		modeFlagDesc    = "storage mode of the server: 'memory' or 'disk'"
		portFlagDefault = "9090"
		portFlagDesc    = "port on which the server runs"
	)
	flag.StringVar(&modeFlag, "mode", modeFlagDefault, modeFlagDesc)
	flag.StringVar(&modeFlag, "m", modeFlagDefault, shorthandDesc+modeFlagDesc)
	flag.StringVar(&portFlag, "port", portFlagDefault, portFlagDesc)
	flag.StringVar(&portFlag, "p", portFlagDefault, shorthandDesc+portFlagDesc)
	flag.Parse()
	fmt.Println("Mode: ", modeFlag)
	fmt.Println("Port: ", portFlag)
}
