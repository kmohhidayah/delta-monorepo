postgres:
	sudo docker container run -d --name=pgdel -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret postgres:alpine

createdb:
	sudo docker exec -it pgdel createdb --username=root --owner=root delta_db

dropdb:
	sudo docker exec -it pgdel dropdb --username=root delta_db

.PHONY: postgres createdb dropdb
