package main

/*
Following will create a config file that looks like this:

	{
		"Debug": 0,
		"AdminName": "",
		"AdminPass": "",
		"Srv": {
			"ListenAddr": "",
			"IpAddr": "",
			"Hostname": ""
		}
	}
*/

import (
	"flag"
	"fmt"

	//"log"
	cfg "./pkg"
)

var Cfg struct {
	Host struct {
		NetworkId uint8  // 0: local, 1: mainnet, 4: rinkeby 5: goreli
		HTTP      string // http://127.0.0.1:7545
		WS        string // http://127.0.0.1:7545
	}
}

var (
	cfg_file = flag.String("c", "", "Config file (*.json)")
	help     = flag.Bool("h", false, "Display help")
)

func main() {

	cfg_file := "./abc"

	if err := cfg.Load(cfg_file, &Cfg); err == nil {

		Cfg.ActiveNetworkId = 0 // networkid
		Cfg.Host.HTTP = "http://localhost"
		Cfg.Host.WS = "ws://localhost"

		if err := cfg.Save(cfg_file, &Cfg); err == nil {
			fmt.Println(Cfg)
		}

	}

}

// func main() {
// 	flag.Parse()
// 	if *help == true || *cfg_file == "" {
// 		flag.PrintDefaults()
// 		os.Exit(0)
// 	}

// 	if err := cfg.Load(*cfg_file, &Cfg); err != nil {
// 		fmt.Println(err)
// 		if err = cfg.Save(*cfg_file, &Cfg); err != nil {
// 			fmt.Println("cfg.Save error: ", err.Error())
// 			os.Exit(1)
// 		} else {

// 			fmt.Println("Please edit your config file at:\n\n\t", *cfg_file)
// 			os.Exit(0)
// 		}
// 	}

// 	Cfg.Debug = 9
// 	cfg.Save(*cfg_file, &Cfg)

// 	// //for read
// 	// if err := cfg.Load(*cfg_file, &Cfg); err == nil {
// 	// 	abc := &Cfg
// 	// 	fmt.Println(abc)
// 	// }

// }
