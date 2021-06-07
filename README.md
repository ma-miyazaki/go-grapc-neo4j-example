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

NOTE: コンテナ起動時に自動で起動するので、手動で起動することは基本ない。

```
docker-compose exec go air
```

## Run client

```
docker-compose exec go go run client/main.go
```

# Neo4j

http://localhost:7474/

## 全グラフの表示
```
MATCH (p) RETURN p
```

## 全グラフの削除
```
MATCH (n) DETACH DELETE n
```

## ユニーク制約

ユニーク制約をかけるとインデックスも作成される

```
CREATE CONSTRAINT ON (n:Person) ASSERT n.uuid IS UNIQUE
CREATE CONSTRAINT ON (n:Person) ASSERT n.email IS UNIQUE
```

## インデックスの作成
```
CREATE INDEX FOR (n:Person) ON (n.lastName, n.firstName)
```
