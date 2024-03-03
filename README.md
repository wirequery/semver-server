# Introduction

Simple server for managing the versions of artifacts using Semantic Versioning

NOTE: This application is currently under development. Please don't rely on this tool yet.

# Installation

```bash
go install github.com/wirequery/semver-server/cmd/sem@latest
```

# Usage

## Server

To start a `sem` server:
```bash
sem serve
```
Optionally, you can provide the flag `--port <port>` to set the port the server listens to.

## REST Interface

Create a new project
```
POST /v1/projects
{
  "group": "<group>",
  "name": "<name>"
}

=>

{
  "major": 0,
  "minor": 1,
  "patch": 0,
}
```
Get version of project
```
GET /v1/projects/:group/:name

=>

{
  "major": <major>,
  "minor": <minor>,
  "patch": <patch>,
}
```
Increment major
```
POST /v1/projects/:group/:name/major

=>

{
  "major": <major>,
  "minor": <minor>,
  "patch": <patch>,
}
```
Increment minor
```
POST /v1/projects/:group/:name/minor

=>

{
  "major": <major>,
  "minor": <minor>,
  "patch": <patch>,
}
```
Increment patch
```
POST /v1/projects/:group/:name/patch

=>

{
  "major": <major>,
  "minor": <minor>,
  "patch": <patch>,
}
```

# Roadmap

- [X] Implement in-memory rudimental prototype
- [ ] Use a real database
- [ ] Error handling
- [ ] Create projects with key validation
- [ ] Cover by tests
- [ ] Integrate with WireQuery

# Recipes

## Using Git history to determine increment

Coming Soon...
