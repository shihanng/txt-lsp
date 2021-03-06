local lspconfig = require("lspconfig")
local configs = require("lspconfig.configs")

if not configs.foo_lsp then
	configs.foo_lsp = {
		default_config = {
			cmd = { "./txt-lsp" },
			filetypes = { "text" },
			root_dir = function(fname)
				return lspconfig.util.find_git_ancestor(fname)
			end,
			settings = {},
		},
	}
end

lspconfig.foo_lsp.setup({})

vim.lsp.set_log_level("debug")
