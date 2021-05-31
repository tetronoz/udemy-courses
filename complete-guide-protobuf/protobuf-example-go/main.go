package main

import (
	"fmt"
	"io/ioutil"
	"log"

	complexpb "example.com/protobuf-example-go/src/complex"
	enumpb "example.com/protobuf-example-go/src/enum_example"
	simplepb "example.com/protobuf-example-go/src/simple"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	//"github.com/golang/protobuf/jsonpb"
	//"github.com/golang/protobuf/proto"
)

func main() {

	sm := doSimple()
	readAndWriteDemo(sm)
	jsonDemo(sm)

	doEnum()
	doComplex()

}

func doComplex() {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "First message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   2,
				Name: "Second message",
			},
			&complexpb.DummyMessage{
				Id:   3,
				Name: "Third message",
			},
		},
	}

	fmt.Println(cm)
}

func doEnum() {
	em := enumpb.EnumMessage{
		Id:           42,
		DayOfTheWeek: enumpb.DayOfTheWeek_THURSDAY,
	}

	fmt.Println(em)
}

func jsonDemo(sm proto.Message) {
	smAsString := toJSON(sm)

	fmt.Println(smAsString)

	sm2 := simplepb.SimpleMessage{}
	fromJSON(smAsString, &sm2)
}

func fromJSON(s string, pb proto.Message) error {
	b := []byte(s)
	if err := protojson.Unmarshal(b, pb); err != nil {
		log.Fatalln("Cant unmarshal from string")
		return err
	}

	return nil
}

func toJSON(pb proto.Message) string {
	marshaler := protojson.MarshalOptions{}
	out, err := marshaler.Marshal(pb)
	if err != nil {
		log.Fatalln("Cant marshal to JSON")
		return ""
	}
	return string(out)
}

func readAndWriteDemo(sm proto.Message) {

	writeToFile("simple.bin", sm)
	newsm, err := readFromFile("simple.bin")

	if err != nil {
		log.Fatalln("Faild to read from simple.bin")
	}

	fmt.Println(newsm.GetName())
}

func doSimple() *simplepb.SimpleMessage {

	sm := simplepb.SimpleMessage{
		Id:        12345,
		IsSimple:  true,
		Name:      "Simple message",
		SampleInt: []int32{1, 2, 3, 4, 5, 6, 7, 8},
	}

	return &sm
}

func readFromFile(filename string) (*simplepb.SimpleMessage, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("Can't read")
		return nil, err
	}

	sm := &simplepb.SimpleMessage{}

	err = proto.Unmarshal(data, sm)
	if err != nil {
		log.Fatalln("Failed to unmarshal")
		return nil, err
	}

	return sm, nil
}

func writeToFile(filename string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Fatal error")
		return err
	}
	if err := ioutil.WriteFile(filename, out, 0644); err != nil {
		log.Fatalln("Failed writing to file")
		return err
	}

	return nil
}
