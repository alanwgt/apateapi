package protoutil

import (
	"github.com/alanwgt/apateapi/models"
	"github.com/alanwgt/apateapi/protos"
)

// UserModelToProto maps one ore more User model into an array of
// UserModel that was created to communicate user objects between users
func UserModelToProto(us ...models.User) []*protos.UserModel {
	var ps []*protos.UserModel

	for _, u := range us {
		p := &protos.UserModel{
			Username: u.Username,
			PubK:     u.PubKey,
		}
		ps = append(ps, p)
	}

	return ps
}

// FriendRequestToProto maps one ore more FriendRequest model into an array of
// FriendRequest proto that was created to communicate objects between users
func FriendRequestToProto(fr ...models.FriendRequest) []*protos.FriendRequest {
	var frs []*protos.FriendRequest

	for _, f := range fr {
		p := &protos.FriendRequest{
			Username:  f.Requester.Username,
			Timestamp: f.CreatedAt.Unix(),
		}
		frs = append(frs, p)
	}

	return frs
}
