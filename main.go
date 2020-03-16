package main

import (
	"context"
	"fmt"
	"log"

	proto "github.com/golang/protobuf/proto"
	"github.com/sorenvonsarvort/velas-sphere/internal/cli"
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

func main() {
	rootCmd := &cobra.Command{
		Use: "sphere",
	}

	rootCmd.AddCommand(cli.NewProviderCommand())
	rootCmd.AddCommand(cli.NewRequesterCommand())

	rootCmd.SilenceErrors = true
	rootCmd.SilenceUsage = true

	err := rootCmd.ExecuteContext(context.TODO())
	if err != nil {
		defer log.Fatal("execution failed: ", err)
	}
}
