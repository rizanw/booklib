.PHONY: docker-up
docker-up:
	docker-compose up --build -d

.PHONY: docker-down
docker-down:
	docker-compose down

.PHONY: build
build:
	go build -v -o bin/booklib cmd/http/*.go

.PHONY: build
run:
	@echo " >> build booklib"
	@make build
	@echo " >> booklib built"
	@echo " >> executing booklib"
	@./bin/booklib
	@echo " >> booklib is running"
