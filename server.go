package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"web/gin-gonic/calculatorpb"
)

type server struct {
	calculatorpb.CalculatorServiceServer
}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Sum calling ...")
	res := &calculatorpb.SumResponse{
		Result: req.GetOnce() + req.GetSecond(),
	}

	return res, nil
}

func (*server) PND(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
}

func main() {
	// router := gin.Default()
	// router.GET("/user/:id", controllers.Get)
	// router.POST("/user/create", controllers.Create)
	// router.PUT("/user/update/:id", controllers.Update)
	// router.DELETE("/user/delete/:id", controllers.Delete)
	// router.Run(":8080")

	fmt.Println("Server running !")

	listen, err := net.Listen("tcp", "0.0.0.0:50001")
	if err != nil {
		log.Fatalf("Connect fail !")
	}

	s := grpc.NewServer()

	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	err = s.Serve(listen)

	if err != nil {
		log.Printf("Calculator Service connect fail !")
	}
}
