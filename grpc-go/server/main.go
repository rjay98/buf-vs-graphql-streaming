package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	// This import path is based on the name declaration in the go.mod,
	// and the gen/proto/go output location in the buf.gen.yaml.
	seatsaverv1 "github.com/rjay98/buf-vs-graphql-streaming/grpc-go/seatsaver/gen/proto/go/v1"
	"google.golang.org/grpc"
)

// SeatSaverServiceServer implements the SeatSaverService API.
type seatSaverServiceServer struct {
	seatsaverv1.UnimplementedSeatSaverServiceServer
}

func (s *seatSaverServiceServer) Ping(ctx context.Context, req *seatsaverv1.PingRequest) (*seatsaverv1.PingResponse, error) {
	log.Println("Got ping server request")
	response := &seatsaverv1.PingResponse{}
	response.RuntimeInfo = strconv.FormatInt(time.Now().Unix(), 10)
	response.Message = "Ping successful"

	return response, nil
}

func (s *seatSaverServiceServer) PingStream(req *seatsaverv1.PingStreamRequest, srv seatsaverv1.SeatSaverService_PingStreamServer) error {
	log.Println("Got ping server request")
	response := &seatsaverv1.PingStreamResponse{
		RuntimeInfo: strconv.FormatInt(time.Now().Unix(), 10),
		Message:     "Ping successful",
	}
	srv.Send(response)
	return nil
}

func (s *seatSaverServiceServer) GetVenues(req *seatsaverv1.GetVenuesRequest, srv seatsaverv1.SeatSaverService_GetVenuesServer) error {
	response := &seatsaverv1.GetVenuesResponse{}
	srv.Send(response)
	return nil
}

func (s *seatSaverServiceServer) GetVenue(ctx context.Context, req *seatsaverv1.GetVenueRequest) (*seatsaverv1.GetVenueResponse, error) {
	response := &seatsaverv1.GetVenueResponse{}
	return response, nil
}

func (s *seatSaverServiceServer) GetOpenSeats(req *seatsaverv1.GetOpenSeatsRequest, srv seatsaverv1.SeatSaverService_GetOpenSeatsServer) error {
	response := &seatsaverv1.GetOpenSeatsResponse{}
	srv.Send(response)
	return nil
}

func (s *seatSaverServiceServer) GetSoldSeats(req *seatsaverv1.GetSoldSeatsRequest, srv seatsaverv1.SeatSaverService_GetSoldSeatsServer) error {
	response := &seatsaverv1.GetSoldSeatsResponse{}
	srv.Send(response)
	return nil
}

func (s *seatSaverServiceServer) GetReservedSeats(req *seatsaverv1.GetReservedSeatsRequest, srv seatsaverv1.SeatSaverService_GetReservedSeatsServer) error {
	response := &seatsaverv1.GetReservedSeatsResponse{}
	srv.Send(response)
	return nil
}

func (s *seatSaverServiceServer) GetSeats(req *seatsaverv1.GetSeatsRequest, srv seatsaverv1.SeatSaverService_GetSeatsServer) error {
	response := &seatsaverv1.GetSeatsResponse{}
	srv.Send(response)
	return nil
}

func (s *seatSaverServiceServer) ReserveSeat(ctx context.Context, req *seatsaverv1.ReserveSeatRequest) (*seatsaverv1.ReserveSeatResponse, error) {
	response := &seatsaverv1.ReserveSeatResponse{}
	return response, nil
}

func (s *seatSaverServiceServer) ReleaseSeat(ctx context.Context, req *seatsaverv1.ReleaseSeatRequest) (*seatsaverv1.ReleaseSeatResponse, error) {
	response := &seatsaverv1.ReleaseSeatResponse{}
	return response, nil
}

func (s *seatSaverServiceServer) BuySeat(ctx context.Context, req *seatsaverv1.BuySeatRequest) (*seatsaverv1.BuySeatResponse, error) {
	response := &seatsaverv1.BuySeatResponse{}
	return response, nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	listenOn := "127.0.0.1:8080"
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", listenOn, err)
	}

	server := grpc.NewServer()
	seatsaverv1.RegisterSeatSaverServiceServer(server, &seatSaverServiceServer{})
	log.Println("Listening on", listenOn)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}
