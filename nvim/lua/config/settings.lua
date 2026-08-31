vim.opt.number = true         -- Show line numbers
vim.opt.relativenumber = true -- Relative line numbers
vim.opt.mouse = 'a'           -- Enable mouse support
vim.opt.ignorecase = true     -- Ignore case in search
vim.opt.smartcase = true      -- Override ignorecase if search has caps
vim.opt.clipboard = 'unnamedplus' -- Use system clipboard
vim.opt.tabstop = 4
vim.opt.shiftwidth = 4
vim.opt.softtabstop = 4
vim.opt.expandtab = true
    
-- --------------------------------------------
-- ------------------- LSP --------------------
-- --------------------------------------------

-- installing
require("mason").setup()

require("mason-lspconfig").setup({
    -- Provide a list of servers to automatically install 
    -- Use lspconfig names (e.g., "lua_ls", not "lua-language-server")
    ensure_installed = { 
        "lua_ls",  
        "pyright",
	    "gopls",
        "golangci-lint",
        "golangci-lint-langserver"
    },
})

-- enabling
vim.lsp.enable('gopls')
vim.lsp.enable('pyright')
vim.lsp.enable('lua_ls')
vim.lsp.enable('golangci-lint')
vim.lsp.enable('golangci-lint-langserver')

vim.lsp.config('golangci-lint-langserver', {
    filetypes = { 'go', 'gomod' },
    init_options = {
        command = { 
            "golangci-lint", 
            "run", 
            "--out-format=json", 
            "--issues-exit-code=1" 
        },
    }
})

vim.lsp.config('gopls', {
  settings = {
    gopls = {
      analyses = {
        unusedparams = false, -- Handled more comprehensively by golangci-lint
      },
      staticcheck = false,   -- Turn off to prevent duplicate staticcheck warnings
    },
  },  
})

-- --------------------------------------------
-- --------------- Treesitter -----------------
-- --------------------------------------------

-- require('nvim-treesitter').install { 'python', 'javascript', 'go', 'lua' }

local ts = require("nvim-treesitter")

local languages = { 
  "lua", "python", "javascript", "typescript", "go" 
}

ts.setup({})

ts.install(languages)

vim.api.nvim_create_autocmd("FileType", {
  pattern = languages,
  callback = function()
    vim.treesitter.start()
      
    vim.wo.foldmethod = "expr"
    vim.wo.foldexpr = "v:lua.vim.treesitter.foldexpr()"
    vim.wo.foldlevel = 99

    vim.bo.indentexpr = "v:lua.require('nvim-treesitter').indentexpr()"
  end,
})

-- --------------------------------------------
-- ------------------- cmp --------------------
-- --------------------------------------------

local cmp = require'cmp'

cmp.setup({
    snippet = {
      -- REQUIRED - you must specify a snippet engine
      expand = function(args)
        vim.fn["vsnip#anonymous"](args.body) -- For `vsnip` users.
        -- require('luasnip').lsp_expand(args.body) -- For `luasnip` users.
        -- require('snippy').expand_snippet(args.body) -- For `snippy` users.
        -- vim.fn["UltiSnips#Anon"](args.body) -- For `ultisnips` users.
        -- vim.snippet.expand(args.body) -- For native neovim snippets (Neovim v0.10+)

        -- For `mini.snippets` users:
        -- local insert = MiniSnippets.config.expand.insert or MiniSnippets.default_insert
        -- insert({ body = args.body }) -- Insert at cursor
        -- cmp.resubscribe({ "TextChangedI", "TextChangedP" })
        -- require("cmp.config").set_onetime({ sources = {} })
      end,
    },
    window = {
      -- completion = cmp.config.window.bordered(),
      -- documentation = cmp.config.window.bordered(),
    },
    mapping = cmp.mapping.preset.insert({
      ['<C-b>'] = cmp.mapping.scroll_docs(-4),
      ['<C-f>'] = cmp.mapping.scroll_docs(4),
      ['<C-Space>'] = cmp.mapping.complete(),
      ['<C-e>'] = cmp.mapping.abort(),
      ['<CR>'] = cmp.mapping.confirm({ select = true }), -- Accept currently selected item. Set `select` to `false` to only confirm explicitly selected items.
    }),
    sources = cmp.config.sources({
      { name = 'nvim_lsp' },
      { name = 'vsnip' }, -- For vsnip users.
      -- { name = 'luasnip' }, -- For luasnip users.
      -- { name = 'ultisnips' }, -- For ultisnips users.
      -- { name = 'snippy' }, -- For snippy users.
    }, {
      { name = 'buffer' },
    })
})

  -- To use git you need to install the plugin petertriho/cmp-git and uncomment lines below
  -- Set configuration for specific filetype.
  --[[ cmp.setup.filetype('gitcommit', {
    sources = cmp.config.sources({
      { name = 'git' },
    }, {
      { name = 'buffer' },
    })
 })
 require("cmp_git").setup() ]]--

  -- Use buffer source for `/` and `?` (if you enabled `native_menu`, this won't work anymore).
cmp.setup.cmdline({ '/', '?' }, {
    mapping = cmp.mapping.preset.cmdline(),
    sources = {
      { name = 'buffer' }
    }
})

  -- Use cmdline & path source for ':' (if you enabled `native_menu`, this won't work anymore).
cmp.setup.cmdline(':', {
    mapping = cmp.mapping.preset.cmdline(),
    sources = cmp.config.sources({
      { name = 'path' }
    }, {
      { name = 'cmdline' }
    }),
    matching = { disallow_symbol_nonprefix_matching = false }
})

-- Set up lspconfig.
local capabilities = require('cmp_nvim_lsp').default_capabilities()
-- Replace <YOUR_LSP_SERVER> with each lsp server you've enabled.
vim.lsp.config('gopls', {
    capabilities = capabilities
})

vim.lsp.config('pyright', {
    capabilities = capabilities
})

vim.lsp.config('lua_ls', {
    capabilities = capabilities
})
