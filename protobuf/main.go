package main

import (
	"encoding/json"
	"fmt"

	"github.com/golang/protobuf/proto"
	pb "github.com/radar07/rest-go/grpc/protofiles"
)

func main() {
	p := &pb.Person{
		Id:    1,
		Name:  "Victor",
		Email: "vick@mail.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "999-888-7870", Type: pb.Person_HOME},
		},
	}

	p1 := &pb.Person{}
	body, _ := proto.Marshal(p)
	_ = proto.Unmarshal(body, p1)

	fmt.Println("original struct: ", p)
	fmt.Println("marshaled data: ", body)
	fmt.Println("unmarshaled data: ", p1)

	body, _ = json.Marshal(p)
	fmt.Println("Json marshaled data: ", string(body))
}
