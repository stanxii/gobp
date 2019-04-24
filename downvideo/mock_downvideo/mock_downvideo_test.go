package mock_downvideo_test

import (
"context"
"fmt"
"testing"
"time"

"github.com/golang/mock/gomock"
"github.com/golang/protobuf/proto"
downvideo "downvideo/grpc/proto"
hwmock "downvideo/mock_downvideo"
)

// rpcMsg implements the gomock.Matcher interface
type rpcMsg struct {
	msg proto.Message
}

func (r *rpcMsg) Matches(msg interface{}) bool {
	m, ok := msg.(proto.Message)
	if !ok {
		return false
	}
	return proto.Equal(m, r.msg)
}

func (r *rpcMsg) String() string {
	return fmt.Sprintf("is %s", r.msg)
}

func TestDownloadVideo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDVClient := hwmock.NewMockDownVideoClient(ctrl)
	req := &downvideo.DVRequest{Name: "unit_test"}
	mockDVClient.EXPECT().DownloadVideo(
		gomock.Any(),
		&rpcMsg{msg: req},
	).Return(&downvideo.DVReply{Message: "Mocked Interface"}, nil)
	testDownloadVideo(t, mockDVClient)
}

func testDownloadVideo(t *testing.T, client downvideo.DownVideoClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.DownloadVideo(ctx, &downvideo.DVRequest{Name: "unit_test"})
	if err != nil || r.Message != "Mocked Interface" {
		t.Errorf("mocking failed")
	}
	t.Log("Reply : ", r.Message)
}
