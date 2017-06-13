package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type statusCode struct {
	Code    int
	Descrip string
}

type statusCodes []statusCode

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	var rcvd string
	var data statusCodes
	rcvd = `[{"Code":200,"Descrip":"StatusOK"},{"Code":301,"Descrip":"StatusMovedPermanently"},{"Code":302,"Descrip":"StatusFound"},{"Code":303,"Descrip":"StatusSeeOther"},{"Code":307,"Descrip":"StatusTemporaryRedirect"},{"Code":400,"Descrip":"StatusBadRequest"},{"Code":401,"Descrip":"StatusUnauthorized"},{"Code":402,"Descrip":"StatusPaymentRequired"},{"Code":403,"Descrip":"StatusForbidden"},{"Code":404,"Descrip":"StatusNotFound"},{"Code":405,"Descrip":"StatusMethodNotAllowed"},{"Code":418,"Descrip":"StatusTeapot"},{"Code":500,"Descrip":"StatusInternalServerError"}]`

	err := json.Unmarshal([]byte(rcvd), &data)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(data)
	for _, v := range data {
		fmt.Println(v.Code, "-", v.Descrip)
	}
	// Marshal data and display it on page
	json, _ := json.Marshal(data)
	w.Write(json)

}
