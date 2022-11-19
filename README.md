# LP-Jcordoba95
Hello! Here is the project I built to learn and practice GoLang. 

## Project Description
LP-Jcordoba95 is an API for a Web platform that provides a simple calculator functionality (addition, subtraction,
multiplication, division, square root, and a random string generation) where each functionality will have a separate cost per request.  

User's will have a starting credit / balance. Each request will be deducted from the user's
balance. If the user's balance isn't enough to cover the request cost, the request shall be denied.
## Installation
### MacOS Instalation
  1. Install Go from https://go.dev/dl/

  2. Download [this](https://go.dev/doc/gopath_code) repository and follow this instructions to set up your go workspace if you haven't done so yet. If you want to change your GOPATH to something else add GOPATH to your shell/bash/zsh initialization file .bash_profile, .bashrc or .zshrc.

    export GOPATH=/something-else

  3. Add `GOPATH/bin` directory to your `PATH` environment variable so you can run Go programs anywhere.

    export PATH=$PATH:$(go env GOPATH)/bin
  
  4. Make sure to re-source source .bash_profile your current session or simply open new terminal.

  5. Create a database to be used for this project. I used [this free service](https://api.elephantsql.com/) .
## Running the API
  1. Make a copy of `env.example.` and create your `.env` and fill out the variables.

  2. `cd` into the project folder and run:

    CompileDaemon -command="./lp-server"

