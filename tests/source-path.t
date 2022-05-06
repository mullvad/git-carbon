  $ . "$TESTDIR/__init__.sh"

Init source repository:

  $ git init --quiet A
  $ echo "It worked!" > A/README.md
  $ echo "answer: 42" > A/.deepconf.yaml
  $ git -C A add README.md
  $ git -C A add .deepconf.yaml
  $ git -C A commit --quiet --message "Initial commit"

Init dest repository:

  $ git init --quiet B

Add carbon copy with different path

  $ git -C B carbon add ../A README.md LISEZMOI.md
  $ cat B/LISEZMOI.md
  It worked!
  $ cat B/.gitcarbon
  [carbon "LISEZMOI.md"]
  	sourceRepository = ../A
  	sourcePath = README.md
  $ git -C B status --short
  A  .gitcarbon
  A  LISEZMOI.md
  $ git -C B commit --quiet --message "Add LISEZMOI.md"

Update source file

  $ echo "It worked again!" > A/README.md
  $ git -C A add README.md
  $ git -C A commit --quiet --message "Update README.md"

Update carbon copy

  $ git -C B carbon update LISEZMOI.md
  Updating LISEZMOI.md from ../A
  $ cat B/LISEZMOI.md
  It worked again!
  $ git -C B status --short
  M  LISEZMOI.md
