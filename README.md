# Go Monitor

A modern terminal-based system monitoring tool written in Go, featuring a sleek and interactive terminal user interface.

## Overview

go-monitor is a lightweight system monitoring application that provides real-time insights into your system's performance. Built with Go and the Charm libraries (Bubbletea, Bubbles, and Lipgloss), it offers an elegant terminal user interface for monitoring system resources.

## Features

- **Real-time Process Monitoring**: View running processes in an interactive table interface
- **System Resource Statistics**:
  - CPU usage monitoring
  - Memory usage tracking (total, used, free, cached, etc.)
  - Network statistics
- **Modern Terminal UI**:
  - Responsive table layout
  - Color-themed interface with light/dark mode support
  - Interactive selection and navigation
  - Regular updates with millisecond precision

## Dependencies

- [charmbracelet/bubbletea](https://github.com/charmbracelet/bubbletea) - Terminal UI framework
- [charmbracelet/bubbles](https://github.com/charmbracelet/bubbles) - TUI components
- [charmbracelet/lipgloss](https://github.com/charmbracelet/lipgloss) - Style definitions
- [shirou/gopsutil](https://github.com/shirou/gopsutil) - System statistics collection

## Installation

```bash
go install github.com/hambosto/go-monitor/cmd/go-monitor@latest
```

Or clone and build from source:

```bash
git clone https://github.com/hambosto/go-monitor.git
cd go-monitor
go build ./cmd/go-monitor
```

### Using Nix Flakes

Add the flake to your inputs:

```nix
# In your flake.nix
inputs = {
  wallpaper-manager.url = "github:hambosto/go-monitor";
};
```

#### NixOS/Home Manager

Enable the module in your configuration:

```nix
# In your configuration.nix or home.nix
{ inputs, ... }: {
  imports = [
    inputs.go-monitor.nixosModules.default
  ];

  programs.go-monitor.enable = true;
}

## Usage

Simply run the binary after installation:

```bash
go-monitor
```

### Interface Controls

- Use arrow keys to navigate through the process table
- Press `q` to quit the application
- The interface updates automatically to show the latest system statistics

## Project Structure

```
.
├── cmd/
│   └── go-monitor/       # Main application entry point
├── internal/
│   ├── system/          # System metrics collection
│   └── tui/            # Terminal UI components
│       ├── components/  # Reusable UI components
│       └── theme/      # UI styling and colors
```

## Contributing

Contributions are welcome! Feel free to submit issues and pull requests.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Charm](https://charm.sh/) for their excellent terminal UI libraries
- [gopsutil](https://github.com/shirou/gopsutil) for system metrics collection