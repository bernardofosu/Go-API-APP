
# 🔁 Go File Watchers for Auto-Rebuild

## 1. Air (Most Popular)

**Air** watches for file changes and automatically recompiles & restarts your Go app.

### ✅ Features:
- Fast, efficient, and widely used
- Minimal config
- Works great with Gin, Echo, etc.

### 🔧 Installation:

```bash
go install github.com/air-verse/air@latest
```

Make sure `$GOPATH/bin` is in your system `PATH`.

### 🚀 Usage:

```bash
air
```

Air will look for `.air.toml` in your project root, but works out-of-the-box without it too.

---

## 2. Reflex

Another file watcher for Go. You can trigger custom commands when files change.

### 🔧 Installation:

```bash
go install github.com/cespare/reflex@latest
```

### 🚀 Example usage:

```bash
reflex -r '\.go$' go run main.go
```

---

## 3. CompileDaemon

A simpler watcher with fewer features than Air.

### 🔧 Install:

```bash
go install github.com/githubnemo/CompileDaemon@latest
```

### 🚀 Run:

```bash
CompileDaemon --command="./myapp"
```

---

## 4. Task + Air (for larger projects)

You can combine [Task](https://taskfile.dev/) (like Makefiles for Go) with Air for better build workflows.

---

### ✅ Recommendation

If you're working with **Gin**, go with **Air** — it's fast, clean, and made for Go development.

---

## ℹ️ Extra Tip: Running Air Without main.go

You don’t need to specify `main.go`.

Air detects the main entry point (like `main.go`) automatically.

### Steps:

1. Navigate to your Go project directory.
2. Run:

```bash
air
```

Air will watch `.go` files by default and restart your app on changes.

If needed, customize behavior using `.air.toml`.

---

Would you like a basic example `.air.toml` config?
