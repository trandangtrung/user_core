
# Project Directory Structure
Refer : https://goframe.org/en/docs/design/project-structure

├── api : The input/output data structure definition for providing external services. Considering version management needs.
├── hack : Contains project development tools, scripts, etc. For example, configuration for CLI tools, and various shell/bat script files.
├── internal
│   ├── cmd
│   ├── config
│   ├── consts : Defines all constants for the project.
│   ├── controller : Entrance/interface layer for receiving and parsing user input parameters.
│   ├── entity
│   ├── initialize
│   ├── middleware
│   ├── repository
│   ├── router
│   ├── service : Management of business logic encapsulation, specific business logic implementation, and encapsulation.
│   └── storage
├── manifest : Contains files for program compilation, deployment, operation, and configuration.
├── resource : Static resource files. These files can be injected into release files in the form of resource packing/image compilation.
├── utility
├── .env.example
├── go.mod
├── go.mod
├── go.sum
├── main.go : Program entry file.
└── Makefile