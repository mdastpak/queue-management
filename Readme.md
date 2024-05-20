# Queue Management

Queue Management is a CLI tool for managing RabbitMQ users and queues. This tool allows you to create, delete, and update users and queues in RabbitMQ.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
  - [Commands](#commands)
    - [User Commands](#user-commands)
    - [Queue Commands](#queue-commands)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/mdastpak/queue-management.git
    cd queue-management
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

3. Build the CLI tool:

    ```sh
    go build -o queue-management
    ```

## Usage

You can use the `queue-management` CLI tool to manage RabbitMQ users and queues.

### Commands

#### User Commands

- Create a new user:

    ```sh
    ./queue-management user create [username] [password] [tags...]
    ```

    Example:

    ```sh
    ./queue-management user create newuser newpassword administrator
    ```

- Delete a user:

    ```sh
    ./queue-management user delete [username]
    ```

    Example:

    ```sh
    ./queue-management user delete newuser
    ```

- Update user permissions:

    ```sh
    ./queue-management user update-permissions [username] [configure] [write] [read]
    ```

    Example:

    ```sh
    ./queue-management user update-permissions newuser ".*" ".*" ".*"
    ```

#### Queue Commands

- Create a new queue:

    ```sh
    ./queue-management queue create [queueName]
    ```

    Example:

    ```sh
    ./queue-management queue create newqueue
    ```

## Configuration

Create a `config.yaml` file in the root of the project with the following structure:

```yaml
rabbitmq:
  url: "amqp://guest:guest@localhost:5672/"
  management_url: "http://guest:guest@localhost:15672/api/"
  username: "guest"
  password: "guest"
