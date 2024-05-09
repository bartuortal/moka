package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/bartuortal/moka/configParser"
)

var (
	portFlag       = flag.Int("port", 8080, "port the server will run on")
	configFileFlag = flag.String("config", "./mokaConf.yaml", "config file to run")
)

func main() {
	var handlerList []configParser.SimpleHandler

	fmt.Printf("using config file %s\n", *configFileFlag)
	file, err := os.Open(*configFileFlag)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	configParser, err := configParser.NewConfigParser(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	handlerList = configParser.Endpoints

	for _, handler := range handlerList {
		http.HandleFunc(fmt.Sprintf("/%s", handler.Endpoint), func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte(handler.Content + "\n"))
			if err != nil {
				fmt.Printf("error while writing message: %s", err)
			}
		})
	}

	fmt.Printf("using port %d\n", *portFlag)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", *portFlag), nil); err != nil {
		fmt.Println("error on listen and serve")
		return
	}
}
