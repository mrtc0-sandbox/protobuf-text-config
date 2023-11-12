package main

import (
	"fmt"
	"log"
	"os"

	config "github.com/mrtc0-sandbox/protobuf-text-config/pkg/config/proto"
	"google.golang.org/protobuf/encoding/prototext"
)

func main() {
	cfg, err := os.ReadFile("config/sample.textproto")
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}

	appCfg := &config.ApplicationConfig{}
	if err := prototext.Unmarshal(cfg, appCfg); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("verbose: %#v\n", appCfg.Verbose)

	for _, provider := range appCfg.Provider {
		fmt.Println(provider.Name, provider.Url)
	}
}
