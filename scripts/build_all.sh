cd backend
GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo
docker image build -t zenika/marcel-backend:dev .

cd ../auth-backend
GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo
docker image build -t zenika/marcel-auth:dev .

cd ../back-office
yarn build
docker image build -t zenika/marcel-backoffice:dev .