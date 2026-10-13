# Hacking on git-carbon

## Release process

### Licensing

Make sure the licensing follows the [REUSE specification][REUSE]. This can be
checked using the associated tool:

```console
$ uvx reuse lint
```

If it complains about missing copyright or licensing information, it can be
added with:

```console
$ uvx reuse annotate -c 'Mullvad VPN AB' -l 'GPL-3.0-only' -r .
```

[REUSE]: https://reuse.software/spec-3.3/#license-files.
