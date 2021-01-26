module github.com/suboat/go-tensorflow-serving

go 1.15

require (
	github.com/golang/protobuf v1.4.3
	github.com/tensorflow/tensorflow v2.4.0+incompatible
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
)

replace github.com/tensorflow/tensorflow v2.4.0+incompatible => github.com/suboat/go-tensorflow v1.0.0
