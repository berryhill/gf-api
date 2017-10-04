sudo docker-compose -f docker-compose.test.yml up --build -d

go test ./... -tags=dev -v

sudo docker stop mongodb