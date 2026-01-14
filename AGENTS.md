# Repository Guidelines

## Project Structure & Module Organization
- This repository is a collection of learning materials and small projects.
- Top-level folders:
  - `rust-compiler-course/`: course content plus the `rusty/` Rust project.
  - `rust-gui-framework-course/`: Markdown lectures and exercises.
  - `scala-kurs-se/`: German-language lecture notes and exercises.
  - `rust/`: currently empty placeholder.
- Rust source lives in `rust-compiler-course/rusty/src/`. Course notes are Markdown in `lectures/`, `appendices/`, `resources/`, and `exercises/` subfolders.

## Build, Test, and Development Commands
- There is no repo-wide build system; commands are scoped to subprojects.
- Rust compiler project:
  - `cd rust-compiler-course/rusty && cargo build` — build the Rusty compiler.
  - `cd rust-compiler-course/rusty && cargo run` — run the compiler binary.
  - `cd rust-compiler-course/rusty && cargo fmt` — format Rust sources with rustfmt.

## Coding Style & Naming Conventions
- Rust code follows standard Rust style (4-space indentation, `snake_case` functions/modules, `CamelCase` types).
- Directory names use `kebab-case` and numeric prefixes for sequence (e.g., `lecture-01`, `appendix-a-llvm`).
- Keep Markdown headings short and title-cased; prefer fenced code blocks for snippets.

## Testing Guidelines
- There is no dedicated test suite in the repository today.
- If you add tests to `rusty/`, use Rust’s built-in test framework and run `cargo test` from `rust-compiler-course/rusty`.
- Name test modules and files to mirror the source module they cover (e.g., `lexer` tests in `lexer.rs`).

## Commit & Pull Request Guidelines
- Recent commit messages are short, sentence-case summaries (e.g., `Add rust compiler course`).
- Keep commits focused on a single course or feature.
- Pull requests should describe the course area affected, link to relevant lecture/exercise files, and include screenshots only if visual materials change.

## Documentation & Content Tips
- Prefer editing Markdown in-place rather than adding new build tooling.
- When adding new lectures or exercises, follow the existing numbering scheme and update the nearest course README.
