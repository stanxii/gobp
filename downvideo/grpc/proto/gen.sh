#/bin/zsh
protoc -I ./ ./downvideo.proto --go_out=plugins=grpc:.