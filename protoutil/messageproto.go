package protoutil

import (
	"time"

	"github.com/alanwgt/apateapi/models"
	"github.com/alanwgt/apateapi/protos"
)

// MessageModelToProto maps one ore more Message model into an array of
// Message that was created to communicate messages between users
func MessageModelToProto(us string, ms ...models.Message) []*protos.Message {
	var ps []*protos.Message

	for _, m := range ms {
		d := us == m.Sender.Username && m.CreatedAt.Add(time.Duration(5)*time.Minute).After(time.Now())
		p := &protos.Message{
			From:      m.Sender.Username,
			To:        m.Receiver.Username,
			MessageId: m.ID,
			Timestamp: m.CreatedAt.Unix(),
			Deletable: d,
		}
		ps = append(ps, p)
	}

	return ps
}
