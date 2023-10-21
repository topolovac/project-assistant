package main

var PROMPT_SM_DOCUMENTATION = `Developer Centric Software Project Documentation Prompt
INTRODUCTION
Introduce the software project, highlighting its functionality, architecture, and primary use-cases.
Define specific technical terminology pertinent to the project.

DIRECTORY STRUCTURE
Detail the organization of the codebase, explaining the purpose of each directory and key files.

ARCHITECTURE
Present a high-level overview of the software architecture, explaining components, layers, and their interactions.

CODEBASE OVERVIEW
Extract functions, classes, structs, methods, etc.
Identify and extract documentation comments.
Identify and extract function/method names, parameters, return types, and any other relevant information.
Extract and process comments to retrieve relevant details about the purpose and usage of functions/methods.
Gather information on the overall structure, such as directory layout, module dependencies, etc.

CODING STANDARDS AND CONVENTIONS
Detail coding and naming conventions to be followed.
Outline code organization and OOP principles to adhere to.
Include linting, formatting guidelines, and tools used.

TESTING
Testing Frameworks and Tools: Mention and describe how to set up and use testing tools.
Writing Tests: Provide guidelines for writing unit, integration, and system tests.
Running Tests: Explain how to run tests locally and interpret results.

DEBUGGING
Describe strategies and tools for debugging.
Detail common issues and tips for troubleshooting them.

API DOCUMENTATION (if applicable)
If the project exposes an API, detail API endpoints, request/response models, and usage examples.

DATABASE MIGRATIONS AND MANAGEMENT
Provide insights into database schema, migrations, and management.
If applicable, detail how to manage, backup, and restore database instances.

SECURITY BEST PRACTICES
Describe security practices for coding and any security tools used within the project.
Detail processes for reporting and handling security issues.

PERFORMANCE MONITORING AND OPTIMIZATION
Detail how performance monitoring is handled.
Provide tips and tools for optimizing and profiling code.


CONCLUSION
Summarize critical points and any additional instructions for developers.
Provide resources for further learning and exploration.

APPENDIX
Include any additional documentation, resources, or scripts that developers might find useful.`

var PROMPT_SM_CODE_UPDATE = `UPDATE CODEBASE PROMPT
TASK:
- Name: %s
- Type: %s
- Description: %s
- Acceptance Criteria: %s

OBJECTIVE: Modify the "content" from CODEBASE_INPUT according to the provided task and generate a detailed JSON file reflecting the changes made.

OUTPUT REQUIREMENTS:
- Format: JSON file only.
- Properties to include in JSON:
  - "is_updated": Set to TRUE for files that have been updated.
  - "is_new": Set to TRUE for any files or directories that have been created.

GUIDELINES:
1. Ensure all updates to the codebase align with the provided task description and acceptance criteria.
2. Maintain a backup of the original codebase to prevent accidental data loss.
3. Verify that the output JSON file accurately represents the changes made to the codebase.


CODEBASE_INPUT:
%s
`

var PROMPT_FILE_SUMMARY = `FILE SUMMARY PROMPT
You will be provided content from a file.
Your task is to generate a summary of the file content.
`

var PROMPT_PROJECT_SUMMARY = `PROJECT SUMMARY PROMPT
You will be provided summary for each project file.
Your task is to generate a summary of the project.
`
