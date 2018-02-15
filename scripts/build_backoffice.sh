cd back-office
yarn build
docker image build -t zenika/marcel-backoffice:dev .
cd ..