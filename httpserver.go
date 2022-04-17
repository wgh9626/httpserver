package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	err := flag.Set("v", "4")
	if err != nil {
		return
	}
	glog.V(4).Infof("starting http server")
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthHandler)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("start http server failed", err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("enter root handler")
	fmt.Printf("request path is %s\n", r.URL.Path)

	if r.URL.Path != "/" {
		fmt.Printf("request path %v is not found! %v\n", r.URL.Path, http.StatusNotFound)
		io.WriteString(w, "request path is not found! 404!\n")
		return
	}

	goVersion := os.Getenv("VERSION")
	if goVersion != "" {
		w.Header().Add("VERSION", goVersion)
	} else {
		err := os.Setenv("VERSION", "go1.17.5")
		if err != nil {
			return
		}
		version := os.Getenv("VERSION")
		w.Header().Add("VERSION", version)
	}

	for k, v := range r.Header {
		//io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
		w.Header().Set(k, strings.Join(v, "; "))
	}
	//for k, v1 := range header {
	//	for _, v2 := range v1 {
	//		writer.Header().Set(k,v2)
	//	}
	//}
	fmt.Fprintf(w, "Hello, wgh!")

	ip, port, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil {
		fmt.Printf("request host is %v:%v\t", ip, port)
		fmt.Printf("http.status=%d\n", http.StatusOK)
	} else {
		fmt.Printf("requset host is wrong: %v\n", err.Error())
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("enter health server")
	
	for k, v := range r.Header {
		w.Header().Set(k, strings.Join(v, "; "))
		
	goVersion := os.Getenv("VERSION")
	if goVersion != "" {
		w.Header().Add("VERSION", goVersion)
	} else {
		err := os.Setenv("VERSION", "go1.17.5")
		if err != nil {
			return
		}
		version := os.Getenv("VERSION")
		w.Header().Add("VERSION", version)
	}
	fmt.Printf("requset path is %v\n", r.URL.Path)
	fmt.Fprintf(w, "This is healthz test html\t")
	io.WriteString(w, "200 OK\n")
	fmt.Printf("res_code is %v", http.StatusOK)
}
