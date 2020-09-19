install_dependencies:
	cd web && npm ci

test_frontend
	cd web && npm test

test_backend:
	cd api && go test ./... -coverprofile cover.out

build_backend:
	docker-compose build api

build_frontend:
	docker-compose build web

lint_frontend:
	cd web && npm run lint

lint_backend:
	cd api && golangci-lint run ./...