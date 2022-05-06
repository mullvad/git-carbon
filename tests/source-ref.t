  $ . "$TESTDIR/__init__.sh"

Init source repository:

  $ git init --quiet A
  $ echo "Hello, main!" > A/README.md
  $ git -C A add README.md
  $ git -C A commit --quiet --message "Initial commit"
  $ git -C A switch --quiet --create notmain
  $ echo "Hello, notmain!" > A/README.md
  $ git -C A commit --all --quiet --message "Update README.md"
  $ git -C A switch --quiet main

Init dest repository:

  $ git init --quiet B

Add carbon copy

  $ git -C B carbon add --ref refs/heads/notmain ../A README.md
  $ cat B/README.md
  Hello, notmain!
  $ cat B/.gitcarbon
  [carbon "README.md"]
  	sourceRepository = ../A
  	sourceRef = refs/heads/notmain
  $ git -C B status --short
  A  .gitcarbon
  A  README.md

Update file

  $ git -C A switch --quiet notmain
  $ echo "Hi, notmain!" > A/README.md
  $ git -C A commit --all --quiet --message "Update README.md"
  $ git -C A switch --quiet main
  $ git -C B carbon update README.md
  Updating README.md from ../A
  $ cat B/README.md
  Hi, notmain!
