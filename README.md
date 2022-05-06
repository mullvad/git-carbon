# git-carbon: duplicate files between repositories

`git-carbon` is a tool that help manage duplicated files across reposistories.
One can think of it like `git-submodule` but for individual files.

## Getting started

### Add a file from an other repository

Add a file from a remote repository to your local worktree:

```
git carbon add ssh://git@github.com:myorg/sharedfiles.git .editorconfig
```

`git-carbon` will refuse to overwrite a file that already exist, unless you use
`--force`:

```
git carbon add --force ssh://git@github.com:myorg/sharedfiles.git .editorconfig
```

`git-carbon` automatically stages its changes, so the only thing left is to
commit:

```
git commit -m "Add shared .editorconfig"
```

### Update shared file

`git-carbon` remembers where it got the files from, so it can apply changes
easily. To get the new version of a shared file, you can use the `update`
subcommand:

```
git carbon update .editorconfig
```

You can also update all files `git-carbon` knows about by passing `--all` to
the subcommand:

```
git carbon update --all
```
