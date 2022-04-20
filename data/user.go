package data

import (
	userpb "github.com/woowonjin/go-grpc-example/protos/v1/user"
)

var Users = []*userpb.UserMessage {
	{
		UserId: "1",
		Name: "Henry",
		PhoneNumber: "01012344123",
		Age: 22,
	},
	{
		UserId: "2",
		Name: "Michael",
		PhoneNumber: "010123442323",
		Age: 55,
	},
	{
		UserId: "3",
		Name: "Jessie",
		PhoneNumber: "01056344123",
		Age: 15,
	},
	{
		UserId: "4",
		Name: "Max",
		PhoneNumber: "01017654123",
		Age: 37,
	},
	{
		UserId: "5",
		Name: "Tony",
		PhoneNumber: "01012340987",
		Age: 25,
	},
}