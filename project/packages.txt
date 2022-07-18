# Starting a Project
go mod init {PROJECT}

# Routes package
go get github.com/go-chi/chi/v5
go get github.com/go-chi/chi/middleware
go get github.com/go-chi/cors

# Rabbitmq Driver
go get github.com/rabbitmq/amqp091-go

# Mongo
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/mongo/options

# Email
go get github.com/vanng822/go-premailer/premailer
go get github.com/xhit/go-simple-mail/v2

# Install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

> Generating the file from .proto file {inside the path}
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative {logs}.proto

go get google.golang.org/grpc
go get google.golang.org/grpc/protobuf
