# Redis CLI Tool in Go 🚀

## Overview 📖

This "redis-cli" project, developed in Go, was created as part of the ["Coding Challenge #37 - Redis CLI Tool"](https://open.substack.com/pub/codingchallenges/p/coding-challenge-37-redis-cli-tool) by John Crickett. It's designed for effective interaction with Redis servers, encompassing serialization and deserialization of RESP messages, establishing network connections, and offering a user-friendly interface for executing Redis commands.

## Features 🌟

- **RESP Serialization/Deserialization** 💾: Handles encoding and decoding of RESP messages.
- **Network Connection** 🌐: Establishes connections to Redis servers.
- **Command Line Interface** ⌨️: Offers an interactive CLI for executing Redis commands.
- **Flexible Server Connection** 🔗: Allows specifying host and port for the Redis server.
- **Extended Command Support** 🛠️: Includes a help command with detailed descriptions of Redis commands.

## Getting Started 🚦

### Prerequisites

- Go (version 1.21 or higher) 🐹
- Redis server 📦

### Installation 🛠️

1. Clone the repository:
   ```bash
   git clone https://github.com/afonsocraposo/redis-cli
   ```

2. Navigate to the project directory:
   ```bash
   cd redis-cli
   ```

3. Build the project:
   ```bash
   go build -o redis-cli cmd/redis-cli/redis-cli.go
   ```

### Usage 📚

#### Basic Execution
Run the CLI tool:
```bash
./redis-cli
```

#### Server Connection
Specify the host and port:
```bash
./redis-cli -h [hostname] -p [port]
```

#### Interactive Mode
Enter commands in the interactive mode:
```bash
localhost:6379> set key value
OK
localhost:6379> get key
value
```

#### Help Command
Access help for Redis commands:
```bash
localhost:6379> help set

SET key value [NX|XX] [GET] [EX seconds|PX milliseconds|EXAT unix-time-seconds|PXAT unix-time-milliseconds|KEEPTTL]
summary: Sets the string value of a key, ignoring its type. The key is created if it doesn't exist.
since: 1.0.0
group: string

```

## Testing 🧪

- The project includes a suite of test cases following the Test-Driven Development (TDD) approach.
- Run the tests to ensure the functionality of serialization, deserialization, and command execution.

## Contributing 🤝

Contributions to enhance the "redis-cli" tool are welcome. Please follow the standard fork, branch, and pull request workflow.

## License 📄

Distributed under the MIT License. See `LICENSE` for more information.
