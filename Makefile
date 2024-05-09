test:
	go test ./...

serve server run: get-air
	air

get-air:
	@which air || go install github.com/cosmtrek/air@latest
