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

### Credentials

The configuration file requires a couple of MetroRetro credentials.

To find them, open any MetroRetro board on your browser, then open
the Storage section of your browser's developer tools (in the Storage
tab on Firefox, in the Application tab on Chrome/Chromium), and copy
the contents of the cookies `metret.sess` and `metret.sess.sig` to
the indicated keys in your configuration file.