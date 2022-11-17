# Documentation
Some basic introduction
## MacOS Instalation
  1. Install Go from https://go.dev/dl/

  2. Download [this](https://go.dev/doc/gopath_code) repository and follow this instructions to set up your go workspace if you haven't done so yet. If you want to change your GOPATH to something else add GOPATH to your shell/bash/zsh initialization file .bash_profile, .bashrc or .zshrc.

    export GOPATH=/something-else

  3. Add `GOPATH/bin` directory to your `PATH` environment variable so you can run Go programs anywhere.

    export PATH=$PATH:$(go env GOPATH)/bin
  
  4. Make sure to re-source source .bash_profile your current session or simply open new terminal.

## Running the server
  1. Make a copy of `env.example.` and create your `.env` and fill out the variables.

  2. `cd` into the project folder and run:

    CompileDaemon -command="./lp-server"

