package benchmark

import (
	"encoding/json"
	"os"
	"testing"

	"andrei/benchmark/protos"

	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/proto"
)

type Student struct {
	Name string
	Age  int64
}

func BenchmarkBsonMarshall(b *testing.B) {
	s := Student{
		Name: "Andrei",
		Age:  21,
	}
	for i := 0; i < b.N; i++ {
		bson.Marshal(s)
	}
}

func BenchmarkJsonMarshall(b *testing.B) {
	s := Student{
		Name: "Andrei",
		Age:  21,
	}
	for i := 0; i < b.N; i++ {
		json.Marshal(s)
	}
}

func BenchmarkProtoMarshall(b *testing.B) {
	s := &protos.Student{
		Name: "Andrei",
		Age:  21,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.Marshal(s)
	}
}

func BenchmarkJsonUnMarshall(b *testing.B) {
	data, _ := os.ReadFile("./andrei.json")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var decoded Student
		json.Unmarshal(data, &decoded)
	}
}

func BenchmarkBsonUnMarshall(b *testing.B) {
	data, _ := os.ReadFile("./andrei.bson")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var decoded Student
		bson.Unmarshal(data, &decoded)
	}
}

func BenchmarkProtoUnMarshall(b *testing.B) {
	data, _ := os.ReadFile("./andrei.protox")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var decoded protos.Student
		proto.Unmarshal(data, &decoded)
	}
}
