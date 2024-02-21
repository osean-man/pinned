# Pinner

A very simple tool to pin commands and store them into a local database. It's kind of like bash history, but 
organized with only things you put into the list. 

## Usage

```bash
Usage:
  pinner [flags]
  pinner [command]

Available Commands:
  add         Add a new pinned command
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        List all pinned commands
  remove      Remove a pinned command
  update      Update a pinned command

Flags:
  -h, --help   help for pinner

```

### Examples
```bash
# Add a new pinned command
pinner add "ls -la"

# Echo a command into the list
echo "ls -la" | pinner add

# List all pinned commands
pinner list

# Update a command
pinner update 

# Remove a pinned command
pinner remove 
```

Keep in mind this is just a little utility I made for myself. It's definitely not secure and upon selecting a command it
will execute it. I've also noticed that there are some issues if you save a command and then you want to execute that and 
tee it into a file. The first line will actually be the command you ran. I don't have the will to fix it because it's not
an issue for me. :)