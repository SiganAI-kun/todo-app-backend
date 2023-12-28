# todo-app-backend

go run cmd/server/main.go


-- ローカルDB構築
docker ps
docker exec -it <container-id> bash

mysql -u root -p < initdb.d/0_database.sql