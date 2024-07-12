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
- There are a few places markes `IMPROVEMENTS` in the code where the code could be improved with indications on what should be done.
- Setup a CI/CD pipeline to run the tests and build the project.
- Add a Dockerfile to build a Docker image for the project.
