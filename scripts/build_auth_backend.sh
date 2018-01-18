cd auth-backend
GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo
docker image build -t zenika/marcel-auth:dev .
cd ..