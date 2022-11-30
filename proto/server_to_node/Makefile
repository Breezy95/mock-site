#create variable for filepath
#easy compile and clean!

compile:
	protoc --go_out=.	--go_opt=paths=source_relative	--go-grpc_out=.	--go-grpc_opt=paths=source_relative server_to_node.proto



clean:
	rm -f *.pb.go