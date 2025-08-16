# KeyForge

**KeyForge** is a cross-platform CLI tool written in Go for generating and analyzing passwords and keys.  
It’s the final showcase project for the upcoming book *CLI: A Practical Guide to Creating Modern Command Line Interfaces* (October 2025).  
Learn more at [moderncli.dev](https://moderncli.dev).

---

## ✨ Features

- Generate different types of passwords:
  - **easy** – memorable, pronounceable strings
  - **strong** – secure random strings with full charset
  - **64wep / 128wep / 256wep** – legacy WEP keys (hex)
- Generate **sets** of keys (like [randomkeygen.com](https://randomkeygen.com/))
- **Analyze** passwords offline (entropy + heuristics)
- Configuration system for future AI integration (OpenAI model + API key)
- Works on Linux, macOS (Intel/Apple), Windows, and Raspberry Pi (ARM/ARM64)

---

## Installation

### Prerequisites
- [Go 1.24+](https://go.dev/dl/) installed

### From source
```bash
git clone https://github.com/derickschaefer/keyforge.git
cd keyforge
make build
./keyforge version

## Go install
go install github.com/derickschaefer/keyforge@latest

## Usage

### Help
keyforge help

### Version
keyforge version
v0.1.0

### Create
keyforge create easy --length 16 --count 3
keyforge create strong --length 24
keyforge create 64wep
keyforge create 128wep --count 2
keyforge create 256wep --json
keyforge create set

### Analyze
keyforge analyze "P@ssw0rd123!"
echo "Tr0ub4dor&3" | keyforge analyze --stdin

### Config management
keyforge config list
keyforge config set model gpt-4o-mini
keyforge config test

## Development
make test
make fmt
make build-binaries
make clean

## 📋 Roadmap

- [x] Easy / strong password generators  
- [x] WEP 64/128/256 hex keys  
- [x] Password analyzer (offline)  
- [ ] AI-powered analyzer (OpenAI GPT-4o-mini integration)  
- [ ] Packaging for GitHub Releases  

---

## 📜 License

[MIT](LICENSE)

---

## 🤝 Contributing

Contributions welcome! Feel free to open issues and pull requests.  
