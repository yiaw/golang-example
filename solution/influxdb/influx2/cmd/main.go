package main

import (
	"flag"
	"fmt"
	"influx2/config"
	"influx2/internal/agent"
	"influx2/internal/app"
	"influx2/internal/pkg/argument"

	"github.com/yiaw/yiaw-go/log"
)

const usage = `Usage of influx2-write
--file                       read query from file
--query                      read stdin from query
-w, --write                  write mode 
    --itemname-field
    --itemname-tag
    --stat
-h, --help                   prints help informations
`

func main() {
	var arg argument.Arguments

	flag.BoolVar(&arg.Write, "write", false, "main start write mode")
	flag.BoolVar(&arg.Write, "w", false, "main start write mode")
	flag.BoolVar(&arg.Read, "read", false, "main start write mode")
	flag.BoolVar(&arg.Read, "r", false, "main start write mode")
	flag.BoolVar(&arg.ReadQuery, "query", true, "main start write mode")
	flag.BoolVar(&arg.ReadQuery, "q", true, "main start write mode")
	flag.StringVar(&arg.ReadFile, "file", "", "read query file ")
	flag.StringVar(&arg.ReadFile, "f", "", "read query file ")
	flag.BoolVar(&arg.W_field, "itemname-field", false, "item name set field")
	flag.BoolVar(&arg.W_tag, "itemname-tag", false, "item name set tag")
	flag.BoolVar(&arg.W_stat, "stat", false, "stat struct write point")
	flag.Usage = func() { fmt.Print(usage) }

	flag.Parse()

	if !arg.ValidCheck() {
		flag.Usage()
	}

	l := log.NewLog()
	log.SetGlobalLog(l)

	log.SVC("%+v\n", arg)

	cfg := config.LoadConfig()

	appService := app.NewApp(agent.NewAgentService(arg, cfg.InfluxDB))
	err := appService.Start(cfg, arg)
	if err != nil {
		log.ERR("app start failed. err=%v\n", err)
	}
}
