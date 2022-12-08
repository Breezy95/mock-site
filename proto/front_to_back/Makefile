#create variable for filepath
#easy compile and clean!

compile:
	protoc --go_out=.	--go_opt=paths=source_relative	--go-grpc_out=.	--go-grpc_opt=paths=source_relative front_to_back.proto



clean:
	rm -f *.pb.go