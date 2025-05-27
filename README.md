# 🛠️ DevLog CLI

A lightweight, developer-first CLI tool for managing project logs, notes, and debug checklists — built entirely with Go.

> Stay organized. Stay focused. Stay in your terminal.

---

## 📦 Overview

**DevLog CLI** is a terminal-based productivity tool for developers who want a local, version-controlled way to track project notes, debugging tasks, and development progress — all without leaving the command line. No GUI. No bloat. Just logs where and when you need them.

---

## 🚀 Features

### ✅ Project Initialization

```bash
devlog init
```
- Sets up a hidden .devlog directory in your current project directory similar to Git
- Initializes local storage files (logs.json, debug.json).

### 📝 Notes Management
```
devlog note add "Fixed bug in auth logic" "auth"
```
- Adds a timestamped development note with an option tag for future queries.
- All notes are saved locally with UNIX TimeDate styled timestamps.

```
devlog note list 
devlog note list [--tag/-t] "tag"
```
- Lists all saved notes in chronological order when no tag specified.
- Lists notes only having a specific tag for more direct querying.

### 🐞 Debug Checklist
```
devlog debug add "Fix infinite loop in parser"
devlog debug list
devlog debug list [--tag/-t] "tag"
devlog debug squash 2
```
- Manage debug tasks you’ve found during debugging.
- Mark bugs as complete (✅) or incomplete (❌) by ID.

## ⚙️ Installation
### Option 1: Go Install
```
go install github.com/Tyler-Arciniaga/devlog@latest
```
- Make sure $GOPATH/bin is in your PATH.
### Option 2: Manual Build
```
git clone https://github.com/yourusername/devlog.git
cd devlog
make install
```
- Handles go install and moving into PATH.
## 🧱 Tech Stack
- Language: Go
- CLI Framework: spf13/cobra
- Storage Format: JSON

## 🧪 Roadmap & Enhancements
| Feature                      | Status     |
| ---------------------------- | ---------- |
| Core CLI (init, note, debug) | ✅ Done     |
| Git commit summary           | 🧠 Planned |
| Markdown export              | 📝 Planned |
| Config file support          | ⚙️ Planned |
| Unit tests                   | 🧪 Planned |
| Homebrew distribution        | 🍺 Planned |

## 📸 Screenshots
![DevLog screenshot1](./Screenshot%202025-05-27%20at%201.18.47 PM.png)

![Devlog screenshot2](./Screenshot%202025-05-27%20at%201.24.55 PM.png)

## 🙌 Contributing
Contributions are welcome! Feel free to open an issue or submit a pull request with suggestions, bugfixes, or features.


