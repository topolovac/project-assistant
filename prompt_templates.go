package main

var DOCUMENTATION_SYSTEM_MESSAGE = `Developer-Centric Software Project Documentation Prompt
Introduction
Introduce the software project, highlighting its functionality, architecture, and primary use-cases.
Define specific technical terminology pertinent to the project.

Setting Up the Development Environment
Prerequisites: Detail all software, hardware, and platform requirements.
Environment Setup: Step-by-step guide to set up the local development environment, including IDE setup, necessary extensions, and configurations.
Dependencies: Describe how to install project dependencies.

Codebase Overview
Repository Access: Provide details about how to access the code repository.
Directory Structure: Detail the organization of the codebase, explaining the purpose of each directory and key files.
Architecture: Present a high-level overview of the software architecture, explaining components, layers, and their interactions.

Branching and Merging Strategies
Branching Strategy: Describe the branching model the project follows (e.g., Gitflow, GitHub flow).
Merging Strategy: Explain the process and guidelines for code merges and pull requests.

Coding Standards and Conventions
Detail coding and naming conventions to be followed.
Outline code organization and OOP principles to adhere to.
Include linting, formatting guidelines, and tools used.

Building the Project
Provide instructions on how to build the project locally.
Describe any build scripts or automation used in the project.

Testing
Testing Frameworks and Tools: Mention and describe how to set up and use testing tools.
Writing Tests: Provide guidelines for writing unit, integration, and system tests.
Running Tests: Explain how to run tests locally and interpret results.

Debugging
Describe strategies and tools for debugging.
Detail common issues and tips for troubleshooting them.

Deployment
Outline deployment procedures, including continuous integration/continuous deployment (CI/CD) pipelines.
Describe deployment environments and how to manage them.


API Documentation (If Applicable)
If the project exposes an API, detail API endpoints, request/response models, and usage examples.

Database Migrations and Management
Provide insights into database schema, migrations, and management.
If applicable, detail how to manage, backup, and restore database instances.

External Integrations
Explain any third-party services or integrations used and how to configure them (without exposing sensitive credentials).

Security Best Practices
Describe security practices for coding and any security tools used within the project.
Detail processes for reporting and handling security issues.

Performance Monitoring and Optimization
Detail how performance monitoring is handled.
Provide tips and tools for optimizing and profiling code.


Conclusion
Summarize critical points and any additional instructions for developers.
Provide resources for further learning and exploration.

Appendix
Include any additional documentation, resources, or scripts that developers might find useful.`