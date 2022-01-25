package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/protobuf/proto"
	dpb "github.com/golang/protobuf/protoc-gen-go/descriptor"
)

var (
	helpFlag    = flag.Bool("h", false, "display help output")
	verboseFlag = flag.Bool("v", false, "display verbose output")
	fileFlag    = flag.String("f", "", "file to inspect")
)

func init() {
	flag.Parse()
}

func main() {
	if *helpFlag {
		printUsage()
		//flag.PrintDefaults()
		os.Exit(0)
	}

	if *fileFlag == "" {
		printUsage()
		os.Exit(1)
	}

	fds := &dpb.FileDescriptorSet{}

	f, err := os.Open(*fileFlag)
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}
	defer f.Close()

	bb, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	if err = proto.Unmarshal(bb, fds); err != nil {
		log.Fatalf("unable to unmarshal file descriptor set: %v", err)
	}

	if fds == nil {
		log.Fatal("fds is nil after unmarshal")
	}

	var messageCount int
	var fileCount int

	for _, v := range fds.File {
		fmt.Printf("File: %v\n", v.GetName())
		fmt.Printf("\tPackage: %v\n", v.GetPackage())

		if *verboseFlag {
			fmt.Printf("\tSyntax: %v\n", v.GetSyntax())
		}

		fmt.Printf("\tMessages:\n")

		for _, m := range v.GetMessageType() {
			fmt.Printf("\t\t%v\n", m.GetName())

			if *verboseFlag {
				fmt.Printf("\t\t\tFields:\n")

				for _, f := range m.GetField() {
					fmt.Printf("\t\t\t\t%v (type: %v number: %v)\n", f.GetName(), f.GetType(), f.GetNumber())
				}
			}

			messageCount += 1
		}

		fileCount += 1

		fmt.Println()
	}

	fmt.Printf("Found %v files with %v messages\n", fileCount, messageCount)
}

func printUsage() {
	fmt.Println("Usage: ./inspect -f file [-v | -h]")
}
