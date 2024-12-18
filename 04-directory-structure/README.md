# Standard Go Project Structure

cmd/: Contains the main application entry points.
    - example
        - cmd/cli/main.go (CLI application)
        - cmd/server/main.go (HTTP server)

internal/: Houses private application and library code.
pkg/: Stores library code that can be used by external applications.

api/: Holds API-related files, such as OpenAPI/Swagger specs.
    - example
        - api/v1/hello.go (API endpoint handler)

web/: Contains web application specific components.
    - example
        - web/controllers/hello.go (HTTP handler)

configs/: Stores configuration files.
test/: For additional external test apps and test data.
docs/: Holds design and user documents.
scripts/: Contains scripts for various build, install, analysis, etc. operations.
build/: Packaging and Continuous Integration files.
deployments/: System and container orchestration deployment configurations.

At the root level, you typically have:
main.go: The main entry point of the application.
go.mod: Defines the module's module path and dependency requirements.
README.md: Project documentation.
