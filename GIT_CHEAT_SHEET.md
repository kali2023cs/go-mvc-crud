# Git Commands Cheat Sheet for New Projects

Follow these steps to initialize and push a new project to GitHub.

### 1. Initialize Git
Run this once inside your project folder.
```bash
git init
```

### 2. Create a .gitignore File
Prevent unnecessary files (like `.env` or binaries) from being tracked.
```bash
# Handled automatically if you create a .gitignore file manually or via your IDE
```

### 3. Stage Your Files
Add all files to the "Staging Area".
```bash
git add .
```

### 4. Create Your First Commit
Save your staged files into a local snapshot.
```bash
git commit -m "Initial commit: Your project description"
```

### 5. Create the Remote Connection
Link your local folder to a GitHub repository URL.
```bash
git remote add origin https://github.com/YOUR_USERNAME/YOUR_REPOSITORY.git
```

### 6. Rename Default Branch (Standard Practice)
Ensure your main branch is named `main`.
```bash
git branch -M main
```

### 7. Push to GitHub
Upload your code to the remote repository. The `-u` flag remembers your settings for next time.
```bash
git push -u origin main
```

---

### Useful Status Commands
- **Check Status**: `git status` (See what files are changed or untracked)
- **Check History**: `git log --oneline` (See past commits)
- **Check Remote**: `git remote -v` (See where your code is pushing to)
