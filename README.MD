# `my-ls-project`

`my-ls-project` is a custom implementation of the `ls` command in Go. This project allows listing files and directories with various sorting and display options similar to the Unix/Linux `ls` command.

## Features

- Lists files and directories in a specified directory.
- Shows hidden files when `-a` is enabled.
- Supports long format `-l` to display additional details about files.
- Sorts files by modification date with the `-t` option.
- Sorts files alphabetically in ascending or descending order based on the `-r` option.
- Supports multi-column display of files.
- Supports recursive mode `-R` to show contents of subdirectories.

## Installation

To use this project, clone the repository and compile the program:

```bash
git clone https://zone01normandie.org/git/gkopoin/my-ls-1.git
cd my-ls-project
```

## Usage

```bash
go run . ls [Options] [Path...]
```

### Options

- `-a` : Show all files, including hidden files (those starting with a dot).
- `-l` : Display file details in long format.
- `-t` : Sort files by modification date, with the most recent first.
- `-R` : Recursively list files in subdirectories.
- `-a` : Display the current directory (`.`) and the parent directory (`..`).

### Examples

1. **List files in the current directory**:

    ```bash
    go run . ls
    ```

2. **List files with long format details**:

    ```bash
    go run . ls -l
    ```

3. **Show all files, including hidden ones**:

    ```bash
    go run . ls -a
    ```

4. **List files sorted by modification date**:

    ```bash
    go run . ls -t
    ```

5. **List files recursively**:

    ```bash
    go run . ls -R
    ```

6. **Combine options**:

    ```bash
    go run . ls -alRrt
    ```

---