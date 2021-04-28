# Metroretro exporter

This is a simple Metroretro exporter for your archiving needs.

## Build

```
git clone https://github.com/babolivier/metroretro-exporter
cd metroretro-exporter
go build
```

## Use

See the [sample configuration file](/config.sample.yaml) for
configuration.

```
./metroretro-exporter -b [BOARD ID]
```

The configuration file will be looked up at `./config.yaml`.
Use `-c` to provide a different location.