package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"web/gin-gonic/calculatorpb"
)

func main() {
	log.Println("client running !")

	cc, err := grpc.Dial("localhost:50001", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Connect fail !")
	}

	defer cc.Close()

	client := calculatorpb.NewCalculatorServiceClient(cc)

	callSum(client)
}

func callSum(c calculatorpb.CalculatorServiceClient) {
	res, err := c.Sum(context.Background(), &calculatorpb.SumRequest{
		Once:   10,
		Second: 20,
	})

	if err != nil {
		log.Print("Error: ", err)
	}

	log.Print("Result: ", res.GetResult())
}
