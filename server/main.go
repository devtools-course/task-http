package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type PairDTO struct {
	X int64 `json:"x" xml:"x" yaml:"x"`
	Y int64 `json:"y" xml:"y" yaml:"y"`
}

type XmlSchema struct {
	XmlName xml.Name `xml:"req"`
	Version string   `xml:"version,attr"`
	Pair    PairDTO  `xml:"pair"`
}

type YamlSchema struct {
	Pair PairDTO `yaml:"pair"`
}

type response struct {
	Result int64 `json:"result" xml:"result" yaml:"result"`
}

func hello(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintln(w, "Hello, world!")
}

func loop(w http.ResponseWriter, req *http.Request) {
	time.Sleep(15 * time.Second)
	w.WriteHeader(http.StatusLoopDetected)
	_, _ = fmt.Fprintln(w, "Exit")
}

func incr(w http.ResponseWriter, req *http.Request) {
	num := req.URL.Query().Get("n")

	numInt, err := strconv.Atoi(num)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintln(w, err)
		return
	}

	w.WriteHeader(200)
	_, _ = fmt.Fprintln(w, numInt+1)
}

func divideJson(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" && req.Header["Content-Type"][0] != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintln(w, "Invalid method")
		return
	}

	var pair PairDTO

	err := json.NewDecoder(req.Body).Decode(&pair)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintln(w, err)
		return
	}

	res, err := divide(w, pair)
	if err != nil {
		return
	}

	yml, err := json.Marshal(response{Result: res})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintln(w, err)
		return
	}

	_, _ = fmt.Fprintln(w, string(yml))
}

func divideXml(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" && req.Header["Content-Type"][0] != "application/xml" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintln(w, "Invalid method")
		return
	}

	var data XmlSchema

	body, err := io.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintln(w, err)
		return
	}

	err = xml.Unmarshal(body, &data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintln(w, err)
		return
	}

	if data.Version != "1.1" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintln(w, "Invalid version")
		return
	}

	res, err := divide(w, data.Pair)
	if err != nil {
		return
	}

	yml, err := xml.Marshal(response{Result: res})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintln(w, err)
		return
	}

	_, _ = fmt.Fprintln(w, string(yml))
}

func divideYaml(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" && req.Header["Content-Type"][0] != "application/yaml" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintln(w, "Invalid method")
		return
	}

	var data YamlSchema

	body, err := io.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintln(w, err)
		return
	}

	err = yaml.Unmarshal(body, &data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintln(w, err)
		return
	}

	res, err := divide(w, data.Pair)
	if err != nil {
		return
	}

	yml, err := yaml.Marshal(response{Result: res})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintln(w, err)
		return
	}

	_, _ = fmt.Fprintln(w, string(yml))
}

func divide(w http.ResponseWriter, pair PairDTO) (int64, error) {
	if pair.Y == 0 {
		w.WriteHeader(http.StatusNotImplemented)
		_, err := fmt.Fprintln(w, "Cannot divide by zero")
		return 0, err
	}

	return pair.X / pair.Y, nil
}

func redirect(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/r1", 301)
}

func redirect1(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/sub-inner", 301)
}

func redirect2(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/inner", 301)
}

func inner(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintln(w, "Inner API")
	if err != nil {
		log.Fatal(err)
	}
}

func file(w http.ResponseWriter, req *http.Request) {
	path := "answer.txt"
	w.Header().Set("Content-Disposition", "attachment; filename=\""+path+"\"")
	http.ServeFile(w, req, "/"+path)
}

func hidden(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/f", 301)
}

func erro(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusServiceUnavailable)
	_, _ = fmt.Fprintln(w, "Service Unavailable")
}

func main() {
	addr := os.Getenv("ADDR")

	http.HandleFunc("/", hello)
	http.HandleFunc("/i", incr)
	http.HandleFunc("/l", loop)
	http.HandleFunc("/dj", divideJson)
	http.HandleFunc("/dx", divideXml)
	http.HandleFunc("/dy", divideYaml)
	http.HandleFunc("/r", redirect)
	http.HandleFunc("/r1", redirect1)
	http.HandleFunc("/sub-inner", redirect2)
	http.HandleFunc("/inner", inner)
	http.HandleFunc("/f", file)
	http.HandleFunc("/h", hidden)
	http.HandleFunc("/e", erro)

	log.Fatal(http.ListenAndServe(":"+addr, nil))
}
