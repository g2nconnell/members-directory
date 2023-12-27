package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/g2nconnell/members-directory/directoryService"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type directoryServiceServer struct {
	directoryService.UnimplementedDirectoryServiceServer
}

func (s directoryServiceServer) CreateMember(ctx context.Context, req *directoryService.CreateMemberRequest) (*directoryService.CreateMemberResponse, error) {
	return &directoryService.CreateMemberResponse{
		Id: &directoryService.MemberID{
			Id: 1,
		},
	}, nil
}

func (s directoryServiceServer) DeleteMember(ctx context.Context, req *directoryService.DeleteMemberRequest) (*directoryService.DeleteMemberResponse, error) {
	return &directoryService.DeleteMemberResponse{
		Id: &directoryService.MemberID{
			Id: 1,
		},
	}, nil
}
func (s directoryServiceServer) UpdateMember(ctx context.Context, req *directoryService.UpdateMemberRequest) (*directoryService.UpdateMemberResponse, error) {
	return &directoryService.UpdateMemberResponse{
		Id: &directoryService.MemberID{
			Id: 1,
		},
	}, nil
}
func (s directoryServiceServer) GetMember(ctx context.Context, req *directoryService.GetMemberRequest) (*directoryService.GetMemberResponse, error) {
	return &directoryService.GetMemberResponse{
		MemberWithId: &directoryService.MemberWithID{
			Id: &directoryService.MemberID{
				Id: 1,
			},
			Member: &directoryService.Member{
				FirstName:  "Glenn",
				LastName:   "Connell",
				MiddleName: "",
				Address: &directoryService.Address{
					Street1: "14100 Tamiami Trl E",
					Street2: "Lot 553",
					Town:    "Naples",
					State:   "FL",
					Zip:     "34114",
				},
				Email: "g2nconnell@gmail.com",
				Phone: &directoryService.PhoneNumber{
					Country: "01",
					Area:    "508",
					Phone:   "414-8144",
				},
				Photo: &directoryService.Photo{
					Pic: []byte(""),
				},
				Birthday: timestamppb.New(time.Date(1961, time.September, 9, 0, 30, 0, 0, time.UTC)),
			},
		},
	}, nil
}
func (s directoryServiceServer) ListMembers(ctx context.Context, req *directoryService.ListMembersRequest) (*directoryService.ListMembersResponse, error) {
	return &directoryService.ListMembersResponse{
		Members: nil,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("Cannot create listener: %s", err)
	}
	serviceRegistrar := grpc.NewServer()
	myServer := &directoryServiceServer{}

	directoryService.RegisterDirectoryServiceServer(serviceRegistrar, myServer)
	err = serviceRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
