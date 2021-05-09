# gof

gof is a command-line utility to post RSS/Atom feeds to the fediverse.
It has first-class support for [Pleroma][Pleroma], and thus should
support Mastodon, too...

gof is for "go feediverse", "go fediverse", "go fedi", or really
whatever you want. gof is a port of [feediverse][feediverse] written in
Go.

gof supports multiple feeds and multiple accounts.

[Pleroma]: https://pleroma.social
[feediverse]: https://github.com/edsu/feediverse

## usage

Before you can start using gof, you'll need to configure it. An example
configuration can be found [here][config-blob]. You can also just copy
the example:

```
$ cd $REPO
$ cp config.yaml.example config.yaml
$ vim config.yaml # don't forget to edit it!
```

You'll need an access token as well. The configuration will be saved in
`$REPO/gof.yaml`:

```
$ go build
$ ./gof
```

Once gof is configured, you might want to add it to your crontab:

```
*/30 * * * * cd /path/to/$REPO; gof
```

[config-blob]: https://git.sr.ht/~mjorgensen/gof/tree/master/gof.yaml.example

## post format

You can specify how the message looks. The variables you have to work
with are `URL`, `Title`, and `Summary`. You don't have to use all
variables.

TODO: `Summary` isn't implemented yet. 

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

## resources

Discussion and patches are welcome and should be directed to my public
inbox for now: [~mjorgensen/public-inbox@lists.sr.ht][lists]. Please use
``--subject-prefix PATCH gof`` for clarity when sending patches.

Bugs, issues, planning, and tasks can all be found at the tracker: 
[~mjorgensen/gof][todo]

[lists]: https://lists.sr.ht/~mjorgensen/public-inbox
[todo]: https://todo.sr.ht/~mjorgensen/gof