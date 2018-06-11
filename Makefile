.PHONY: test test-down reset-dev reset-test

test:
	docker-compose up -d

test-down:
	docker-compose down

reset-test:
	docker-compose down && \
	docker-compose up -d