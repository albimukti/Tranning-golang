# Generate dengan protoc
protoc --go_out=plugins=grpc:./proto proto/helloworld.proto

# Generate dengan buf
buf generate