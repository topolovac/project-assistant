FILE_PATH:./config.go
FILE_NAME:config.go
CONTENT:The provided code is a Go program that reads a configuration file in JSON format and retrieves the configuration settings. 
It utilizes command-line flags to specify the OpenAI API key and the directory to create documentation for. 
The program first defines the default configuration settings and then reads the configuration file using the specified directory and file name. 
If there is an error reading or unmarshalling the file, the program returns the default configuration settings. 
After retrieving the configuration settings, the program adds the configuration file itself to the list of files to ignore.

FILE_PATH:./helpers.go
FILE_NAME:helpers.go
CONTENT:The given code is a Go program that performs various operations related to directories and files.

The main functions defined in the code include:
1. `getDirectoryInfo`: Retrieves information about a directory and its contents recursively.
2. `shouldEntryBeIgnored`: Determines whether a directory entry (file or directory) should be ignored based on the configuration settings.
3. `createOutputDir`: Creates an output directory if it doesn't already exist.
4. `createMDFile`: Creates a Markdown file with the specified content at the given path.
5. `logStructAsJSON`: Logs a struct as JSON.
6. `updateDirectory`: Updates all files within a directory, creating new files and recursively updating subdirectories.
7. `updateOrCreateFile`: Updates or creates a file with the specified content at the given path.
8. `flattenDirectory`: Flattens a directory structure into a list of files.

There are also several supporting struct types defined in the code, such as `Config`, `Directory`, `File`, and `FlattenedDir`. These structs represent different aspects of the directory and file information.

Overall, the code provides functionality to retrieve directory information, manage files and directories, and perform operations such as updating and creating files.

FILE_PATH:./main.go
FILE_NAME:main.go
CONTENT:The code is a Go application that performs various tasks related to creating documentation, updating a codebase, and generating project summaries. 

The main function initializes the application and retrieves the configuration. It then calls the necessary functions to create documentation, update the codebase, and generate a project summary.

The CreateDocumentation function creates documentation by calling the OpenAI CreateDocumentationRequest method. It handles the error and saves the output as a markdown file.

The UpdateCodebase function updates the codebase by calling the OpenAI CreateCodeUpdateRequest method. It handles the error and updates the directory based on the response.

The createProjectSummary function creates a project summary by calling the OpenAI CreateFileSummaryRequest method for each file in the directory. It then creates a summary for the entire project and saves it as markdown files.

Overall, the code demonstrates how to use the OpenAI API to automate documentation, codebase updates, and project summary generation.

FILE_PATH:./openai.go
FILE_NAME:openai.go
CONTENT:This file contains a Go package for interacting with the OpenAI platform. It defines a struct called `OpenAI` which has methods for initializing the OpenAI client, making various API requests, and generating completion requests for different purposes.

The `OpenAI` struct has a field for storing the JWT token and a client object for interacting with the OpenAI API. The `Init` method is used to initialize the JWT and client.

The `completionRequest` method is used to make a chat completion request to the OpenAI API. It takes a list of chat completion messages as input and returns the generated completion as a string.

The `CreateDocumentationRequest` method is used to create a documentation request. It takes directory information as input, marshals it into JSON, and makes a completion request with a system message and the JSON content.

The `CreateCodeUpdateRequest` method is used to create a code update request. It takes a code update command as input, marshals the codebase into JSON, generates a prompt string, and makes a completion request.

The `CreateFileSummaryRequest` method is used to create a file summary request. It takes file content as input, makes a completion request with a system message and the file content.

The `CreateProjectSummaryRequest` method is similar to the `CreateFileSummaryRequest` method but is used for project summaries. It takes project content as input and makes a completion request with a system message and the project content.

FILE_PATH:./types.go
FILE_NAME:types.go
CONTENT:The provided file contains several struct definitions for a code directory structure and configuration settings. The main structs include:

1. Directory: Represents a directory in the codebase and contains fields such as name, a list of files, and a list of subdirectories.

2. File: Represents a file in the codebase and contains fields such as name, type, content, and update status.

3. IgnoreSettings: Contains lists of files, directories, and file patterns to be ignored in the codebase.

4. Config: Contains configuration settings for the codebase, including the root path, OpenAI key, output directory, and ignore settings.

5. Flags: Represents command-line flags, such as JWT token and directory name.

6. CodeUpdateCommand: Represents a code update command, including task type, name, description, acceptance criteria, and the codebase directory structure.

These structs provide the necessary data structures to manage and manipulate the code directory and configuration settings.

