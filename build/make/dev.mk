export DOCKER_IP = $(shell which docker-machine > /dev/null 2>&1 && docker-machine ip $(DOCKER_MACHINE_NAME))

check-docker-env-vars:
	@echo "DOCKER_HOST = \"$$DOCKER_HOST\""; \
	echo "DOCKER_IP = \"$$DOCKER_IP\""; \
	if [ -z "$$DOCKER_HOST" ]; then \
		echo "DOCKER_HOST is not set. Check your docker-machine is running or do \"eval \$$(docker-machine env)\"?" 1>&2; \
		exit 1; \
	fi; \
	if [ -z "$$DOCKER_IP" ]; then \
		echo "DOCKER_IP is not set. Check your docker-machine is running or do \"eval \$$(docker-machine env)\"?" 1>&2; \
		exit 1; \
	fi; \
	if [ "$$DOCKER_IP" = "127.0.0.1" ]; then \
		echo "DOCKER_IP is set to a loopback address. Check your docker-machine is running or do \"eval \$$(docker-machine env)\"?" 1>&2; \
		exit 1; \
	fi

dev-up: check-docker-env-vars
	docker-compose up -d

dev-down: check-docker-env-vars
	docker-compose stop
	docker-compose rm -f

go-up: check-docker-env-vars
	docker-compose up -d --force-recreate --no-deps go_chatbot_lab

go-down: check-docker-env-vars
	docker-compose stop go_chatbot_lab

dev: go-down go-up

quicktest:
	go install -v
	go test ./... \
		| awk '$$1 != "?" { \
			gsub("github.com/bossjones/go-chatbot-lab/", "", $$2); \
			printf("%-7s %-40s %8s\n", $$1, $$2, $$3) \
		  }'

quicktest-verbose:
	@echo "*** Building ..."
	go install -v
	@echo
	@for pkg in $$(go list ./... | grep -v /vendor/); do \
		path=$$(echo $$pkg | sed 's#github.com/bossjones/go-chatbot-lab#.#'); \
		if ! ls $$path/*_test.go > /dev/null 2>&1; then \
			continue; \
		fi; \
		echo "**** Running: go test -v $$path..."; \
		go test -v --timeout 5s $$path | egrep --color=none -- '--- |^ok|^Run|^SUCCESS.*'; \
		echo; \
	done
