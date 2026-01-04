build:
	@clear
	@date
	go build .
	@date

run:
	@clear
	@date
	go run .
	@date

TEST_URL ?= "http://localhost:8080/pizzas"
test-get-pizzas:
	@clear
	@date
	curl $(TEST_URL) | jq

test-post-pizza:
	@clear
	@date
	curl -X POST -d @tests/post_pizza.json $(TEST_URL) | jq # pass -i to curl to see headers

PIZZA_ID ?= 1
test-get-pizza:
	@clear
	@date
	curl "http://localhost:8080/pizza/$(PIZZA_ID)" | jq
