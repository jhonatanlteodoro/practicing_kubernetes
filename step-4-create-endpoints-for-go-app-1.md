# Modify your go app 1 

## Endpoints to read and write on some file
### Algo como isso:
```go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var FILE_NAME = os.Getenv("REQUEST_FILE_NAME")
var DEFEALT_FILE_NAME = "/tmp/upa_cavalinho.json"

func GetFileName() string {
	file := FILE_NAME
	if file == "" {
		file = DEFEALT_FILE_NAME
	}
	return file
}

type Data struct {
	Request_id int   `json:"request_id"`
	Date       int64 `json:"date"`
}

type Requests struct {
	Requests []Data `json:"requests"`
}

func ReadRequestFile(file_name string) (Requests, error) {
	file, err := ioutil.ReadFile(file_name)
	if err != nil {
		fmt.Printf("Fail attempt read file: %s, error: %s \n", file_name, err)
		return Requests{}, err
	}
	data := Requests{}
	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		fmt.Printf("Fail attempt convert data into Request struct. Erro: %s \n", err)
		return Requests{}, err
	}
	return data, nil
}

func WriteRequestToFile(file_name string) error {
	data, err := ReadRequestFile(file_name)

	var new_request_id int
	new_unix_date := time.Now().Unix()
	if len(data.Requests) > 0 {
		new_request_id = data.Requests[len(data.Requests)-1].Request_id + 1
	} else {
		new_request_id = 1
	}

	data.Requests = append(
		data.Requests,
		Data{
			Request_id: new_request_id,
			Date:       new_unix_date,
		},
	)

	json_data, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Printf("Fail attempt marshelling data. Error: %s \n", err)
		return err
	}

	err = ioutil.WriteFile(file_name, json_data, 0644)
	if err != nil {
		fmt.Printf("Fail attempt write file: %s. Error: %s", file_name, err)
	}
	return err
}

func ReadRequestJsonFileHandler(rw http.ResponseWriter, r *http.Request) {
	data, _ := ReadRequestFile(GetFileName())
	json_data, _ := json.Marshal(data)
	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(
		[]byte(json_data),
	)
}

func WriteRequestJsonFileHandler(rw http.ResponseWriter, r *http.Request) {
	err := WriteRequestToFile(GetFileName())
	rw.Header().Set("Content-type", "application/json")
	
	if err != nil {
		rw.WriteHeader(http.StatusNoContent)
		rw.Write(
			[]byte(`{"request_status": "not created"}`),
		)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write(
		[]byte(`{"request_status": "created"}`),
	)
}

func main() {
	http.HandleFunc("/read", ReadRequestJsonFileHandler)
	http.HandleFunc("/write", WriteRequestJsonFileHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## Endpoint to use as helthcheck

### Algo como isso
```go
func HelthCheckHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(
		[]byte(`{"status": "Ok"}`),
	)
}

func main() {
	...
    ...
	http.HandleFunc("/ready", HelthCheckHandler)
	...
}
```
