package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/anabiozz/yotunheim/backend/common"
	"github.com/anabiozz/yotunheim/backend/common/datastore"
	"github.com/anabiozz/yotunheim/backend/common/utility"
	"github.com/anabiozz/yotunheim/backend/endpoints"
	"github.com/anabiozz/yotunheim/backend/handlers"
	_ "github.com/anabiozz/yotunheim/backend/metrics/all"
	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/mux"

	"github.com/anabiozz/yotunheim/backend/internal/config"
)

var (
	hub *Hub
)

func main() {

	hub = NewHub()
	go hub.Run()

	//****************************************************************************************//

	// set up logs parameters
	// log.SetOutput(os.Stdout)
	// log.SetFlags(log.Ltime | log.LUTC)

	//****************************************************************************************//

	// get uuid
	uuid, err := utility.GenerateUUID()
	if err != nil {
		if _, ok := err.(utility.UUIDError); ok {
			utility.BugMsg = err.Error()
		}
		utility.HandleError(uuid, err, utility.BugMsg)
	}

	//****************************************************************************************//

	// Cretate new datastore
	db, err := datastore.NewDatastore(datastore.INFLUXDB, "http://influxdb:8086")
	if err != nil {
		if _, ok := err.(datastore.Err); ok {
			utility.BugMsg = err.Error()
		}

		utility.HandleError(uuid, err, utility.BugMsg)
	}

	//****************************************************************************************//

	// Create new config
	newConfig := config.NewConfig()
	// Filling new config getting data from default config
	err = newConfig.LoadConfig()
	// fmt.Println(newConfig.InputFilters)
	inputs := newConfig.InputFilters["inputs"].([]interface{})

	// Filling InputFilters
	for _, value := range inputs {
		log.Println(value)
		newConfig.AddInput(value.(string))
	}

	if len(newConfig.Inputs) == 0 {
		log.Fatalf("ERROR: no inputs found, did you provide a valid config file?")
	}

	//**********************************File notifier for bundle reloading**********************//

	go func() {

		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			fmt.Println("ERROR", err)
		}
		defer watcher.Close()

		// out of the box fsnotify can watch a single file, or a single directory
		if err := watcher.Add("/go/src/github.com/anabiozz/yotunheim/backend/public/bundle.js"); err != nil {
			fmt.Println("ERROR", err)
		}

		for {
			select {
			// watch for events
			case event := <-watcher.Events:

				fmt.Printf("EVENT! %#v\n", event)

				message := bytes.TrimSpace([]byte("reload"))
				hub.Broadcast <- message

			// watch for errors
			case err := <-watcher.Errors:
				fmt.Println("ERROR", err)
			}
		}
	}()

	// ********************************************************************************************//

	env := common.Env{DB: db}

	// Mux
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.DashboardHandler(w, r)
	}).Methods("GET")

	router.HandleFunc("/api/get-json", func(w http.ResponseWriter, r *http.Request) {
		endpoints.GetJSONnEndpoint(w, r, &env, newConfig)
	}).Methods("GET")

	router.HandleFunc("/reload", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(hub, w, r)
	})

	router.PathPrefix("/").Handler(http.StripPrefix("/",
		http.FileServer(http.Dir("../go/src/github.com/anabiozz/yotunheim/backend/public"))))

	srv := &http.Server{
		Addr: "yotunheim:8888",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router, // Pass our instance of gorilla/mux in.
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
