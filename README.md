# gof

gof is a command-line utility to post RSS/Atom feeds to the fediverse.
It has first-class support for Pleroma, and thus should support Mastodon,
too...

gof is for "go feediverse", "go fediverse", "go fedi", or really
whatever you want. gof is a port of [feediverse][feediverse] written in
Go.

gof supports multiple feeds and multiple accounts.

[feediverse]: https://github.com/edsu/feediverse

## requirements

* Go 1.16

## installation

```
go install github.com/prplecake/gof@latest
```

## usage

Before you can start using gof, you'll need to configure it. An example
configuration can be found [here][config-blob]. You can also just copy
the example:

```
$ cd $REPO
$ cp gof.yaml.example gof.yaml
$ vim gof.yaml # don't forget to edit it!
```

You'll need an access token as well. You can get on from the [Fediverse
Instance Access Token Generator][fediverse-access-token].

[fediverse-access-token]:https://tools.splat.soy/fediverse-access-token/

Build the thing:

```
$ go build
```

Then you can use it:

```
$ ./gof
```

You could also specify the configuration file to use via the command
line:

```
$ ./gof -c /path/to/your/gof.yaml
```

This would allow you to place the executable (and configuration)
anywhere on your system. Once gof is configured, you might want to add it to
your crontab, or your other favorite task scheduler:

```
*/30 * * * * cd /path/to/$REPO; gof
```

[config-blob]:https://github.com/prplecake/gof/blob/master/gof.yaml.example

## post format

You can specify how the message looks. The variables you have to work
with are `URL`, `Title`, and `Summary`. You don't have to use all
variables.

An example template:

```yaml
template: '{{.Title}}: {{.URL}}'
```

If you want the message to include line breaks, use YAML's multiline
syntax:

```yaml
template: |-
  {{.Title}}

  {{.URL}}
```

### Pleroma instances

Formatted posts are also supported. You can choose from plaintext,
Markdown, HTML, or BBCode, as long as they’re supported by your Pleroma
instance.

```yaml
template: |-
  **{{.Title}}**

  > {{.Summary}}

  {{.URL}}
format: markdown
```

See configuration details [in the wiki][wiki-formatting].

[wiki-formatting]:https://github.com/prplecake/gof/wiki/Configuration#format
