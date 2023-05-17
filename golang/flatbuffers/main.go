package main

import (
	"fmt"

	"tutorial/rdgmartins"

	flatbuffers "github.com/google/flatbuffers/go"
)

func main() {

	builder := flatbuffers.NewBuilder(1024)

	courseName := builder.CreateString("Tutorial")
	courseDescription := builder.CreateString("Tutorial de treinamento")

	// serializando um objeto
	rdgmartins.CourseStart(builder)
	rdgmartins.CourseAddId(builder, 1)
	rdgmartins.CourseAddName(builder, courseName)
	rdgmartins.CourseAddDescription(builder, courseDescription)
	course := rdgmartins.CourseEnd(builder)
	builder.Finish(course)

	// deserializando os dados
	buf := builder.FinishedBytes()
	c := rdgmartins.GetRootAsCourse(buf, 0) // le o binário direto do buffer
	fmt.Println(c.Id())
	fmt.Println(string(c.Name()))
	fmt.Println(string(c.Description()))
}
