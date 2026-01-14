# Exercise Sheet 11: Cross-Platform Deployment

> **Estimated Time:** 1.5 hours  
> **Difficulty:** ⭐⭐⭐

---

## Exercise 1: Web Build (30 min)

1. Add wasm-bindgen to your project
2. Create `index.html` with canvas
3. Build with wasm-pack
4. Run in browser with a local server

Verify: same visual output as desktop.

---

## Exercise 2: Platform Abstraction Layer (30 min)

Create platform-agnostic file loading:
```rust
pub async fn load_font(name: &str) -> Vec<u8>;
pub async fn load_image(path: &str) -> ImageData;
```

Implement for both native (std::fs) and web (fetch API).

---

## Exercise 3: Desktop Packaging (20 min)

1. Add app icon (create or find one)
2. Configure cargo-bundle
3. Build macOS .app or Windows .exe
4. Test the packaged application

---

## Exercise 4: Accessibility Basics (20 min)

Add screen reader support:
1. Mark buttons with accessible names
2. Mark text labels appropriately  
3. Implement keyboard navigation (Tab to move, Enter to activate)
4. Test with VoiceOver/NVDA

---

## Theory Problems

### T1: WASM Size
Your desktop binary is 5MB. The .wasm is 2MB.
1. Why is WASM smaller?
2. What techniques reduce WASM size further?
3. What's the trade-off with `wasm-opt -O3`?

### T2: WebGPU Fallbacks
Not all browsers support WebGPU.
1. What's the fallback strategy?
2. How do you detect capabilities?
3. What features might not work in WebGL2 fallback?

---

## Challenge: Mobile Build

Build for iOS or Android:
1. Set up cargo-mobile
2. Create minimal mobile app
3. Handle touch correctly
4. Deploy to simulator/device
