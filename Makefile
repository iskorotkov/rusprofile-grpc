IMAGE = iskorotkov/rusprofile-grpc
TAG = v0.1.0

.PHONY: proto-generate build run test build-image push-image run-image

proto-generate:
	cd ./api && buf generate

build:
	go build ./...

run:
	go run ./cmd/rusprofile-grpc

test:
	go test ./...

build-image:
	docker build -t $(IMAGE):$(TAG) -f ./build/rusprofile-grpc.dockerfile .

push-image:
	docker push $(IMAGE):$(TAG)

run-image:
	docker run -it -p 8080:8080 -p 8888:8888 --rm $(IMAGE):$(TAG)
