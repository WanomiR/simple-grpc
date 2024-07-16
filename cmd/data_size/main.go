package main

import (
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	desc "github.com/wanomir/simple-grpc/pkg/note_v1"
	"google.golang.org/protobuf/proto"
)

func main() {
	session := &desc.NoteInfo{
		Title:    gofakeit.BeerName(),
		Content:  gofakeit.IPv4Address(),
		Author:   gofakeit.Name(),
		IsPublic: gofakeit.Bool(),
	}

	dataJson, _ := json.Marshal(session)
	fmt.Printf("\n\ndataJson len %d byte \n%v\n", len(dataJson), dataJson)

	dataPb, _ := proto.Marshal(session)
	fmt.Printf("dataPb len %d byte \n%v\n", len(dataPb), dataPb)
}
