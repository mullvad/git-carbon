  $ . "$TESTDIR/__init__.sh"

Init source repository:

  $ git init --quiet A
  $ echo "Old stuff…" > A/README.md
  $ git -C A add README.md
  $ git -C A commit --quiet --message "Initial commit"

Init dest repository:

  $ git init --quiet B

Add carbon copy

  $ git -C B carbon add ../A README.md
  $ cat B/README.md
  Old stuff…
  $ git -C B commit --quiet --message "Added README"

Update source file

  $ echo "It worked!" > A/README.md
  $ git -C A add README.md
  $ git -C A commit --quiet --message "Update"

Update carbon copy

  $ git -C B carbon update README.md
  Updating README.md from ../A
  $ cat B/README.md
  It worked!
  $ git -C B status --short
  M  README.md
