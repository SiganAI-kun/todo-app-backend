# todo-app-backend

go run cmd/server/main.go

go fmt ./...

-- ローカルDB構築

docker ps

docker exec -it <container-id> bash

mysql -u root -p < initdb.d/0_database_tasks.sql
mysql -u root -p < initdb.d/1_insert_tasks.sql


