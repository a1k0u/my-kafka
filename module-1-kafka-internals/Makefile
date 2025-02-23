step-0:
	docker compose up -d

step-1:
	docker exec module-1-kafka-internals-kafka-0-1 kafka-topics.sh \
		--create \
		--if-not-exists \
		--topic yandex-practicum \
		--partitions 3 \
		--replication-factor 2 \
		--bootstrap-server localhost:9092
