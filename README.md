## Go DSA | behindthe.dev

This repository is a supplementary resource to articles posted on behindthe.dev covering DSA, they are here for readers to be able to view working code, as well as have an attempt at their own implementation.


### How It Works

This project uses a simple two-part system to guide your learning:

    📁 /examples

    This is your "answer key." It contains complete, production-ready, and well-tested implementations of each data structure. Use this code to study best practices, see idiomatic Go, or compare it against your own solution after you've finished a challenge.

    📁 /todo

    This is your hands-on coding challenge. For each data structure, you'll find a boilerplate file (_todo.go) with all the necessary structs and method signatures. The logic, however, is missing. The implementation is strictly up to you, you are encouraged to not use any 3rd party resources / AI, if you get stuck try whiteboarding it, if you can't overcome it peek into examples.


### Here’s how to get started on your first data structure:

    Pick a Challenge: Navigate to the internal/todo/ directory and choose a data structure you want to implement, for example, stack.

    Open the File: Inside that directory, open the stack_todo.go file. You will see the empty methods waiting for you.

    Write the Code: Implement the logic for each method (Push, Pop, Peek, etc.). Think about the edge cases, especially for empty structures.

    Test Your Solution: Once you've written your implementation, run the provided tests from your terminal to see if your code is correct. The tests are your guide and validator.

### How to Test Your Code

You'll run all commands from the root directory of the project.

#### Running Tests for a Specific Challenge

To check your work on the stack challenge, run:

```bash
go test ./internal/todo/stack
```

#### Running All Tests in the Project

To run every test in both the examples and todo directories at once, use:

```bash
go test
```
#### Checking Test Coverage

Test coverage shows you which lines of your code were actually run by your tests.

Step 1: Generate a Coverage Profile
This command runs all the tests and saves the results to a file named coverage.out.

```bash
go test ./... -coverprofile=coverage.out
```

Step 2: View the Interactive HTML Report
This command opens a report in your browser that visually highlights your code.

```bash
go tool cover -html=coverage.out
```
