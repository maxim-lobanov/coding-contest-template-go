# coding-contest-template-go

A Go template for competitive programming and coding contests. Provides a structured workflow for creating, testing, and solving algorithmic problems.

## Quick Start

1. Create a new task
    - `script/create-task A1` -> Creating task with name `tasks/A1` from `tasks/template`
    - `script/create-task A2 --from A1` - Creating task with name `tasks/A2` from `tasks/A1`
2. Fill input to `main.in`
3. Fill samples to `sample_1.in` and `sample_1.out`
    - You can add more samples if necessary: `sample_2.in / sample_2.out`, `sample_N.in / sample_N.out`, etc
5. Write code to `solution` method in `tasks/A1/main.go`
6. Run task
    - `script/run-task A1`

<img width="795" height="315" alt="image" src="https://github.com/user-attachments/assets/5717e701-8635-42cb-8ce3-06a8edb386b3" />

## Helper Packages

The template provides useful helper packages to simplify solution writing:

- **`internal/cast`** - Easy-to-use parsing methods for input handling:
  - `cast.ParseInt`, `cast.ParseFloat`, `cast.ParseIntArray`, etc.
  - `cast.ToString` for converting values to strings

- **`internal/linq`** - Generic methods for slice operations, similar to C# LINQ