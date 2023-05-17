package main

import flatbuffers "github.com/google/flatbuffers/go"

func main() {
	builder := flatbuffers.NewBuilder(1024)

	courseName := builder.CreateString("Tutorial")
	courseDescription := builder.CreateString("Tutorial de treinamento")

	rdgmartins.CourseStart(builder)
}
