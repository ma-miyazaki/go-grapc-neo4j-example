# Development

## Run docker container

```
docker-compose up -d
```

## Generate codes

```
docker-compose exec go protoc --proto_path ./proto --go_out=plugins=grpc:./pb/calc calc.proto
```

## Run server

```
docker-compose exec -d go go run server/main.go
```

## Run client

```
docker-compose exec go go run client/main.go
```

# Neo4j

```
MATCH (p) RETURN p
```
