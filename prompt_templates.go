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

var PROMPT_SM_CODE_UPDATE = `UPDATE CODEBASE
YOU WILL BE GIVEN A TASK TO UPDATE THE CODEBASE
YOU WILL BE GIVEN FOLLOWING INFORMATION:
TASK_NAME: <TASK_NAME>
TASK_DESCRIPTION: <TASK_DESCRIPTION>
ACCEPTANCE_CRITERIA: <ACCEPTANCE_CRITERIA>
CODEBASE: <CODEBASE>
THE OUTPUT SHOULD BE ONLY A JSON FILE WITH THE UPDATED CODEBASE
SET THE IS_UPDATED FLAG TO TRUE FOR ANY FILES THAT HAVE BEEN UPDATED
SET THE IS_NEW FLAG TO TRUE FOR ANY FILES OR DIRECTORIES THAT HAVE BEEN CREATED
`
