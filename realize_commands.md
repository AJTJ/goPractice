https://github.com/oxequa/realize

# Realize basics

`realize start`
- It will create a .realize.yaml file if doesn't already exist, add the working directory as project and run your workflow.

## New Project?
- Grab `yaml` file for basicChi
- ensure it has all the commands
```
commands:
  build:
    status: true
  run:
    status: true
  generate:
    status: true
  install:
    status: true
```

## Running
- Once realize has started, it will rebuild and rerun the file every time it is saved.