function genProto {
  DOMAIN=$1 # DOMAIN=$1 为rental DOMAIN=$0 为auth
  PROTO_PATH=./${DOMAIN}/api
  GO_OUT_PATH=./${DOMAIN}/api-gen/v1

  mkdir -p $GO_OUT_PATH
  protoc -I $PROTO_PATH \
    --go_out $GO_OUT_PATH \
    --go_opt paths=source_relative \
    --go-grpc_out $GO_OUT_PATH\
    --go-grpc_opt paths=source_relative \
    --grpc-gateway_out $GO_OUT_PATH \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt grpc_api_configuration=$PROTO_PATH/${DOMAIN}.yaml \
    $PROTO_PATH/*.proto
}

genProto user
