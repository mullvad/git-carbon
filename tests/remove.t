  $ . "$TESTDIR/__init__.sh"

Init source repository:

  $ git init --quiet source
  $ touch source/{foo,bar}
  $ git -C source add foo bar
  $ git -C source commit --quiet --message "Initial commit"

Init local repository:

  $ git init --quiet local
  $ cd local

Add carbon copy

  $ git carbon add --quiet ../source foo
  $ git carbon add --quiet ../source bar baz
  $ git commit --quiet --message 'Add carbon copy'
  $ ls -C
  baz  foo

Remove carbon copy

  $ git carbon remove foo
  $ ls -C
  baz
  $ git status --short
  M  .gitcarbon
  D  foo

Try to remove a carbon copy with local modifications

  $ git reset --hard --quiet
  $ echo "CHANGED" > foo
  $ git carbon remove foo
  error: the following file has local modifications:
  	foo
  (use -k to keep the file, or -f to force removal)
  [1]
  $ ls -C
  baz  foo
  $ git status --short
   M foo

Force removing with --force

  $ git carbon remove --force foo
  $ ls -C
  baz
  $ git status --short
  M  .gitcarbon
  D  foo

Keep file with --keep

  $ git reset --hard --quiet
  $ git carbon remove --keep foo
  $ ls -C
  baz  foo
  $ git status --short
  M  .gitcarbon

Remove a file which is not a carbon copy

  $ touch quux
  $ git add quux
  $ git commit --quiet --message "Add quux"
  $ git carbon remove quux
  error: the following file is not a carbon copy:
  	quux
  [1]

Remove baz, should also remove the .gitcarbon file

  $ git carbon remove baz
  $ ls -C
  foo  quux
  $ git status --short
  D  .gitcarbon
  D  baz
