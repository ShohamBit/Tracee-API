#!/bin/bash

# Check if protoc-gen-go is installed, if not install it
if ! command -v protoc-gen-go &> /dev/null
then
    echo "protoc-gen-go not found, installing..."
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
else
    echo "protoc-gen-go is already installed."
fi

# Check if protoc-gen-go-grpc is installed, if not install it
if ! command -v protoc-gen-go-grpc &> /dev/null
then
    echo "protoc-gen-go-grpc not found, installing..."
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
else
    echo "protoc-gen-go-grpc is already installed."
fi



# Directory containing .proto files
TRACEE_PROTOS="../../api/v1beta1/*.proto"
# Run protoc for all .proto files
    protoc \
        --proto_path=.
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-json_out=orig_name=true,paths=source_relative:. \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative $TRACEE_PROTOS

# Check if the protoc command succeeded
if [ $? -eq 0 ]; then
    echo "Protoc command completed successfully."
    notify-send "Protoc Command" "Protoc command completed successfully."
else
    echo "Protoc command failed."
    notify-send "Protoc Command" "Protoc command failed."
fi