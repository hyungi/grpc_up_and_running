package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	/*"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"*/
	wrapper "github.com/golang/protobuf/ptypes/wrappers"
	"io"
	"log"
	"net"

	pb "server/ecommerce"

	"strings"
)

//go:generate protoc -I ecommerce ecommerce/order_management.proto --go_out=./ecommerce --go-grpc_out=./ecommerce

const (
	port           = ":50051"
	orderBatchSize = 3
)

var orderMap = make(map[string]pb.Order)

type server struct {
	orderMap map[string]*pb.Order
	// UnimplementedOrderManagementServer must be embedded to have forward compatible implementations.
	pb.UnimplementedOrderManagementServer
}

// Simple RPC
func (s *server) AddOrder(ctx context.Context, orderReq *pb.Order) (*wrapper.StringValue, error) {
	log.Printf("Order Added. ID : %v", orderReq.Id)
	orderMap[orderReq.Id] = *orderReq
	return &wrapper.StringValue{Value: "Order Added: " + orderReq.Id}, nil
}

// Simple RPC
func (s *server) GetOrder(ctx context.Context, orderId *wrapper.StringValue) (*pb.Order, error) {
	ord, exists := orderMap[orderId.Value]
	if exists {
		return &ord, status.New(codes.OK, "").Err()
	}

	return nil, status.Errorf(codes.NotFound, "Order does not exist. : ", orderId)
}

// Server-side Streaming RPC
func (s *server) SearchOrders(searchQuery *wrappers.StringValue, stream pb.OrderManagement_SearchOrdersServer) error {
	for key, order := range orderMap {
		log.Printf("key: %s, order: %s", key, fmt.Sprintf("%v", &order))
		for _, itemStr := range order.Items {
			if strings.Contains(itemStr, searchQuery.Value) {
				// Send the matching orders in a stream
				err := stream.Send(&order)
				if err != nil {
					return fmt.Errorf("error sending message to stream : %v", err)
				}
				log.Print("Matching Order Found : " + key)
				break
			}
		}
	}
	return nil
}

// Client-side Streaming RPC
func (s *server) UpdateOrders(stream pb.OrderManagement_UpdateOrdersServer) error {
	IDs := make([]string, 0)
	ordersStr := "Updated Order IDs : "
	for {
		order, err := stream.Recv()
		if err == io.EOF {
			// Finished reading the order stream.
			ordersStr += strings.Join(IDs, ", ")
			return stream.SendAndClose(&wrapper.StringValue{Value: "Orders processed " + ordersStr})
		}

		if err != nil {
			return err
		}
		// Update order
		orderMap[order.Id] = *order

		log.Printf("Order ID : %s - %s", order.Id, "Updated")
		IDs = append(IDs, order.Id)
	}
}

// Bi-directional Streaming RPC
func (s *server) ProcessOrders(stream pb.OrderManagement_ProcessOrdersServer) error {
	batchMarker := 1
	combinedShipmentMap := make(map[string]pb.CombinedShipment)
	for {
		orderId, err := stream.Recv()
		log.Printf("Reading Proc order : %s", orderId)
		if err == io.EOF {
			// Client has sent all the messages
			// Send remaining shipments
			log.Printf("EOF : %s", orderId)
			for _, shipment := range combinedShipmentMap {
				if err := stream.Send(&shipment); err != nil {
					return err
				}
			}
			return nil
		}
		if err != nil {
			log.Println(err)
			return err
		}

		destination := orderMap[orderId.GetValue()].Destination
		shipment, found := combinedShipmentMap[destination]

		if found {
			ord := orderMap[orderId.GetValue()]
			shipment.OrdersList = append(shipment.OrdersList, &ord)
			combinedShipmentMap[destination] = shipment
		} else {
			comShip := pb.CombinedShipment{Id: "cmb - " + (orderMap[orderId.GetValue()].Destination), Status: "Processed!"}
			ord := orderMap[orderId.GetValue()]
			comShip.OrdersList = append(shipment.OrdersList, &ord)
			combinedShipmentMap[destination] = comShip
			log.Print(len(comShip.OrdersList), comShip.GetId())
		}

		if batchMarker == orderBatchSize {
			for _, comb := range combinedShipmentMap {
				log.Printf("Shipping : %v -> %v", comb.Id, len(comb.OrdersList))
				if err := stream.Send(&comb); err != nil {
					return err
				}
			}
			batchMarker = 1
			combinedShipmentMap = make(map[string]pb.CombinedShipment)
		} else {
			batchMarker++
		}
	}
}

func main() {
	initSampleData()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterOrderManagementServer(s, &server{})
	// Register reflection service on gRPC server. https://grpc.github.io/grpc/core/md_doc_server_reflection_tutorial.html
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func initSampleData() {
	orderMap["102"] = pb.Order{Id: "102", Items: []string{"Google Pixel 3A", "Mac Book Pro"}, Destination: "Mountain View, CA", Price: 1800.00}
	orderMap["103"] = pb.Order{Id: "103", Items: []string{"Apple Watch S4"}, Destination: "San Jose, CA", Price: 400.00}
	orderMap["104"] = pb.Order{Id: "104", Items: []string{"Google Home Mini", "Google Nest Hub"}, Destination: "Mountain View, CA", Price: 400.00}
	orderMap["105"] = pb.Order{Id: "105", Items: []string{"Amazon Echo"}, Destination: "San Jose, CA", Price: 30.00}
	orderMap["106"] = pb.Order{Id: "106", Items: []string{"Amazon Echo", "Apple iPhone XS"}, Destination: "Mountain View, CA", Price: 300.00}
	orderMap["107"] = pb.Order{Id: "107", Items: []string{"Lenovo ThinkPad", "Galaxy S21"}, Destination: "Mountain View, CA", Price: 1700.00}
	orderMap["108"] = pb.Order{Id: "108", Items: []string{"Dell XPS"}, Destination: "San Diego, CA", Price: 1000.00}
}
