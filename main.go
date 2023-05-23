package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println(`
██████╗ ██╗   ██╗███████╗███████╗██╗   ██╗
██╔══██╗██║   ██║██╔════╝██╔════╝╚██╗ ██╔╝
██████╔╝██║   ██║█████╗  █████╗   ╚████╔╝ 
██╔══██╗██║   ██║██╔══╝  ██╔══╝    ╚██╔╝  
██████╔╝╚██████╔╝██║     ██║        ██║   
╚═════╝  ╚═════╝ ╚═╝     ╚═╝        ╚═╝                                          
                in golang1!!1!!!
	`)

	http.HandleFunc("/", proxyHandler)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	log.Printf("Server has been started! Listening on port %s\n", PORT)
	log.Printf("Link to view: https://%s.%s.repl.co\n", os.Getenv("REPL_SLUG"), os.Getenv("REPL_OWNER"))

	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	urlParam := r.URL.String()
	mainurl := "https://google.com"
  	url := mainurl + urlParam

	req, err := http.NewRequest(r.Method, url, nil)
	if err != nil {
		log.Printf("Error creating request: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Origin", mainurl)
	req.Header.Set("Access-Control-Allow-Origin", "*")
	req.Header.Set("Access-Control-Allow-Methods", "*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error fetching asset: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading asset body: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	for key, value := range resp.Header {
		w.Header().Set(key, value[0])
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}

