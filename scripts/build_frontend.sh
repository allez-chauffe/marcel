cd frontend
yarn build
docker image build -t zenika/marcel-frontend:dev .
cd ..