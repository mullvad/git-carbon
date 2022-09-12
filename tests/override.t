  $ . "$TESTDIR/__init__.sh"

Init source repository:

  $ git init --quiet A
  $ echo "It worked!" > A/README.md
  $ git -C A add README.md
  $ git -C A commit --quiet --message "Initial commit"

Init dest repository:

  $ git init --quiet B
  $ echo "Local content" > B/README.md

Refuse to add a carbon copy if the file already exists

  $ git -C B carbon add --quiet ../A README.md
  Error: README.md already exists
  [1]
  $ cat B/README.md
  Local content

Adding anyway with --force

  $ git -C B carbon add --quiet --force ../A README.md
  $ cat B/README.md
  It worked!

Overwriting an existing carbon copy

  $ git init --quiet C
  $ echo "I am from C" > C/README.md
  $ git -C C add README.md
  $ git -C C commit --quiet --message "Initial commit"

  $ git -C B carbon add --quiet --force ../C README.md
  $ cat B/README.md
  I am from C
  $ cat B/.gitcarbon
  [carbon "README.md"]
  	sourceRepository = ../C

