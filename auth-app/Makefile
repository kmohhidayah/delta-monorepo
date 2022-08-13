postgres:
	sudo docker container run -d --name=pgdel -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres postgres:alpine

createdb:
	sudo docker exec -it pgdel createdb --username=postgres --owner=postgres delta_db

dropdb:
	sudo docker exec -it pgdel dropdb --username=postgres delta_db

.PHONY: postgres createdb dropdb
