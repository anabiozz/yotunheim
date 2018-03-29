package main

import (
	"fmt"
	"heimdall_project/yotunheim/backend/common"
	"heimdall_project/yotunheim/backend/common/datastore"
	"heimdall_project/yotunheim/backend/common/utility"
	"heimdall_project/yotunheim/backend/endpoints"
	"heimdall_project/yotunheim/backend/handlers"
	"heimdall_project/yotunheim/backend/internal/config"
	_ "heimdall_project/yotunheim/backend/metrics/all"
	"log"
	"os"

	"github.com/kataras/iris"
)

func handleError(key string, err error, message string) {
	log.SetPrefix(fmt.Sprintf("[logID: %v]: ", key))
	log.Printf("%#v", err)
	log.Printf("[%v] %v", key, message)
}

var (
	bugMsg = "There was an unexpected issue; please report this as a bug."
)

func main() {
	// set up logs parameters
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	// Iris parameters
	app := iris.Default()
	tmpl := iris.HTML("../public", ".html")
	tmpl.Layout("index.html")
	app.RegisterView(tmpl)

	// get uuid
	uuid, err := utility.GenerateUUID()
	if err != nil {
		if _, ok := err.(utility.UUIDError); ok {
			bugMsg = err.Error()
		}
		handleError(uuid, err, bugMsg)
	}

	// Cretate new datastore
	db, err := datastore.NewDatastore(datastore.INFLUXDB, "http://localhost:8086")
	if err != nil {
		if _, ok := err.(datastore.DatastoreErr); ok {
			bugMsg = err.Error()
		}

		handleError(uuid, err, bugMsg)
	}

	env := common.Env{DB: db}

	// Create new config
	newConfig := config.NewConfig()
	// Filling new config getting data from default config
	err = newConfig.LoadConfig()
	inputs := newConfig.InputFilters["inputs"].([]interface{})

	// Filling InputFilters
	for _, value := range inputs {
		newConfig.AddInput(value.(string))
	}

	if len(newConfig.Inputs) == 0 {
		log.Fatalf("ERROR: no inputs found, did you provide a valid config file?")
	}

	// Handlers
	app.Handle("GET", "/", handlers.HomeHandler)
	app.Handle("GET", "/dashboard", handlers.DashboardHandler)

	//Endpoints
	app.Handle("GET", "/api/get-json", endpoints.GetJSONnEndpoint(&env, newConfig))

	app.StaticWeb("/", "../public")
	app.Run(
		iris.Addr("localhost:8080"),
		// disables updates:
		iris.WithoutVersionChecker,
		// skip err server closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)
}
