  $ . "$TESTDIR/__init__.sh"


Init source repository:

  $ git init --quiet A
  $ echo "READMEv1" > A/README.md
  $ echo "FOOBARv1" > A/FOOBAR.md
  $ git -C A add README.md FOOBAR.md
  $ git -C A commit -qm "Initial commit"

Init dest repository:

  $ git init --quiet B

Add carbon copies

  $ git -C B carbon add ../A README.md
  $ git -C B carbon add ../A FOOBAR.md
  $ git -C B commit -qm "Add carbon copies"

Update upstream files

  $ echo "READMEv2" > A/README.md
  $ echo "FOOBARv2" > A/FOOBAR.md
  $ git -C A commit -qam "Initial commit"

Update all copies

  $ git -C B carbon update --all
  Updating FOOBAR.md from ../A
  Updating README.md from ../A
  $ cat B/README.md
  READMEv2
  $ cat B/FOOBAR.md
  FOOBARv2
