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
	go run ./app/*.go

cleanup:
	docker compose down
	docker volume ls -q | grep "kafka-volume" | xargs docker volume rm
