[![Go](https://github.com/prplecake/gof/actions/workflows/go.yml/badge.svg)](https://github.com/prplecake/gof/actions/workflows/go.yml)
[![GitHub release (latest SemVer including pre-releases)](https://img.shields.io/github/v/release/prplecake/gof?include_prereleases)](https://github.com/prplecake/gof/releases/latest)

# gof

gof is a command-line utility to post RSS/Atom feeds to the fediverse.

Confirmed working with:

* Mastodon and forks such as,
  * glitch-soc
  * hometown
* Pleroma

gof is for "go feediverse", "go fediverse", "go fedi", or really
whatever you want. gof started as a port of [feediverse][feediverse],
written in Go.

gof supports multiple feeds and multiple accounts.

[feediverse]: https://github.com/edsu/feediverse

## requirements

* Go 1.21

## installation

Download the latest release for your system from the
[Releases page](https://github.com/prplecake/gof/releases/latest).

### from source

Clone the repo and build the thing:

```shell
git clone https://github.com/prplecake/gof
cd gof && go build
```

## usage

Before you can start using gof, you'll need to configure it. An example
configuration can be found [here][config-blob]. You can also just copy
the example:

```shell
cp gof.example.yaml gof.yaml
vim gof.yaml # don't forget to edit it!
```

You'll need an access token as well. On Mastodon you can get some from
your settings page, and for others without a PAT UI, you can get on from
the [Fediverse Instance Access Token Generator][fediverse-access-token].

[fediverse-access-token]:https://tools.splat.soy/pleroma-access-token/

Then you can use it:

```shell
./gof
```

You could also specify the configuration file to use via the command
line:

```shell
./gof -c /path/to/your/gof.yaml
```

This would allow you to place the executable (and configuration)
anywhere on your system. Once gof is configured, you might want to add
it to your crontab, or your other favorite task scheduler:

```text
*/30 * * * * cd /path/to/$REPO; gof
```

[config-blob]:https://github.com/prplecake/gof/blob/master/gof.example.yaml

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

### Instances supporting formatted posts

Formatted posts are also supported. You can choose from plaintext,
Markdown, HTML, or BBCode, as long as they’re supported by your
instance. Here's an example with Markdown:

```yaml
template: |-
  **{{.Title}}**

  > {{.Summary}}

  {{.URL}}
format: markdown
```
