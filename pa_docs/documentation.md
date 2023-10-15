# project-assistant

Project assistant is an AI tool that assists developers by generating documentation, suggesting improvements, and more.

## Directory Structure

- `.gitignore`: Specifies files and directories to be ignored by Git.
- `README.md`: Documentation file providing an overview of the project.
- `config.go`: Contains functions related to loading and retrieving project configurations.
- `helpers.go`: Contains helper functions for working with directories and files.
- `main.go`: Entry point of the application and contains the main logic.
- `openai.go`: Contains functions for interacting with the OpenAI API.
- `prompt_templates.go`: Stores templates for prompting AI responses.
- `types.go`: Defines the data structures used in the project.

## Architecture

The software architecture consists of the following components and layers:

- **Application**: The main component that orchestrates the flow of the program. It interacts with the OpenAI component and performs file and directory operations.
- **OpenAI**: The component responsible for communicating with the OpenAI API. It handles authentication and sends requests for generating documentation.
- **File and Directory Operations**: These components handle reading and writing files, retrieving directory information, and ignoring certain files and directories based on the configuration.

## Codebase Overview

### Functions

- `getConfig()`: Reads the project configuration from a file and returns the parsed configuration object.
- `getFlags()`: Retrieves command-line flags for the OpenAI API key and directory name.
- `getDirectoryInfo()`: Recursively retrieves directory information, including files and subdirectories, given a root path.
- `shouldEntryBeIgnored()`: Determines whether a file or directory should be ignored based on the configuration settings.
- `createOutputDir()`: Creates the output directory specified in the configuration.
- `createMDFile()`: Creates a Markdown file with the given content at the specified path.
- `logStructAsJSON()`: Logs a struct as a JSON string.

### Methods

- `Init()`: Initializes the OpenAI component with the provided API key.
- `completionRequest()`: Sends a completion request to the OpenAI API for generating AI responses.
- `CreateDocumentationRequest()`: Generates documentation by sending a completion request to OpenAI, including the directory information as input.

## Coding Standards and Conventions

- Coding style should follow the standard Go formatting guidelines.
- Variable and function names should be descriptive and follow camelCase notation.
- Code should be organized into logical modules and packages.
- Object-oriented programming principles should be followed when designing code.

## Testing

### Testing Frameworks and Tools

Testing in Go is typically done using the built-in `testing` package and the `go test` command. To run tests locally, execute the following command in the project root directory:

```
go test ./...
```

### Writing Tests

Tests should be written using the `testing` package and follow the naming convention `func TestXxx(*testing.T)`. The `*testing.T` parameter is used to manage test state and report test failures. Tests should cover different scenarios and edge cases.

### Running Tests

To run tests, navigate to the project root directory and execute the command:

```
go test ./...
```

This will run all tests in the project.

## Debugging

