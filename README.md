# Development

## Run docker container

```
docker-compose up -d
```

## Generate codes

```
docker-compose exec go protoc --proto_path=proto --go_out=pb/calc/ --go_opt=paths=source_relative --go-grpc_out=pb/calc/ --go-grpc_opt=paths=source_relative calc.proto
docker-compose exec go protoc --proto_path=proto --go_out=pb/employee/ --go_opt=paths=source_relative --go-grpc_out=pb/employee/ --go-grpc_opt=paths=source_relative employee.proto
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

http://localhost:7474/

```
MATCH (p) RETURN p
```
