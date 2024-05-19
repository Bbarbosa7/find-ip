# Find IP - CLI Application

## Overview

Find IP is a command-line application written in Go that allows users to search for IP addresses and server names associated with a given hostname. The application leverages the `urfave/cli` package to provide a straightforward interface for performing these lookups.

## Features

- **IP Lookup**: Search for IP addresses associated with a specified hostname.
- **Server Lookup**: Search for server names (NS records) associated with a specified hostname.

## Installation

To install and run the Find IP application, you need to have Go installed on your machine. You can download and install Go from the [official website](https://golang.org/dl/).

### Steps to Install

1. Clone the repository:

    ```sh
    git clone https://github.com/Bbarbosa7/find-ip.git
    cd find-ip
    ```

2. Build the application:

    ```sh
    go build -o find-ip
    ```

3. Run the application:

    ```sh
    ./find-ip
    ```

## Usage

The Find IP application provides two main commands: `ip` and `server`. Both commands accept a `--host` flag to specify the hostname you want to look up.

### IP Lookup

To search for IP addresses associated with a hostname:

```sh
./find-ip ip --host www.example.com
```

or using the alias:

```sh
./find-ip i --host www.example.com
```

### Server Lookup

To search for server names (NS records) associated with a hostname:

```sh
./find-ip server --host www.example.com
```

or using the alias:

```sh
./find-ip s --host www.example.com
```

## Project Structure

The project is organized as follows:

```
.
├── app
│   └── app.go      # Contains the main CLI application logic
├── find-ip         # Compiled binary after building the application
├── go.mod          # Go module file
├── go.sum          # Go dependencies file
└── main.go         # Entry point of the application
```

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request with your changes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For any inquiries or issues, please contact [brayann.barbosa@proton.me](mailto:brayann.barbosa@proton.me).

---

By following the instructions above, you should be able to set up and use the Find IP application to search for IP addresses and server names associated with any hostname. Happy searching!
