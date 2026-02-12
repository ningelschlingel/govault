# GoVault

A lightweight macOS Finder & CLI utility for quick file encryption, built with Go and the `age` encryption library.

### 🔒 Features
- **Native Integration:** Right-click any file to lock/unlock via macOS Quick Actions.
- **Strong Encryption:** Uses the [age](https://github.com/FiloSottile/age) format (X25519, ChaCha20-Poly1305, and scrypt).

### 🚀 Installation

1. **Build the binary:**
   ```bash
   go build -o govault .
   ````

2. Install to System:
    Move the binary to your local path:
    ```bash
    sudo cp govault /usr/local/bin/
    ```
3. Enable Right-Click Menu:
    Copy the `GoVault-Toggle.workflow` folder to `~/Library/Services/`.

    _Note: You may need to enable it under System Settings > Extensions > Finder_.

# 📖 Usage

- Via Finder: Right-click any file → Quick Actions → GoVault-Toggle. A native macOS dialog will prompt for your password.
- Via CLI:
    ```bash
    govault <filename> <password>
    ```

# ⚠️ Disclaimer

This was a 2-hour "test run" project to explore the Go ecosystem. While it uses a state-of-the-art encryption library, this implementation is provided "as-is."

Password Recovery: There is none. If you lose the password, the data is gone.

File Safety: Always keep a backup of critical data before encrypting.

# Outlook

If at any point more time is spent on this project, it will be used to:

- Make use of AGE's ssh-key functionality as a default with a password-backup (e.g. for decrypting on another machine)
- Improve fail safety / error handling, rather optimistiv right now
- Fix password leak issue in CLI; priority was finder-integration
- Add support for folders and make use of go-routines for concurrency
- Tests

