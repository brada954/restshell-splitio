# Restshell Example

Restshell becomes more powerful when it is built with custom commands and functions for developers particular applications and environments. This example repository demonstrates how to build a custom version of restshell with third-party packages or custom commands and functions. Custom command packages can be created and shared with others.

## Getting Started

It is recommended to clone or fork this repository to build a custom restshell with your custom content. Your custom restshell can reference other custom content in other repositories.

Following this example will ensure custom shells can pickup new features and commands from restshell and other third-party providers.

## Build instructions

```bash
cd rs
go build
```

The build produces a new restshell called rs.exe (or rs) which includes all the base functionality of restshell with the example commands loaded. You can rename the directory to any name for your executable.

Run the following to see the help command

```bash
rs help
```

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

## Extending RestShell

RestShell is easily extendable by developing custom commands to perform typical operations. Custom commands can perform domain actions against a set REST Api's as well as perform operational tasks. Custom commands can peform any programmable task a user may desire. The goal of extensibility is to enable developers to simplify their day to day work and provide an eco-system of features.

Custom commands can provide high-level operation capabilities like below:

1. myapplogin
2. adduser --fn john --ln doe --email john.doe@example.com
3. finduser john

The example commands provide a demonstartion how to develop your own custom command and integrate them into restshell. The init.go file can be modified to include any custom content desrired as demonstarted.

There are projects developing custom commands for features like:

1. Communicating with a mongo db
2. Capturing JSON output of shell programs like kubectl or aws cli for processing in the tool
