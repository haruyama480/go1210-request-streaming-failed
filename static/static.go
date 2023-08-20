package static

import "embed"

//go:embed index.html main.wasm wasm_exec.js favicon.ico
var Assets embed.FS
