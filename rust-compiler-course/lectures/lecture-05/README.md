# Lecture 5: Building a Lexer in Rust

> *"The lexer is the foundation — get it right, and everything else flows smoothly."*

Time to write real code! In this lecture, we'll implement a complete hand-written lexer for Rusty using Rust's iterator patterns and pattern matching.

## 🎯 Learning Objectives

By the end of this lecture, you will be able to:
- Implement a hand-written lexer in Rust
- Use `Peekable` iterators for lookahead
- Handle string escapes and multi-character tokens
- Report errors with precise source locations
- Write comprehensive tests for the lexer

## 📖 Topics

### 5.1 Lexer Architecture

Our lexer will:
- Take a `&str` source code as input
- Produce a `Vec<Token>` as output
- Track line and column numbers
- Handle errors gracefully (collect and continue)

```rust
pub struct Lexer<'a> {
    source: &'a str,
    chars: Peekable<CharIndices<'a>>,
    current_pos: usize,
    line: usize,
    column: usize,
    tokens: Vec<Token>,
    errors: Vec<LexerError>,
}
```

### 5.2 The Token Structure

```rust
#[derive(Debug, Clone, PartialEq)]
pub struct Token {
    pub kind: TokenKind,
    pub span: Span,
    pub lexeme: String,  // Original text
}

#[derive(Debug, Clone, Copy, PartialEq)]
pub struct Span {
    pub start: usize,
    pub end: usize,
    pub line: usize,
    pub column: usize,
}

#[derive(Debug, Clone, PartialEq)]
pub enum TokenKind {
    // Literals
    Integer(i64),
    Float(f64),
    String(String),
    
    // Identifiers and keywords
    Identifier(String),
    
    // Keywords
    Let, Mut, Fn, Return, If, Else, While, Loop, For, In,
    Match, Struct, Enum, Impl, Break, Continue, True, False,
    SelfValue, SelfType,
    
    // Operators
    Plus, Minus, Star, Slash, Percent,
    EqEq, NotEq, Lt, Gt, LtEq, GtEq,
    AndAnd, OrOr, Not,
    Eq,
    PlusEq, MinusEq, StarEq, SlashEq,
    
    // Delimiters
    LParen, RParen, LBrace, RBrace, LBracket, RBracket,
    
    // Punctuation
    Comma, Colon, Semicolon, Dot, Arrow, FatArrow, DoubleColon, Ampersand,
    DotDot, DotDotEq,
    
    // Special
    Eof,
}
```

### 5.3 Lexer Implementation

```rust
use std::iter::Peekable;
use std::str::CharIndices;

impl<'a> Lexer<'a> {
    pub fn new(source: &'a str) -> Lexer<'a> {
        Lexer {
            source,
            chars: source.char_indices().peekable(),
            current_pos: 0,
            line: 1,
            column: 1,
            tokens: Vec::new(),
            errors: Vec::new(),
        }
    }
    
    /// Tokenize the entire source code
    pub fn tokenize(mut self) -> Result<Vec<Token>, Vec<LexerError>> {
        while !self.is_at_end() {
            self.scan_token();
        }
        
        // Add EOF token
        self.tokens.push(Token {
            kind: TokenKind::Eof,
            span: Span {
                start: self.source.len(),
                end: self.source.len(),
                line: self.line,
                column: self.column,
            },
            lexeme: String::new(),
        });
        
        if self.errors.is_empty() {
            Ok(self.tokens)
        } else {
            Err(self.errors)
        }
    }
    
    fn is_at_end(&mut self) -> bool {
        self.chars.peek().is_none()
    }
    
    /// Consume and return the next character
    fn advance(&mut self) -> Option<char> {
        if let Some((pos, ch)) = self.chars.next() {
            self.current_pos = pos + ch.len_utf8();
            if ch == '\n' {
                self.line += 1;
                self.column = 1;
            } else {
                self.column += 1;
            }
            Some(ch)
        } else {
            None
        }
    }
    
    /// Look at the next character without consuming it
    fn peek(&mut self) -> Option<char> {
        self.chars.peek().map(|(_, ch)| *ch)
    }
    
    /// Look at the character after next
    fn peek_next(&self) -> Option<char> {
        let mut iter = self.chars.clone();
        iter.next();
        iter.next().map(|(_, ch)| ch)
    }
    
    /// Consume the next character if it matches expected
    fn match_char(&mut self, expected: char) -> bool {
        if self.peek() == Some(expected) {
            self.advance();
            true
        } else {
            false
        }
    }
    
    fn scan_token(&mut self) {
        let start_pos = self.chars.peek().map(|(i, _)| *i).unwrap_or(self.source.len());
        let start_line = self.line;
        let start_column = self.column;
        
        let ch = match self.advance() {
            Some(c) => c,
            None => return,
        };
        
        let kind = match ch {
            // Whitespace
            ' ' | '\t' | '\r' | '\n' => return,  // Skip whitespace
            
            // Single-character tokens
            '(' => TokenKind::LParen,
            ')' => TokenKind::RParen,
            '{' => TokenKind::LBrace,
            '}' => TokenKind::RBrace,
            '[' => TokenKind::LBracket,
            ']' => TokenKind::RBracket,
            ',' => TokenKind::Comma,
            ';' => TokenKind::Semicolon,
            
            // Multi-character tokens
            '+' => {
                if self.match_char('=') { TokenKind::PlusEq }
                else { TokenKind::Plus }
            }
            '-' => {
                if self.match_char('>') { TokenKind::Arrow }
                else if self.match_char('=') { TokenKind::MinusEq }
                else { TokenKind::Minus }
            }
            '*' => {
                if self.match_char('=') { TokenKind::StarEq }
                else { TokenKind::Star }
            }
            '/' => {
                if self.match_char('/') {
                    // Line comment
                    self.skip_line_comment();
                    return;
                } else if self.match_char('*') {
                    // Block comment
                    self.skip_block_comment();
                    return;
                } else if self.match_char('=') {
                    TokenKind::SlashEq
                } else {
                    TokenKind::Slash
                }
            }
            '%' => TokenKind::Percent,
            '=' => {
                if self.match_char('=') { TokenKind::EqEq }
                else if self.match_char('>') { TokenKind::FatArrow }
                else { TokenKind::Eq }
            }
            '!' => {
                if self.match_char('=') { TokenKind::NotEq }
                else { TokenKind::Not }
            }
            '<' => {
                if self.match_char('=') { TokenKind::LtEq }
                else { TokenKind::Lt }
            }
            '>' => {
                if self.match_char('=') { TokenKind::GtEq }
                else { TokenKind::Gt }
            }
            '&' => {
                if self.match_char('&') { TokenKind::AndAnd }
                else { TokenKind::Ampersand }
            }
            '|' => {
                if self.match_char('|') { TokenKind::OrOr }
                else {
                    self.error(start_line, start_column, "Unexpected character '|'");
                    return;
                }
            }
            ':' => {
                if self.match_char(':') { TokenKind::DoubleColon }
                else { TokenKind::Colon }
            }
            '.' => {
                if self.match_char('.') {
                    if self.match_char('=') { TokenKind::DotDotEq }
                    else { TokenKind::DotDot }
                } else {
                    TokenKind::Dot
                }
            }
            
            // String literals
            '"' => return self.scan_string(start_pos, start_line, start_column),
            
            // Character literals
            '\'' => return self.scan_char(start_pos, start_line, start_column),
            
            // Numbers
            '0'..='9' => return self.scan_number(ch, start_pos, start_line, start_column),
            
            // Identifiers and keywords
            'a'..='z' | 'A'..='Z' | '_' => {
                return self.scan_identifier(ch, start_pos, start_line, start_column);
            }
            
            _ => {
                self.error(start_line, start_column, 
                    &format!("Unexpected character '{}'", ch));
                return;
            }
        };
        
        let end_pos = self.current_pos;
        let lexeme = self.source[start_pos..end_pos].to_string();
        
        self.tokens.push(Token {
            kind,
            span: Span {
                start: start_pos,
                end: end_pos,
                line: start_line,
                column: start_column,
            },
            lexeme,
        });
    }
    
    fn skip_line_comment(&mut self) {
        while let Some(ch) = self.peek() {
            if ch == '\n' {
                break;
            }
            self.advance();
        }
    }
    
    fn skip_block_comment(&mut self) {
        let mut depth = 1;  // Support nested comments
        
        while depth > 0 {
            match (self.advance(), self.peek()) {
                (Some('/'), Some('*')) => {
                    self.advance();
                    depth += 1;
                }
                (Some('*'), Some('/')) => {
                    self.advance();
                    depth -= 1;
                }
                (None, _) => {
                    self.error(self.line, self.column, "Unterminated block comment");
                    return;
                }
                _ => {}
            }
        }
    }
    
    fn scan_string(&mut self, start_pos: usize, start_line: usize, start_column: usize) {
        let mut value = String::new();
        
        loop {
            match self.advance() {
                Some('"') => break,
                Some('\\') => {
                    match self.advance() {
                        Some('n') => value.push('\n'),
                        Some('r') => value.push('\r'),
                        Some('t') => value.push('\t'),
                        Some('\\') => value.push('\\'),
                        Some('"') => value.push('"'),
                        Some('0') => value.push('\0'),
                        Some(ch) => {
                            self.error(self.line, self.column,
                                &format!("Invalid escape sequence '\\{}'", ch));
                        }
                        None => {
                            self.error(start_line, start_column, "Unterminated string");
                            return;
                        }
                    }
                }
                Some('\n') => {
                    self.error(start_line, start_column, "Unterminated string (newline in string)");
                    return;
                }
                Some(ch) => value.push(ch),
                None => {
                    self.error(start_line, start_column, "Unterminated string");
                    return;
                }
            }
        }
        
        let end_pos = self.current_pos;
        let lexeme = self.source[start_pos..end_pos].to_string();
        
        self.tokens.push(Token {
            kind: TokenKind::String(value),
            span: Span {
                start: start_pos,
                end: end_pos,
                line: start_line,
                column: start_column,
            },
            lexeme,
        });
    }
    
    fn scan_char(&mut self, start_pos: usize, start_line: usize, start_column: usize) {
        // TODO: Implement character literal scanning
        // Similar to string but expects exactly one character
    }
    
    fn scan_number(&mut self, first: char, start_pos: usize, start_line: usize, start_column: usize) {
        let mut is_float = false;
        
        // Check for hex/binary/octal
        if first == '0' {
            match self.peek() {
                Some('x') | Some('X') => {
                    self.advance();
                    return self.scan_hex(start_pos, start_line, start_column);
                }
                Some('b') | Some('B') => {
                    self.advance();
                    return self.scan_binary(start_pos, start_line, start_column);
                }
                Some('o') | Some('O') => {
                    self.advance();
                    return self.scan_octal(start_pos, start_line, start_column);
                }
                _ => {}
            }
        }
        
        // Consume integer part
        while let Some(ch) = self.peek() {
            if ch.is_ascii_digit() || ch == '_' {
                self.advance();
            } else {
                break;
            }
        }
        
        // Check for decimal point (but not range operator ..)
        if self.peek() == Some('.') && self.peek_next() != Some('.') {
            if let Some(next) = self.peek_next() {
                if next.is_ascii_digit() {
                    self.advance(); // consume '.'
                    is_float = true;
                    
                    // Consume fractional part
                    while let Some(ch) = self.peek() {
                        if ch.is_ascii_digit() || ch == '_' {
                            self.advance();
                        } else {
                            break;
                        }
                    }
                }
            }
        }
        
        // Check for exponent
        if let Some('e') | Some('E') = self.peek() {
            is_float = true;
            self.advance();
            
            if let Some('+') | Some('-') = self.peek() {
                self.advance();
            }
            
            while let Some(ch) = self.peek() {
                if ch.is_ascii_digit() {
                    self.advance();
                } else {
                    break;
                }
            }
        }
        
        let end_pos = self.current_pos;
        let lexeme = &self.source[start_pos..end_pos];
        
        // Remove underscores for parsing
        let clean: String = lexeme.chars().filter(|&c| c != '_').collect();
        
        let kind = if is_float {
            match clean.parse::<f64>() {
                Ok(n) => TokenKind::Float(n),
                Err(_) => {
                    self.error(start_line, start_column, "Invalid float literal");
                    return;
                }
            }
        } else {
            match clean.parse::<i64>() {
                Ok(n) => TokenKind::Integer(n),
                Err(_) => {
                    self.error(start_line, start_column, "Invalid integer literal");
                    return;
                }
            }
        };
        
        self.tokens.push(Token {
            kind,
            span: Span {
                start: start_pos,
                end: end_pos,
                line: start_line,
                column: start_column,
            },
            lexeme: lexeme.to_string(),
        });
    }
    
    fn scan_hex(&mut self, start_pos: usize, start_line: usize, start_column: usize) {
        while let Some(ch) = self.peek() {
            if ch.is_ascii_hexdigit() || ch == '_' {
                self.advance();
            } else {
                break;
            }
        }
        
        let end_pos = self.current_pos;
        let lexeme = &self.source[start_pos..end_pos];
        let hex_part: String = lexeme[2..].chars().filter(|&c| c != '_').collect();
        
        match i64::from_str_radix(&hex_part, 16) {
            Ok(n) => {
                self.tokens.push(Token {
                    kind: TokenKind::Integer(n),
                    span: Span { start: start_pos, end: end_pos, line: start_line, column: start_column },
                    lexeme: lexeme.to_string(),
                });
            }
            Err(_) => {
                self.error(start_line, start_column, "Invalid hexadecimal literal");
            }
        }
    }
    
    fn scan_binary(&mut self, start_pos: usize, start_line: usize, start_column: usize) {
        // TODO: Similar to scan_hex but for binary
    }
    
    fn scan_octal(&mut self, start_pos: usize, start_line: usize, start_column: usize) {
        // TODO: Similar to scan_hex but for octal
    }
    
    fn scan_identifier(&mut self, first: char, start_pos: usize, start_line: usize, start_column: usize) {
        while let Some(ch) = self.peek() {
            if ch.is_ascii_alphanumeric() || ch == '_' {
                self.advance();
            } else {
                break;
            }
        }
        
        let end_pos = self.current_pos;
        let lexeme = &self.source[start_pos..end_pos];
        
        // Check for keywords
        let kind = match lexeme {
            "let" => TokenKind::Let,
            "mut" => TokenKind::Mut,
            "fn" => TokenKind::Fn,
            "return" => TokenKind::Return,
            "if" => TokenKind::If,
            "else" => TokenKind::Else,
            "while" => TokenKind::While,
            "loop" => TokenKind::Loop,
            "for" => TokenKind::For,
            "in" => TokenKind::In,
            "match" => TokenKind::Match,
            "struct" => TokenKind::Struct,
            "enum" => TokenKind::Enum,
            "impl" => TokenKind::Impl,
            "break" => TokenKind::Break,
            "continue" => TokenKind::Continue,
            "true" => TokenKind::True,
            "false" => TokenKind::False,
            "self" => TokenKind::SelfValue,
            "Self" => TokenKind::SelfType,
            _ => TokenKind::Identifier(lexeme.to_string()),
        };
        
        self.tokens.push(Token {
            kind,
            span: Span {
                start: start_pos,
                end: end_pos,
                line: start_line,
                column: start_column,
            },
            lexeme: lexeme.to_string(),
        });
    }
    
    fn error(&mut self, line: usize, column: usize, message: &str) {
        self.errors.push(LexerError {
            line,
            column,
            message: message.to_string(),
        });
    }
}

#[derive(Debug, Clone)]
pub struct LexerError {
    pub line: usize,
    pub column: usize,
    pub message: String,
}
```

### 5.4 Testing the Lexer

```rust
#[cfg(test)]
mod tests {
    use super::*;
    
    fn tokenize(input: &str) -> Vec<TokenKind> {
        let lexer = Lexer::new(input);
        lexer.tokenize()
            .unwrap()
            .into_iter()
            .map(|t| t.kind)
            .collect()
    }
    
    #[test]
    fn test_simple_tokens() {
        assert_eq!(
            tokenize("+ - * /"),
            vec![TokenKind::Plus, TokenKind::Minus, TokenKind::Star, TokenKind::Slash, TokenKind::Eof]
        );
    }
    
    #[test]
    fn test_multi_char_operators() {
        assert_eq!(
            tokenize("== != <= >= && ||"),
            vec![
                TokenKind::EqEq, TokenKind::NotEq,
                TokenKind::LtEq, TokenKind::GtEq,
                TokenKind::AndAnd, TokenKind::OrOr,
                TokenKind::Eof
            ]
        );
    }
    
    #[test]
    fn test_keywords() {
        assert_eq!(
            tokenize("let mut fn if else"),
            vec![
                TokenKind::Let, TokenKind::Mut, TokenKind::Fn,
                TokenKind::If, TokenKind::Else, TokenKind::Eof
            ]
        );
    }
    
    #[test]
    fn test_identifiers() {
        assert_eq!(
            tokenize("foo bar _baz hello123"),
            vec![
                TokenKind::Identifier("foo".into()),
                TokenKind::Identifier("bar".into()),
                TokenKind::Identifier("_baz".into()),
                TokenKind::Identifier("hello123".into()),
                TokenKind::Eof
            ]
        );
    }
    
    #[test]
    fn test_numbers() {
        assert_eq!(
            tokenize("42 3.14 1_000_000"),
            vec![
                TokenKind::Integer(42),
                TokenKind::Float(3.14),
                TokenKind::Integer(1_000_000),
                TokenKind::Eof
            ]
        );
    }
    
    #[test]
    fn test_strings() {
        let tokens = tokenize(r#""hello" "world\n""#);
        assert_eq!(tokens[0], TokenKind::String("hello".into()));
        assert_eq!(tokens[1], TokenKind::String("world\n".into()));
    }
    
    #[test]
    fn test_function() {
        let input = "fn add(a: i64, b: i64) -> i64 { a + b }";
        let kinds = tokenize(input);
        assert_eq!(kinds[0], TokenKind::Fn);
        assert_eq!(kinds[1], TokenKind::Identifier("add".into()));
        assert_eq!(kinds[2], TokenKind::LParen);
    }
}
```

---

## ✅ Summary

| Component | Purpose |
|-----------|---------|
| `Lexer` struct | Holds state during tokenization |
| `advance()` | Consume next character |
| `peek()` | Look ahead without consuming |
| `match_char()` | Conditional advance |
| `scan_*` methods | Token-specific scanning |
| `Span` | Source location tracking |

---

**Next:** [Lecture 6: Parsing & Building ASTs](../lecture-06/README.md)

**Exercises:** [Exercise Sheet 5](exercises.md)
