include .env

db:
	docker run \
	--env MONGO_INITDB_ROOT_USERNAME=${MONGO_DB_USERNAME} \
	--env MONGO_INITDB_ROOT_PASSWORD=${MONGO_DB_PASSWORD} \
	-p "27017:27017/tcp" \
	--name mongodb -d mongo:4.4

clean:
	@docker stop mongodb
	@docker rm mongodb
