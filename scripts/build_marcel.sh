GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo ./cmd/marcel
docker image build -t zenika/marcel:dev .
