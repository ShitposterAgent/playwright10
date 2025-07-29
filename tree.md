# shitposteragent Project Structure (Finalized)

A robust, scalable structure for a cross-language social media automation CLI tool using Go (for orchestration, CLI, API) and Playwright/TypeScript (for browser automation), with AI integration (Ollama, etc.).

---

## Root Directory

```
shitposteragent/
├── gocli/                        # All Go code for CLI, API, orchestration
│   ├── main.go                   # Entry point for CLI tool
│   ├── go.mod                    # Go module file (if Go code is a module)
│   ├── config/                   # Config management
│   │   └── config.go
│   ├── api/                      # API server for integration
│   │   └── server.go
│   ├── tasks/                    # Task runner/manager
│   │   └── tasks.go
│   ├── ai/                       # AI/Ollama integration
│   │   └── ollama.go
│   ├── playwright/               # Orchestrates Playwright scripts
│   │   └── runner.go
│   ├── utils/                    # Utility functions (logging, paths, etc.)
│   │   └── utils.go
│   └── types/                    # Shared types and structs
│       └── types.go
│
├── automation/                   # All TypeScript automation modules/scripts
│   ├── common/                   # Shared Playwright helpers/utilities
│   │   └── playwright-helpers.ts
│   ├── twitter/                  # Twitter automation scripts (e.g., post.ts, login.ts)
│   ├── facebook/                 # Facebook automation scripts
│   ├── instagram/                # Instagram automation scripts
│   ├── threads/                  # Threads automation scripts
│   ├── linkedin/                 # LinkedIn automation scripts
│   └── types/                    # Shared TypeScript types/interfaces
│
├── tests/                        # Playwright test specs
│   └── example.spec.ts
├── tests-examples/               # Example Playwright test specs
│   └── demo-todo-app.spec.ts
│
├── .github/                      # GitHub workflows and CI
│   └── workflows/
│       └── playwright.yml
│
├── package.json                  # Node/TypeScript dependencies
├── playwright.config.ts          # Playwright config
├── pnpm-lock.yaml                # pnpm lockfile (if using pnpm)
├── go.mod                        # Go module file (if Go code is a module)
├── shitposteragent.json          # Sample config file (for reference)
├── README.md                     # Project documentation
├── tree.md                       # This project structure file
└── .gitignore                    # Git ignore file
```

---

## Directory/Component Descriptions

- **gocli/**: All Go code for CLI, API, config, orchestration, and integration logic.
  - `main.go`: CLI entry point.
  - `config/`: Handles config loading, saving, and validation.
  - `api/`: HTTP API server for integration with other tools/services.
  - `tasks/`: Task runner and automation logic.
  - `ai/`: AI/Ollama integration helpers.
  - `playwright/`: Orchestrates Playwright scripts from Go.
  - `utils/`: Utility functions (logging, paths, etc.).
  - `types/`: Shared Go types and structs.

- **automation/**: All TypeScript automation modules/scripts for browser automation.
  - `common/`: Shared Playwright helpers/utilities.
  - One subdirectory per social platform (twitter, facebook, etc.), each with its own scripts.
  - `types/`: Shared TypeScript types/interfaces.

- **tests/** and **tests-examples/**: Playwright test specs and example specs.

- **package.json**, **playwright.config.ts**, **pnpm-lock.yaml**: TypeScript/Playwright project files are in the root directory (not in a subfolder).

- **shitposteragent.json**: Sample config file (used as a template for actual configs in production).

- **tree.md**: This file, describing the recommended project structure.

- **README.md**: Project overview, setup, and usage instructions.

- **.gitignore**: Git ignore rules for the project.

- **.github/**: GitHub Actions and CI/CD workflows.

---

## Notes

- Actual config values are stored in the user's home directory under `~/.shitposteragent/shitposteragent.json` in production.
- The Go CLI orchestrates Playwright scripts by invoking Node.js/TypeScript code for each platform.
- AI integration (Ollama, etc.) is handled via Go, with the option to extend to other providers.
- The structure is modular and extensible for new platforms, features, and integrations.
