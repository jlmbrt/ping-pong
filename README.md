Very simple ping/pong server written in Go, created for debugging networks issues

Send an http request on `/ping` and it answers with `pong`. **It's magic ! 🪄🤩**

Multi-arch support: `amd64,arm64`

## Entrypoints

- `/` : return a Hello message with the server name
- `/ping` : return `pong`
- `/ready`: return an empty response with code 200

## Configuration

Use env variables for configuration:

- `ADDR`, the binding address (default `0.0.0.0`)
- `PORT`, the listening port (default `8080`)
- `NAME`, the server name (default random value)

## Docker

**Image**: `jlmbrt/ping-pong`

```shell
docker run -p 8080:8080 -d jlmbrt/ping-pong
```
