# Elastic-Stack

## Setup

Clone

```sh
git clone https://github.com/comoyi/elastic-stack.git
```

Step into directory

```sh
cd elastic-stack
```

Copy .env.example to .env

```sh
cp .env.example .env
```

## Usage

Build

```sh
docker-compose build
```

Up

```sh
docker-compose up
```

## Log mocker
```sh
chmod +x deploy-log-mocker.sh
./deploy-log-mocker.sh
```

## Open Kibana
[http://127.0.0.1:5601](http://127.0.0.1:5601)

