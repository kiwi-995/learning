//! Rusty - A Rust-inspired programming language compiler
//! 
//! This is the main entry point for the Rusty compiler.
//! The compiler is structured into several modules:
//! 
//! - `lexer`: Tokenizes source code
//! - `parser`: Builds an Abstract Syntax Tree
//! - `typechecker`: Performs semantic analysis and type checking
//! - `interpreter`: Tree-walk interpreter
//! - `compiler`: Compiles to bytecode
//! - `vm`: Bytecode virtual machine

mod lexer;
mod parser;
mod ast;
mod typechecker;
mod interpreter;
mod compiler;
mod vm;

use std::env;
use std::fs;

fn main() {
    let args: Vec<String> = env::args().collect();
    
    if args.len() < 2 {
        println!("Rusty Compiler v0.1.0");
        println!("Usage: rusty <command> [options]");
        println!();
        println!("Commands:");
        println!("  run <file>       Run a Rusty program");
        println!("  compile <file>   Compile to bytecode");
        println!("  repl             Start interactive REPL");
        return;
    }
    
    match args[1].as_str() {
        "run" => {
            if args.len() < 3 {
                eprintln!("Error: missing file argument");
                return;
            }
            run_file(&args[2]);
        }
        "compile" => {
            if args.len() < 3 {
                eprintln!("Error: missing file argument");
                return;
            }
            compile_file(&args[2]);
        }
        "repl" => {
            run_repl();
        }
        _ => {
            eprintln!("Unknown command: {}", args[1]);
        }
    }
}

fn run_file(path: &str) {
    match fs::read_to_string(path) {
        Ok(source) => {
            println!("Running: {}", path);
            // TODO: Implement file execution
            // 1. Lex the source
            // 2. Parse to AST
            // 3. Type check
            // 4. Interpret or compile and run
            println!("Source length: {} bytes", source.len());
        }
        Err(e) => {
            eprintln!("Error reading file: {}", e);
        }
    }
}

fn compile_file(path: &str) {
    println!("Compiling: {}", path);
    // TODO: Implement compilation
}

fn run_repl() {
    println!("Rusty REPL v0.1.0");
    println!("Type 'exit' to quit");
    println!();
    
    // TODO: Implement REPL
}
