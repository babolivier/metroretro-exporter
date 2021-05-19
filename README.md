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

## Format

The exporter generates markdown following this format:

```markdown
Section
  * Note 1
    * Comment 1
    * Comment 2
  * Note 2
    * Comment 3
```

Each note and comment starts with initials.

For notes, this is always the initials of the note's author.

For comments, if the comment starts with `XX: [...]`, with `XX` being
any uppercase string, `XX` is used. Otherwise, the initials of the
comment's author are used.

This is so one attendee (or more) can use comments to minute the
discussion around a note without needing every attendee to repeat
themselves in a comment.

### Comment example

This comment from John Doe:

> Hello world

Will translate into:

> JD: Hello world

But this comment (still from John Doe):

> BA: Hello world

Will translate into:

> BA: Hello world
