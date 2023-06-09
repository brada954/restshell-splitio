# Restshell Split.io

Commands to be included with Restshell to access Split.io configurations

## Build instructions

```bash
cd rs
go build
```

The build produces a new restshell called rs.exe (or rs) which includes all the base functionality of restshell with the Split.io commands loaded. You can rename the directory to any name for your executable.

Run the following to see the help command

```bash
rs help
```
As an alternative a restshell fork can be created with references to the commands in this repository

## Tests

The rs\tests folder contains tests that verify the function of the restshell command to verify operation with updates to the shell package. The rs shell should be started from the tests folder to run the local .rsconfig file required by the tests.

Tests are run as follows:

```bash
cd rs
go build
cd tests
../rs run test
```

All assertions should pass when the script completes.

## Split.io Commands

HELP command lists all the commands available and the Split.io commands start with 'split'.

The 'ABOUT Split' command provides additional informtion about using the commands.
