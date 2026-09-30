# gubbylist-
gubby.list over powered word list generating maded by go language 
# 🚀 gubby.list 🚀

An ultra-fast, advanced, and **completely storage-less infinite wordlist generator** written in Go (Golang). Specially optimized for high-performance and low-resource environments like **Termux (Android)** and Linux.

---

## ⚡ Key Features
- **Zero Disk Write:** It does NOT write massive text files to your phone's storage. Everything is generated and deduplicated instantly inside the RAM.
- **CUPP + Flexible Mode:** Can perform targeted user profiling or accept completely custom flexible inputs.
- **Infinite Loop:** Generates non-stop combinations dynamically until you explicitly stop it using `Ctrl + C`.
- **Live Memory Endpoint:** Once stopped, it creates a high-speed local HTTP web link. You can pass this URL directly into other tools (like Hydra, Gobuster, etc.) without downloading any physical files.

---

## 🛠️ Step-by-Step Installation & Usage Guide

Open your Termux terminal and execute the following commands sequentially:

### Step 1: Update Packages & Install Dependencies
First, make sure your Termux repository is up-to-date and install the Go compiler along with Git:
```bash
apt update && apt upgrade -y
apt install golang git -y
```

### Step 2: Clone the Repository
Download the source code of the tool from GitHub into your system:
```bash
git clone https://github.com
```

### Step 3: Fire Up the Tool
Run the program instantly using the Go compiler:
```bash
go run main.go
```

---

## 🛑 How to stop and get the link?
1. When the tool is running in infinite mode, let it generate as many variations as you need.
2. Once you are satisfied, press **`Ctrl + C`** on your keyboard (or use the Termux special key bar).
3. The tool will stop generation immediately and print your live memory endpoint link: `http://localhost:9999/gubby.list`
4. Copy this URL directly into your favorite security testing tools!

   
