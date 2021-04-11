# Copyright (c) 2021, The NLP Odyssey Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

FROM golang:1.16.3-alpine3.13 as Builder
WORKDIR /go/src/langdet
COPY . .

RUN go mod download && \
    CGO_ENABLED=0 go build -ldflags="-extldflags=-static" -o /go/bin/langdet cmd/langdet/main.go

FROM scratch

COPY --from=Builder /go/bin/langdet /langdet

ENTRYPOINT ["/langdet"]
CMD ["help"]
