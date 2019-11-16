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
`./gof.yaml`, unless you specify otherwise:

```
gof
```

Once gof is configured, you might want to add it to your crontab:

```
*/30 * * * * /path/to/gof
```

## resources

Additional documentation can be found in [gof's wiki][wiki].

Discussion and patches are welcome and should be directed to my public
inbox for now: [~mjorgensen/public-inbox@lists.sr.ht][lists]. Please use
``--subject-prefix PATCH gof`` for clarity when sending patches.

Bugs, issues, planning, and tasks can all be found at the tracker: 
[~mjorgensen/gof][todo]

[wiki]: https://man.sr.ht/~mjorgensen/gof
[lists]: https://lists.sr.ht/~mjorgensen/public-inbox
[todo]: https://todo.sr.ht/~mjorgensen/gof