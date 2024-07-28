#!/bin/bash

# Define the command to run your Go ls implementation
GO_LS_CMD="go run . ls"

# Function to pause and wait for user input
pause() {
    read -p "Press [Enter] to continue..."  # Wait for user input
}

# Function to clear the console
clear_console() {
    clear  # Clear the console screen
}

# Function to run and compare commands
run_and_compare() {
    local cmd="$1"  # Command to run
    local description="$2"  # Description of the test
    
    clear_console  # Clear the screen
    echo -e "$description"  # Print the test description
    echo -e "\n"  # Print a newline

    # Run the Go command
    echo "Running command: $GO_LS_CMD $cmd"
    echo -e "\n"
    $GO_LS_CMD $cmd  # Execute the Go command
    echo -e "\n"

    # Run the system `ls` command
    echo -e "System 'ls' command"
    echo -e "\n"
    ls $cmd  # Execute the system `ls` command

    echo -e "\n"
    echo "-----------------------------------------------------------------------------------------------------"
    echo -e "\n"

    # Pause to allow the user to review the output
    pause
}

# Run the tests one by one
run_and_compare "-l" "Testing 'go run . ls -l' vs 'ls -l'"
run_and_compare "" "Testing 'go run . ls' vs 'ls'"
run_and_compare "-a" "Testing 'go run . ls -a' vs 'ls -a'"
run_and_compare "-t" "Testing 'go run . ls -t' vs 'ls -t'"
run_and_compare "-la" "Testing 'go run . ls -la' vs 'ls -la'"
run_and_compare "-lt" "Testing 'go run . ls -lt' vs 'ls -lt'"
run_and_compare "-lR internal/" "Testing 'go run . ls -lR internal/' vs 'ls -lR internal/'"
run_and_compare "/usr/bin/" "Testing 'go run . ls /usr/bin/' vs 'ls /usr/bin/'"

echo "All tests completed."

# Il y a un test en plus pour go run . ls -l /usr/bin et ls -l /usr/bin/
