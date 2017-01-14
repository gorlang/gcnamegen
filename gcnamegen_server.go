package gcnamegen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type RespNames struct {
	Status string
	Names  []string
}

func init() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/names", namesHandler)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "defaultHandler")
}

func namesHandler(w http.ResponseWriter, r *http.Request) {

	//request_origin := "http://localhost:8888" // dev
	request_origin := "http://knutas.com" // prod

	fmt.Println("namesHandler_")

	respNames := RespNames{}
	var ctx Context
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&ctx)
	if err != nil {
		respNames.Status = "error"
	} else {
		fmt.Println("ctx", ctx)
		respNames = buildRespNames(&ctx)
		fmt.Println("respNames", respNames)
	}

	b, err := json.Marshal(&respNames)
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "\t")
	// TODO handle errors

	w.Header().Add("Access-Control-Allow-Origin", request_origin)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(out.Bytes())
}

/*
	Calls the names generator and returns result
*/

func buildRespNames(ctx *Context) RespNames {

	fmt.Println("buildRespNames")

	ctx.NameCount = 10 // hard coded to limit service

	names := GenerateNames(ctx)
	r := RespNames{"ok", names}

	return r
}
