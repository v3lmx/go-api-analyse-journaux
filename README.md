# go-api-analyse-journaux

This project implements an API to analyse the content of journal files. It returns a CSV file with the maximum occurences of data for each time segment.

## Setup

### Prerequisites

- Go 1.22
- make

## Usage

To run the project :

```bash
make run
```

To run the tests :

```bash
make test
```

## TODO

The following improvements could be made:

- Tests for the `api` and `repository` packages. Tests are currently only done for the `core` package.
- 
