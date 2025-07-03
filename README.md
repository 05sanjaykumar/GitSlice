# 🧩 GitSlice CLI

**A CLI tool to extract specific folders or file from GitHub repositories using sparse-checkout.**  
Designed for speed, simplicity, and developers who don't want to clone massive repos unnecessarily.

---

## 🚀 Features

- ⚡️ Extract a specific folder or file from any GitHub repo
- 🔍 Automatically resolves deep nested paths like `/tree/branch/src/utils/`
- 🌲 Uses Git sparse-checkout for minimal clone footprint
- 🧹 Cleans up temp files after extraction
- 🔄 Supports both full clone and sparse mode
- 🧠 Intelligent path matching without needing full repo history

---

## 📦 Installation

### With Go

```bash
go install github.com/05sanjaykumar/gitslice@latest
````

### Or clone and build:

```bash
git clone https://github.com/05sanjaykumar/gitslice
cd gitslice
go build -o gitslice main.go
```

---

## 🛠️ Usage

### Basic Syntax

```bash
gitslice <github-folder-or-file-url>
```

### Examples

```bash
gitslice https://github.com/supabase/storage/tree/fix/pgboss-on-error-callback/src/auth
```

```bash
gitslice https://github.com/vercel/next.js/tree/canary/packages/next
```

```bash
gitslice https://github.com/05sanjaykumar/Flappy-Bird-OpenCV/blob/main/assets/background-day.png
```

### Output

* Extracts folder/file to the current directory.
* Auto-detects branch and path.

---

## 📋 CLI Options

```bash
gitslice --help      # Show help
gitslice --version   # Show version
```

---

## 🧠 How It Works

1. Parses the GitHub URL and extracts the owner, repo, branch, and folder or file path.
2. Clones the repo in sparse mode without full checkout.
3. Sets up sparse-checkout for the specific folder or file.
4. Switches to the correct branch and checks out only the required path.
5. Moves the folder or file to your working directory and cleans up.

---

## ⚠️ Limitations & Edge Cases

Despite its power, `GitSlice` does have a few **known limitations**:

* **Branch-specific folders only**: If a folder or a file exists *only* in a non-default branch (e.g., not in `main`), `GitSlice` might not detect or clone it correctly unless the branch is explicitly specified in the URL.

* **Private repositories**: GitSlice currently only works with **public GitHub repositories**. Support for private repos (with authentication) is not yet implemented.

* **Path guessing limitations**: The CLI resolves folder or file paths by pattern matching on local clone results. In very complex or dynamically generated repos, edge cases may slip through.

* **Shallow branch detection**: We don't fetch the full remote branch list — only the default or specified branch is used for sparse-checkout. If you mistype the branch name or forget it in the URL, cloning will fail.

* **No submodule support (yet)**: Git submodules aren't handled or fetched. This tool assumes you're targeting standalone folders.

* **GitHub-only**: Currently supports only GitHub URLs. Support for GitLab, Bitbucket, etc., is on the roadmap.

---

## 🔒 Requirements

* Git installed and available in PATH
* Go 1.18+ (for building from source)

---

## 📄 License

MIT © [Sanjay Kumar](https://github.com/05sanjaykumar)

---

## ❤️ Contributing

Pull requests and issues are welcome!
If you hit an edge case or repo structure the tool doesn't support, open an issue and let’s improve it together.

---

## 🌟 Why GitSlice?

Because cloning an entire repo when you just want one folder or file is like ordering a pizza and getting the whole restaurant.
GitSlice gives you just the slice you need.


