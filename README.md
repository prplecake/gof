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

The first time you use gof, it'll ask you about your Pleroma instance.
You'll need an access token as well. The configuration will be saved in
`~/.gof`, unless you specify otherwise:

```
gof
```

Once gof is configured, you might want to add it to your crontab:

```
*/30 * * * * /path/to/gof
```