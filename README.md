# go-cms-gql

GraphQL application for managing contents. Written in Go with gqlgen.

## Core Features

- Basic authentication.
- Content management.
- Generate content by topic, title and read duration.
- Generate tags content.

## Tech Stack

- Go
- Gqlgen
- MongoDB
- Chi
- Github Actions
- Docker

## How to Use

1. Clone this repository.

2. Copy the configuration file.

```sh
cp .env.example .env
```

3. Fill the configuration inside `.env` file.

4. Generate the admin account.

```sh
go run helper/admin/generate.go
```

5. Run the application.

```sh
go run server.go
```

## Notes for Using with Docker

1. Make sure to set the `APP_MODE` in the `.env` into `production`.

2. Adjust the `MONGO_URI` to use `mongdb-service` as the host.

3. Fill the username and password in the `MONGO_INITDB_ROOT_USERNAME` and `MONGO_INITDB_ROOT_PASSWORD`.

4. Generate the admin account.

```sh
./init_admin.sh
```

5. Run the application.

```sh
docker compose up -d
```

6. Stop the application.

```sh
docker compose down
```

## Documentation

The application documentation is available [here](https://documenter.getpostman.com/view/5781191/2sAY4xC3Ch#1ad11b01-ee75-4619-81d3-8db4181334a1).

## Notes for Recommendation Features

1. In order to use recommendation features (tag and content generation), make sure to insert the OpenAI API key.
