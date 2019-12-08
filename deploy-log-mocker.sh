echo "Start build"
cd log-mocker
go build
cd ..
echo "Copy log-mocker to docker container"
docker cp ./log-mocker/log-mocker filebeat:/usr/local/bin
echo "Start log-mocker"
docker-compose exec -d filebeat /usr/local/bin/log-mocker
echo "Log Mocker started"

