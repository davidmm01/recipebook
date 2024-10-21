validate:
	go run backend --command validate

serve:
	go run backend --command serve

stack:
	docker-compose build --no-cache
	docker-compose up
