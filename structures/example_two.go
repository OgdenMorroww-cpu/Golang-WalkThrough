package structures

import (
	"log"
	"os"
)

type Job struct {
	Command string
	Logger  *log.Logger
}

func ExampleTwo() {
	job := &Job{"Demo", log.New(os.Stdout, "Jobs: ", log.Ldate)}
	job.Logger.Print("test")
}
