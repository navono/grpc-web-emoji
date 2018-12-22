package v1

import (
	"context"
	
	"github.com/navono/gRPC-Web-emoji/pkg/api/v1"
	emoji "gopkg.in/kyokomi/emoji.v1"
)

// server is used to implement the EmojiService interface
type emojiServer struct{}

func NewEmojiServiceServer() v1.EmojiServiceServer {
	return &emojiServer{}
}

// Emojize takes a input string via EmojizeRequest, replaces known keywords with
// actual emoji characters and returns it via a EmojizeReply instance.
func (s *emojiServer) Emojize(c context.Context, r *v1.EmojizeRequest) (*v1.EmojizeReply, error) {
	return &v1.EmojizeReply{EmojizedText: emoji.Sprint(r.Text)}, nil
}
