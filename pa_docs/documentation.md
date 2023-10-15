# project-assistant

Project assistant is an AI tool that assists developers by generating documentation, suggesting improvements, and more.

## Directory Structure

- `.gitignore`: Gitignore file specifying files and directories to ignore.
- `README.md`: Readme file containing project overview and description.
- `config.go`: Contains functions for reading and parsing configuration files.
- `helpers.go`: Contains helper functions for directory traversal, file manipulation, and logging.
- `main.go`: Entry point of the application, responsible for initializing the project and generating documentation.
- `openai.go`: Contains functions for interacting with the OpenAI API.
- `prompt_templates.go`: Contains the template for the documentation prompt.
- `types.go`: Defines the data structures used in the project.

## Architecture

The software architecture follows a simple structure. The main components are:

- `config.go`: Responsible for reading and parsing the project configuration file.
- `helpers.go`: Contains helper functions for directory traversal, file manipulation, and logging.
- `openai.go`: Handles interactions with the OpenAI API for generating documentation.
- `main.go`: Entry point of the application that initializes the necessary components and generates the documentation.

The components interact with each other to read the project's directory structure, generate documentation using the OpenAI API, and save the documentation to a specified output directory.

## Codebase Overview

### `config.go`

- `getConfig()`: Reads and parses the project configuration file, returning the configuration object.

### `helpers.go`

- `getDirectoryInfo(path string, config *Config) (Directory, error)`: Recursively reads the directory structure, skipping ignored files and directories, and returns the directory information.
- `shouldEntryBeIgnored(entry_name string, config *Config) bool`: Determines whether an entry (file or directory) should be ignored based on the configuration settings.
- `createOutputDir(config *Config) error`: Creates the output directory specified in the configuration.
- `createMDFile(content string, path string) error`: Creates a Markdown file with the specified content at the given path.
- `logStructAsJSON(s interface{})`: Logs the JSON representation of a structure.

### `main.go`

- `main()`: Entry point of the application. Gets the project configuration, reads the directory structure, initializes the OpenAI API client, generates and saves the documentation.

### `openai.go`

- `Init(jwt string)`: Initializes the OpenAI API client with the provided JWT (JSON Web Token).
- `completionRequest(messages []openai.ChatCompletionMessage) (string, error)`: Sends a completion request to the OpenAI API and returns the response content.
- `CreateDocumentationRequest(dir_info Directory) (string, error)`: Creates a documentation request by sending the directory information to the OpenAI API and returns the generated documentation.

### `prompt_templates.go`

- Contains a string constant `PROMPT_SM_DOCUMENTATION` that serves as the template for the documentation prompt.

### `types.go`

- Defines the data structures used in the project, including `Directory`, `File`, `IgnoreSettings`, and `Config`.
- Defines a global constant `CONFIG_DEFAULT_OUTPUT_DIR` for the default output directory.

## Coding Standards and Conventions

- The code follows Go's official coding conventions and style guide.
- Code organization follows the recommended package and file structure.
- OOP principles are not heavily used as the code is structured around functional programming principles.
- Linting and formatting guidelines are not explicitly mentioned in the provided code.

## Testing

Testing frameworks and tools are not implemented in the provided code.

## Debugging

Strategies and tools for debugging are not explicitly implemented in the provided code.

## API Documentation

No external APIs are utilized in the provided code.

## Database Migrations and Management

The project does not involve database operations.

## Security Best Practices

Security best practices are not explicitly handled or implemented in the provided code.

## Performance Monitoring and Optimization

Performance monitoring and optimization strategies are not implemented in the provided code.

## Conclusion

The provided code serves as a starting point for