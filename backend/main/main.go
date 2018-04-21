package main

import (
	"fmt"
	"log"
	"os"

	"github.com/anabiozz/yotunheim/backend/common"
	"github.com/anabiozz/yotunheim/backend/common/datastore"
	"github.com/anabiozz/yotunheim/backend/common/utility"
	"github.com/anabiozz/yotunheim/backend/endpoints"
	"github.com/anabiozz/yotunheim/backend/handlers"
	_ "github.com/anabiozz/yotunheim/backend/metrics/all"

	"github.com/anabiozz/yotunheim/backend/internal/config"

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

	//****************************************************************************************//
	// set up logs parameters
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	// Iris parameters
	app := iris.Default()
	tmpl := iris.HTML("../go/src/github.com/anabiozz/yotunheim/backend/public", ".html")
	tmpl.Layout("index.html")
	app.RegisterView(tmpl)

	//****************************************************************************************//

	// get uuid
	uuid, err := utility.GenerateUUID()
	if err != nil {
		if _, ok := err.(utility.UUIDError); ok {
			bugMsg = err.Error()
		}
		handleError(uuid, err, bugMsg)
	}

	//****************************************************************************************//

	// Cretate new datastore
	db, err := datastore.NewDatastore(datastore.INFLUXDB, "http://influxdb:8086")
	if err != nil {
		if _, ok := err.(datastore.DatastoreErr); ok {
			bugMsg = err.Error()
		}

		handleError(uuid, err, bugMsg)
	}

	fmt.Println(db)

	//****************************************************************************************//

	// Create new config
	newConfig := config.NewConfig()
	// Filling new config getting data from default config
	err = newConfig.LoadConfig()
	fmt.Println(newConfig.InputFilters)
	inputs := newConfig.InputFilters["inputs"].([]interface{})

	// Filling InputFilters
	for _, value := range inputs {
		log.Println(value)
		newConfig.AddInput(value.(string))
	}

	if len(newConfig.Inputs) == 0 {
		log.Fatalf("ERROR: no inputs found, did you provide a valid config file?")
	}

	//****************************************************************************************//

	env := common.Env{DB: db}

	// Handlers
	app.Handle("GET", "/", handlers.HomeHandler)
	app.Handle("GET", "/dashboard", handlers.DashboardHandler)

	//Endpoints
	app.Handle("GET", "/api/get-json", endpoints.GetJSONnEndpoint(&env, newConfig))

	app.StaticWeb("/", "../go/src/github.com/anabiozz/yotunheim/backend/public")
	app.Run(
		iris.Addr("yotunheim:8080"),
		// disables updates:
		iris.WithoutVersionChecker,
		// skip err server closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)
}
