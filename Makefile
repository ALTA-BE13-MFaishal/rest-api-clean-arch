test:
	go test ./features/user/... -coverprofile=cover.out && go tool cover -html=cover.out
	go test ./features/auth/... -coverprofile=cover.out && go tool cover -html=cover.out