local lspconfig = require("lspconfig")
local configs = require("lspconfig.configs")

if not configs.foo_lsp then
	configs.foo_lsp = {
		default_config = {
			cmd = { "/home/neovim/lua-language-server/run.sh" },
			filetypes = { "txt" },
			root_dir = function(fname)
				return lspconfig.util.find_git_ancestor(fname)
			end,
			settings = {},
		},
	}
end

lspconfig.foo_lsp.setup({})

vim.lsp.set_log_level("debug")
