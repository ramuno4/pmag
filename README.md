# Pmag
Command Line Interface for mananging programming projects with different frameworks and languages

## Setup
#### Clone locally
`git clone https://github.com/Jon1105/pmag`
#### Config
Fill in `config.yaml` with the appropriate format. Example:
```yaml
languages:
  - 
    name: 'Go'
    acros: ['go', 'golang']
    path: 'C:/Users/username/Documents/Programming/Go/src/github.com/Jon1105'
    templatePath: 'templates/Go'
    initialCommand: 'go mod init github.com/Jon1105/{{projectName}}'

defaultEditorPath: 'C:/Users/username/AppData/Local/Programs/Microsoft VS Code/bin/code
```
*More info about inidividual configuration options in config.yaml*

#### Setup Github V3 api (optional)
* Go to [github settings](https://github.com/settings/tokens) and sign in
* Generate a new token and select "repo" to allow repository creation access
* Copy generated token into `config.yaml`'s ghKey variable

#### Add to path
* Run `go build` and add the created executable to your path so you can access it from any directory in the terminal
* If `%GOBIN%` is already in your path, you can simply run `go install` which will build the project and move the executable to `%GOBIN%`

#### Delete project (optional)
Thanks to the go:embed package, the created executable can be run standalone, so the source, text and configuration files are no longer needed on the machine running the tool.
Note: this does not include the templates for each language. These must be remain on the machine for proper functionality of the template feature
You can therefore remove the pmag project from your computer while keeping the executable installed, which will continue to function as expected

## Usage
### Commands
#### Create
Usage:
    pmag create language project name [flags...]  
    -r, -readme Toggle creation a README.md file  
    -v, -vcs    Toggle initialization a version control system (git, github)  
    -p          Toggle visibility of the github repository created (if created)  
#### Open
Usage:
    pmag open [language] [project name] 
#### Vcs
Usage:
    pmag vcs [project name]
#### Help
Usage:
    pmag help [command]

## Project Structure
### `./conf`
Package to manage reading and accessing data from `config.yaml`
### `./cmd`
Package for cobra package
### `./utilities`
Package for multiple use functions throughout the application
### `./vcs`
Package for managing the vcs command
#### `./vcs/git`
Package for managing creation of git vcs
#### `./vcs/github`
Package for managing creation of github vcs