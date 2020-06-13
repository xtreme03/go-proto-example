package main

import ( 
	"fmt"
	"log"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"go-proto-example/src/simple" 
	"go-proto-example/src/enum"
	"go-proto-example/src/complex"
	"io/ioutil"

)

func main(){
	//generates a  sample protobuf message
	msg:=generateSimple()
	//function to write the message to a file
	writeToFile("simple.bin",msg)
	//function to read from file
	msg2 := &simplepb.SimpleMessage{}

	readFromFile("simple.bin",msg2)
	fmt.Println(msg2)

	jsonMsg:=toJson(msg)
	fmt.Println(jsonMsg)
	msg3 := &simplepb.SimpleMessage{}
	JsonToProto(jsonMsg,msg3)
	fmt.Println(msg3)

	//enum operations
	doEnum()
	//complex proto operations
	doComplex()

}
func doComplex(){
	cm := complexpb.ComplexMessage{
		SingleMessage: &complexpb.DummyMessage{
			Id: 1,
			Name:"First",
		},
		MultipleMessages : []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id: 2,
				Name:"Second",
			},
			&complexpb.DummyMessage{
				Id: 2,
				Name:"third",
			},
		},
	}
	fmt.Println(cm)
}
func doEnum(){
	em := enumpb.EnumMessage{
		Id :34,
		DayOfWeek: enumpb.DayofWeek_MONDAY,
	}
	fmt.Println(em)
}
func JsonToProto(in string , pb proto.Message){
	err := jsonpb.UnmarshalString(in,pb)
	if err!=nil{
		log.Fatalln("Can't Unmarshal")
	}

}
func toJson(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out,err := marshaler.MarshalToString(pb)
	if (err != nil) {
		log.Fatalln("Can't Marshal To String")
		return ""
	}
	return out
}
func readFromFile(fname string,pb proto.Message) error{
	in,err := ioutil.ReadFile(fname)
	if (err!= nil){
		log.Fatalln("Something is not right")
		return err

	}
	err2 := proto.Unmarshal(in,pb)
	if (err2!=nil){
		log.Fatalln("can't unmarshal",err)
		return err
	}
	return nil

}
func writeToFile(fname string,pb proto.Message) error{
	out,err := proto.Marshal(pb)
	if (err!= nil){
		log.Fatalln("Can't serialize",err)
		return err

	}
	if err :=ioutil.WriteFile(fname,out,0644); err != nil {
		log.Fatalln("Can't write to file",err)
		return err
	}
	fmt.Println("written ")
	return nil
}


//generates a  sample protobuf message
func generateSimple() *simplepb.SimpleMessage {
	msg := simplepb.SimpleMessage{
		MessageCode:123,
		Message:"hello",
		List : [] int32{1,2,3},
	}
	fmt.Println(msg.GetMessageCode())
	fmt.Println(msg.GetMessage())
	fmt.Println(msg.GetList())

	return &msg

}