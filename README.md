# txt-lsp

txt-lsp is a toy project where I play around with [Language Server Protocol (LSP)](https://microsoft.github.io/language-server-protocol/).

## Development

### Code quality

We use [pre-commit](https://pre-commit.com/) to maintain the code quality of this project. Refer to [.pre-commit-config.yaml](./.pre-commit-config.yaml) for the list of linters that we are using. Refer to [this page](https://pre-commit.com/#install) to install pre-commit and the git hook script.

```
pre-commit install
```

### Interacting with LSP client

[.vimrc.lua](./.vimrc.lua) shows how to use the language server with the Neovim LSP client.

Use the following to access the client-side logs.

```
tail $HOME/.cache/nvim/lsp.log
```
