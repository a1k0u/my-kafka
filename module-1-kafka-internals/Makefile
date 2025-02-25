step-1:
	docker compose up -d

step-2:
	docker exec module-1-kafka-internals-kafka-0-1 kafka-topics.sh \
		--create \
		--if-not-exists \
		--topic yandex-practicum \
		--partitions 3 \
		--replication-factor 2 \
		--bootstrap-server localhost:9092

step-3:
	docker exec module-1-kafka-internals-kafka-0-1 kafka-consumer-groups.sh \
		--all-groups \
		--delete \
		--bootstrap-server localhost:9092

	docker exec -u root module-1-kafka-internals-kafka-0-1 sh -c \
		"echo '$$(< offset.json)' > offset.json && kafka-delete-records.sh  --bootstrap-server localhost:9092 --offset-json-file offset.json"

	go run ./app/*.go \
		--bootstrap-server localhost:9094 \
		--topic yandex-practicum \
		--produce-interval 1s

cleanup:
	docker compose down
	docker volume ls -q | grep "kafka-volume" | xargs docker volume rm
