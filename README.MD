# StrongBody API

### Requirements
- [Golang version >= 1.20](https://go.dev/doc/install)
- [GoFrame v2.x](https://goframe.org/en/docs/cli/install)
- PostgreSQL

### How to run
1. Clone this repository:
     ```shell
   git clone https://github.com/quannv/strongbody-api.git
   cd strongbody-api
   ```
2. Set the environment variables in a `.env` file:
   ```shell
   cp .env.example .env
   ```

3. Install the dependencies:
   ```shell
   go mod tidy
   ```

4. Run the application:
    ```shell
   go run main.go
   ```

## Project Directory Structure
Refer : https://goframe.org/en/docs/design/project-structure