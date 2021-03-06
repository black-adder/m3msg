// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// mockgen rules for generating mocks for exported interfaces (reflection mode).
//go:generate sh -c "mockgen -package=producer $PACKAGE/producer Message,Producer | mockclean -pkg $PACKAGE/producer -out $GOPATH/src/$PACKAGE/producer/producer_mock.go"
//go:generate sh -c "mockgen -package=consumer $PACKAGE/consumer Message | mockclean -pkg $PACKAGE/consumer -out $GOPATH/src/$PACKAGE/consumer/consumer_mock.go"
//go:generate sh -c "mockgen -package=proto $PACKAGE/protocol/proto Encoder,Decoder | mockclean -pkg $PACKAGE/protocol/proto -out $GOPATH/src/$PACKAGE/protocol/proto/proto_mock.go"
//go:generate sh -c "mockgen -package=topic $PACKAGE/topic Service | mockclean -pkg $PACKAGE/topic -out $GOPATH/src/$PACKAGE/topic/topic_mock.go"

// mockgen rules for generating mocks for unexported interfaces (file mode).
//go:generate sh -c "mockgen -package=writer -destination=$GOPATH/src/$PACKAGE/producer/writer/consumer_service_writer_mock.go -source=$GOPATH/src/$PACKAGE/producer/writer/consumer_service_writer.go"
//go:generate sh -c "mockgen -package=writer -destination=$GOPATH/src/$PACKAGE/producer/writer/router_mock.go -source=$GOPATH/src/$PACKAGE/producer/writer/router.go"
//go:generate sh -c "mockgen -package=writer -destination=$GOPATH/src/$PACKAGE/producer/writer/shard_writer_mock.go -source=$GOPATH/src/$PACKAGE/producer/writer/shard_writer.go"

package mocks
