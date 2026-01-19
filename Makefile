build:
	@clear
	@date
	go build .
	@date

run:
	@clear
	@date
	go run cmd/main.go
	@date

TEST_URL ?= "http://localhost:8080/pizzas"
get-pizzas:
	@clear
	@date
	curl $(TEST_URL) | jq

FLAVOR ?= quatrino
test-post-pizza:
	@clear
	@date
	curl -X POST -d @tests/$(FLAVOR).json $(TEST_URL) | jq # pass -i to curl to see headers

post-pizzas:
	@clear
	@date
	curl -X POST -d @tests/margherita.json $(TEST_URL) | jq # pass -i to curl to see headers
	curl -X POST -d @tests/quatrino.json $(TEST_URL) | jq # pass -i to curl to see headers
	curl -X POST -d @tests/tuna_with_cheese.json $(TEST_URL) | jq # pass -i to curl to see headers
	curl -X POST -d @tests/tuscany.json $(TEST_URL) | jq # pass -i to curl to see headers

PIZZA_ID ?= 1
get-pizza:
	@clear
	@date
	curl "$(TEST_URL)/$(PIZZA_ID)"

delete-pizza:
	@clear
	@date
	curl -X DELETE "$(TEST_URL)/$(PIZZA_ID)" | jq

update-pizza:
	@clear
	@date
	curl -X PUT -d @tests/update.json "$(TEST_URL)/$(PIZZA_ID)"

negative-price:
	@clear
	@date
	curl -i -X POST -d @tests/negative_pizza.json "$(TEST_URL)"

post-review:
	@clear
	@date
	curl -s -X POST -d @tests/review.json "$(TEST_URL)/$(PIZZA_ID)/reviews"
	@echo ""
	@echo "-----------------"
	@echo ""
	curl -s -X POST -d @tests/review_error.json "$(TEST_URL)/$(PIZZA_ID)/reviews"
