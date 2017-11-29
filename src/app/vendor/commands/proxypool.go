package commands

import (
	"fmt"
	"github.com/henson/ProxyPool/api"
	"github.com/henson/ProxyPool/getter"
	"github.com/henson/ProxyPool/models"
	"github.com/henson/ProxyPool/storage"
	"github.com/henson/ProxyPool/util"
	"github.com/rakyll/statik/fs"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"runtime"
	"sync"
	"time"
)

func check() {
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}
	f, err := statikFS.Open("/scripts/phantomjs_fetcher.js")
	if err != nil {
		panic(err)
	}
	js_name := "./phantomjs_fetcher.js"
	fr, err := os.Open(js_name)
	if err != nil && os.IsNotExist(err) {
		fw, err := os.Create(js_name)
		if err != nil {
			panic(err)
		}
		defer fw.Close()

		_, err = io.Copy(fw, f)
		if err != nil {
			panic(err)
		}
	}
	defer fr.Close()

	fmt.Println("exists")
}

func run(host_listen string, cfg *util.Config) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ipChan := make(chan *models.IP, 2000)
	storage.SetConfig(cfg)
	conn := storage.NewStorage()

	// Start HTTP
	go func() {
		api.Run(host_listen)
	}()

	// Check the IPs in DB
	go func() {
		storage.CheckProxyDB()
	}()

	// Check the IPs in channel
	for i := 0; i < 50; i++ {
		go func() {
			for {
				storage.CheckProxy(<-ipChan)
			}
		}()
	}

	// Start getters to scraper IP and put it in channel
	for {
		x := conn.Count()
		log.Printf("Chan: %v, IP: %v\n", len(ipChan), x)
		if len(ipChan) < 100 {
			go runloop(ipChan)
		}
		time.Sleep(10 * time.Minute)
	}
}

func runloop(ipChan chan<- *models.IP) {
	var wg sync.WaitGroup
	funs := []func() []*models.IP{
		getter.Data5u,
		getter.IP66,
		getter.KDL,
		getter.GBJ,
		getter.Xici,
		getter.XDL,
		getter.IP181,
		//getter.YDL,		//失效的采集脚本，用作系统容错实验
		getter.PLP,
	}
	for _, f := range funs {
		wg.Add(1)
		go func(f func() []*models.IP) {
			temp := f()
			for _, v := range temp {
				ipChan <- v
			}
			wg.Done()
		}(f)
	}
	wg.Wait()
	log.Println("All getters finished.")
}

func setupRunCommand(rootCmd *cobra.Command) {
	var mgo_addr string
	var mgo_db string
	var mgo_col string
	cmd := &cobra.Command{
		Use:   "run",
		Short: "啟動服務",
		Run: func(cmd *cobra.Command, args []string) {
			check()
			run(":8080", &util.Config{
				Mongo: util.MongoConfig{
					Addr:  mgo_addr,
					DB:    mgo_db,
					Table: mgo_col,
					Event: "",
				},
			})
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&mgo_addr, "mgo-addr", "a", "mongodb://127.0.0.1:27017?maxPoolSize=15", "MongoDB連線參數")
	flags.StringVarP(&mgo_db, "mgo-db", "d", "proxypool", "MongoDB Database Name")
	flags.StringVarP(&mgo_col, "mgo-col", "c", "pool", "MongoDB Collection name")

	rootCmd.AddCommand(cmd)
}

func setupCheckCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "初始化",
		Run: func(cmd *cobra.Command, args []string) {
			// fmt.Println("TODO")
			// fmt.Println("check config.json")
			// fmt.Println("check pjs")
			check()
		},
	}
	rootCmd.AddCommand(cmd)
}

func SetupProxyPoolCommand(parentCmd *cobra.Command) {
	rootCmd := &cobra.Command{
		Use:   "proxy-pool",
		Short: "ProxyPool Service",
	}

	// setupCheckCommand(rootCmd)
	setupRunCommand(rootCmd)

	parentCmd.AddCommand(rootCmd)
}
