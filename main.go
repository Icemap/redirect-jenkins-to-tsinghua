package main

import (
	"arbiter_agent/log"
	"flag"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var srcLoc = "https://mirrors.tuna.tsinghua.edu.cn/jenkins/updates/current/update-center.json"
var echoServer *echo.Echo

func main () {
	port := flag.String("p", "18081", "Service Port")
	flag.StringVar(&srcLoc, "s", "https://mirrors.tuna.tsinghua.edu.cn/jenkins/updates/current/update-center.json", "Src Jenkins Update Center Json")
	help := flag.Bool("h", false, "Man")

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	echoServer = echo.New()
	echoServer.HideBanner = true
	echoServer.Use(middleware.Recover())
	echoServer.GET("/json", proxyJenkinsJson)
	echoServer.Logger.Info(echoServer.Start(":" + *port))
}

func proxyJenkinsJson(context echo.Context) error {
	resp, err := http.Get(srcLoc)
	if err != nil {
		log.Error(err.Error())
		return context.String(http.StatusInternalServerError, "Source URL can't reached.")
	}

	baData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err.Error())
		return context.String(http.StatusInternalServerError, "Source URL data can't be read.")
	}

	sData := string(baData)
	sData = strings.Replace(sData, "http://updates.jenkins-ci.org/download", "https://mirrors.tuna.tsinghua.edu.cn/jenkins", -1)
	sData = strings.Replace(sData, "http://www.google.com/", "http://www.baidu.com/", -1)

	return context.String(http.StatusOK, sData)
}


