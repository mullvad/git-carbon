  $ source "$TESTDIR/__init__.sh"

Init source repository:

  $ git init --quiet A
  $ echo "It worked!" > A/README.md
  $ echo "answer: 42" > A/.deepconf.yaml
  $ git -C A add README.md
  $ git -C A add .deepconf.yaml
  $ git -C A commit --quiet --message "Initial commit"

Init dest repository:

  $ git init --quiet B

Add carbon copy

  $ git -C B carbon add README.md ../A
  $ cat B/README.md
  It worked!
  $ cat B/.gitcarbon
  [cc "README.md"]
  	sourceRepository = ../A
  $ git -C B status --short
  A  .gitcarbon
  A  README.md

Add an other file from:

  $ git -C B carbon add .deepconf.yaml ../A
  $ cat B/.deepconf.yaml
  answer: 42
  $ cat B/.gitcarbon
  [cc "README.md"]
  	sourceRepository = ../A
  [cc ".deepconf.yaml"]
  	sourceRepository = ../A
