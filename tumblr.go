package main

import (
	"github.com/qbhy/go-utils"
	"path"
	"fmt"
	"strings"
	"tumblr-crawler/downloader"
	"runtime"
	"sync"
	io "io/ioutil"
)

func ParseSites(filename string) string {
	data, _ := io.ReadFile(filename)
	wrapSites := string(data)
	wrapSites = strings.Replace(wrapSites, "\t", ",", -1)
	wrapSites = strings.Replace(wrapSites, "\n", ",", -1)
	wrapSites = strings.Replace(wrapSites, "\r", ",", -1)
	wrapSites = strings.Replace(wrapSites, " ", ",", -1)
	return wrapSites
}

var waitGroup sync.WaitGroup //定义一个同步等待的组

var num = 14 //定义一工并发多少数量
var cnum chan int

func main() {

	config := downloader.NewConfig()
	currentPath := utils.CurrentPath()
	sites := ""
	var proxies downloader.ProxyConfig

	fmt.Println(proxies)
	fmt.Println(sites)

	proxyPath := path.Join(currentPath, "proxies.json")
	if exists, _ := utils.PathExists(proxyPath); exists {
		proxies = downloader.ProxyConfig{}
		config.Load(proxyPath, &proxies)
		fmt.Println(proxies)
	}

	sitesPath := path.Join(currentPath, "sites.txt")
	if exists, _ := utils.PathExists(sitesPath); exists {
		fmt.Println(ParseSites(sitesPath))
	}

	maxProcesses := runtime.NumCPU() //获取cpu个数
	runtime.GOMAXPROCS(maxProcesses) //限制同时运行的goroutines数量

	// 下面这个for循环的意义就是利用信道的阻塞，一直从信道里取数据，直到取得跟并发数一样的个数的数据，则视为所有goroutines完成。

	site := downloader.NewSite("tbr91677", proxies)

	site.StartDownload()

	fmt.Println("WE DONE!!!")
}