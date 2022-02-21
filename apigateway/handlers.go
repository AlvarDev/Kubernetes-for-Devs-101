package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	pb "apigatewayservice/pb"

	"google.golang.org/protobuf/encoding/protojson"
)

func (fe *frontendServer) getNewIdHandler(w http.ResponseWriter, r *http.Request) {

	req := &pb.GetNewIdRequest{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = protojson.Unmarshal(body, req)
	if err != nil {
		fmt.Println(err)
	}

	newRemainder, err := fe.getNewId(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	m := protojson.MarshalOptions{
		Indent:          "  ",
		EmitUnpopulated: true,
	}

	jsonBytes, _ := m.Marshal(newRemainder)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
