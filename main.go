package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	proto "github.com/golang/protobuf/proto"
	"github.com/sorenvonsarvort/velas-sphere/internal/command"
	"github.com/sorenvonsarvort/velas-sphere/internal/entropy"
	"github.com/sorenvonsarvort/velas-sphere/internal/resources"
	"github.com/spf13/cobra"
)

func serializationDemo() {
	test := &resources.TaskExecutionRequest{
		Id:    "1",
		Input: "hello",
	}
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	newTest := &resources.TaskExecutionRequest{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	fmt.Println(newTest.GetId())
}

func entropyDemo() {
	data := []byte{}

	for i := 0; i < 256; i++ {
		data = append(data, byte(i))
	}

	fmt.Println(entropy.Shannon(data))
}

func entropyFromFileDemo(path string) {
	file, _ := os.Open(path)
	defer file.Close()
	fileBytes, _ := ioutil.ReadAll(file)

	fmt.Println(entropy.Shannon(fileBytes))
}

func main() {
	// contractDeploymentDemo()
	// return
	rootCmd := &cobra.Command{
		Use: "velas-sphere",
	}

	rootCmd.AddCommand(command.NewPluginCommand())
	rootCmd.AddCommand(command.NewNodeCommand())

	rootCmd.SilenceErrors = true
	rootCmd.SilenceUsage = true

	err := rootCmd.ExecuteContext(context.TODO())
	if err != nil {
		defer log.Fatal("execution failed: ", err)
	}
}
