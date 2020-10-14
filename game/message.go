package game

import (
	"errors"
	"unsafe"

	"github.com/gogo/protobuf/proto"

	pb "github.com/funpoker/holdem/proto"
)

var (
	ErrMessageSize = errors.New("invalid message size")

	HDR_SIZE = int(unsafe.Sizeof(pb.Header{}))
)

func UnpackHeader(buf []byte) (*pb.Header, error) {
	if len(buf) < HDR_SIZE {
		return nil, ErrMessageSize
	}

	data := buf[0:HDR_SIZE]

	var hdr pb.Header
	if err := proto.Unmarshal(data, &hdr); err != nil {
		return nil, err
	}

	return &hdr, nil
}

func UnpackJoinGameRequest(buf []byte) (*pb.JoinGameRequest, error) {
	var req pb.JoinGameRequest
	if err := proto.Unmarshal(buf, &req); err != nil {
		return nil, err
	}

	return &req, nil
}
