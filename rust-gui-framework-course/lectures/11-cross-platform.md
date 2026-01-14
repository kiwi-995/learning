# Lecture 11: Cross-Platform Deployment & Architecture

> **Duration:** ~1.5-2 hours  
> **Prerequisites:** Lectures 01-10  
> **Goal:** Deploy your framework to web, desktop, and mobile

---

## Overview

Your GUI framework should run everywhere. We'll cover platform abstraction and deployment strategies.

---

## Part 1: Platform Abstraction (20 min)

### Architecture

```
┌─────────────────────────────────────────┐
│           Your Application              │
├─────────────────────────────────────────┤
│          GUI Framework Layer            │
├─────────────────────────────────────────┤
│         Platform Abstraction            │
├──────────┬──────────┬──────────┬────────┤
│  Desktop │   Web    │   iOS    │ Android│
│  (wgpu)  │ (WebGPU) │ (Metal)  │(Vulkan)│
└──────────┴──────────┴──────────┴────────┘
```

### Platform Trait

```rust
pub trait Platform {
    type Window;
    type Renderer;
    
    fn create_window(&self, config: WindowConfig) -> Self::Window;
    fn create_renderer(&self, window: &Self::Window) -> Self::Renderer;
    fn run_event_loop<F: FnMut(Event)>(self, handler: F);
}
```

---

## Part 2: Web Deployment (25 min)

### wasm-bindgen Setup

```toml
# Cargo.toml
[lib]
crate-type = ["cdylib"]

[dependencies]
wasm-bindgen = "0.2"
web-sys = { version = "0.3", features = ["Window", "Document", "HtmlCanvasElement"] }
wgpu = { version = "0.19", features = ["webgl"] }
```

### Building for Web

```bash
# Install tools
cargo install wasm-pack

# Build
wasm-pack build --target web

# Generated files in pkg/
# - my_gui_bg.wasm
# - my_gui.js
```

### HTML Integration

```html
<!DOCTYPE html>
<html>
<head>
    <style>
        canvas { width: 100%; height: 100vh; display: block; }
    </style>
</head>
<body>
    <canvas id="canvas"></canvas>
    <script type="module">
        import init from './pkg/my_gui.js';
        await init();
    </script>
</body>
</html>
```

---

## Part 3: Desktop Packaging (20 min)

### Windows/macOS/Linux

```bash
# Simple: cargo build --release
# Binary in target/release/

# For distribution, use cargo-bundle
cargo install cargo-bundle
cargo bundle --release
```

### Cargo.toml for bundling

```toml
[package.metadata.bundle]
name = "My GUI App"
identifier = "com.example.myguiapp"
icon = ["icons/icon.icns", "icons/icon.ico"]
version = "1.0.0"
```

---

## Part 4: Conditional Compilation (20 min)

### Platform-Specific Code

```rust
#[cfg(target_arch = "wasm32")]
mod web_platform;

#[cfg(not(target_arch = "wasm32"))]
mod native_platform;

#[cfg(target_os = "macos")]
fn get_system_font() -> &'static str { "SF Pro" }

#[cfg(target_os = "windows")]
fn get_system_font() -> &'static str { "Segoe UI" }

#[cfg(target_os = "linux")]
fn get_system_font() -> &'static str { "Ubuntu" }
```

### Feature Flags

```toml
[features]
default = ["desktop"]
desktop = []
web = ["wasm-bindgen", "web-sys"]
mobile = []
```

---

## Part 5: Accessibility (20 min)

### AccessKit Integration

```rust
use accesskit::{NodeBuilder, Role, TreeUpdate};

fn build_accessibility_tree(root: &Widget) -> TreeUpdate {
    let mut nodes = Vec::new();
    
    fn visit(widget: &Widget, nodes: &mut Vec<(NodeId, Node)>) {
        let mut node = NodeBuilder::new(widget.role());
        
        if let Some(label) = widget.accessible_label() {
            node.set_name(label);
        }
        
        node.set_bounds(widget.bounds().into());
        
        nodes.push((widget.id(), node.build()));
        
        for child in widget.children() {
            visit(child, nodes);
        }
    }
    
    visit(root, &mut nodes);
    TreeUpdate { nodes, /* ... */ }
}
```

---

## Part 6: CI/CD Pipeline (10 min)

### GitHub Actions

```yaml
name: Build

on: [push]

jobs:
  build:
    strategy:
      matrix:
        os: [ubuntu-latest, windows-latest, macos-latest]
    runs-on: ${{ matrix.os }}
    steps:
      - uses: actions/checkout@v3
      - uses: dtolnay/rust-toolchain@stable
      - run: cargo build --release
      
  web:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: dtolnay/rust-toolchain@stable
        with:
          targets: wasm32-unknown-unknown
      - run: cargo install wasm-pack
      - run: wasm-pack build --target web
```

---

## Key Takeaways

1. **wgpu runs everywhere** - native and web
2. **wasm-pack** for web builds
3. **cargo-bundle** for desktop packaging
4. **Conditional compilation** for platform differences
5. **Accessibility is important** - use AccessKit

---

## Up Next

**Lecture 12**: Capstone project - build a complete application!
