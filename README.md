# lmp-exporter

A prometheus exporter that returns the metrics visible on the "website" of an LMP server.

## Usage

1. Clone repo
2. Install dependencies: `go mod tidy`
3. Build (`go build -o exporter main.go`)
4. Configure
    - See `config-example.json` for default values
    - Save changed version as `config.json` in working directory
5. Run (`./exporter`)
    - Or `./exporter -config ~/config.json` if the config is somewhere else.
6. (Optional) You can use the `example-systemd.service` file to create the service.
    - Make sure you edit the placeholder values to fit your setup!

### Queries

TODO

### Example response

```
TODO
```
