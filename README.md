# KOATUU

## Development

Build the binary

```sh
make
```

Start postgres

```sh
docker-compose up -Vd postgres
```

Run sql migrations

```sh
migrate -source file://migrations -database postgres://postgres:password@127.0.0.1/koatuu\?sslmode=disable up
```

Run the web server

```sh
./bin/server
```

## Test

Start postgres

```sh
docker-compose up -Vd postgres
```

Run sql migrations

```sh
migrate -source file://migrations -database postgres://postgres:password@127.0.0.1/koatuu\?sslmode=disable up
```

Run tests

```sh
go test -v ./...
```

## License

Project released under the terms of the MIT [license](./LICENSE).