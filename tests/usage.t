  $ . "$TESTDIR/__init__.sh"


Running `git carbon` with no subcommand should show usage information.
  $ git carbon
  git-carbon is a tool that helps manage duplicated files across repositories.
  
  Like git-submodule but for individual files.
  
  Usage:
    git-carbon [command]
  
  Available Commands:
    add         Add a carbon-copy to the current repository.
    completion  Generate the autocompletion script for the specified shell
    help        Help about any command
    remove      Remove a carbon-copy from the current repository.
    update      Update carbon copies from their respective repository.
  
  Flags:
    -h, --help   help for git-carbon
  
  Use "git-carbon [command] --help" for more information about a command.


Running `git carbon add` with no args:
  $ git carbon add
  Error: accepts between 2 and 3 arg(s), received 0
  Usage:
    git-carbon add [--force] [--ref REF] REPOSITORY FILE [DESTINATION]
  
  Flags:
    -f, --force        Add file even if it already exist in the repository
    -h, --help         help for add
    -q, --quiet        Suppress output
    -r, --ref string   Ref of source repository
  
  [1]


Running `git carbon remove` with no args:
  $ git carbon remove
  Error: accepts 1 arg(s), received 0
  Usage:
    git-carbon remove FILE
  
  Flags:
    -f, --force   Remove the file even if there are local modifications
    -h, --help    help for remove
    -k, --keep    Remove the file from .gitcarbon but keep the local copy
  
  [1]


Running `git carbon update` with no args:
  $ git carbon update
  Error: no files specified, pass FILE... or use --all
  Usage:
    git-carbon update (FILE... | --all)
  
  Flags:
    -a, --all             Update all files git-carbon knows about.
        --branch string   Override the branch to update from.
    -h, --help            help for update
    -q, --quiet           Suppress output.
  
  [1]
