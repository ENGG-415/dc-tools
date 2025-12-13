# dc-tools

Packages to facilitate assignments and project work in **ENGG-415: Distributed Computing** at Thayer School of Engineering at Dartmouth.

## Overview

This repository contains utility packages designed to support distributed computing coursework and projects, providing foundational tools for network communication, logging, game simulation, card manipulation, and I/O operations.

**Primary author:** Mike Kokko

## Packages

### `logconfig`

Configuration and setup utilities for structured logging throughout the application.

- **File:** `logconfig.go`
- **Purpose:** Centralized logging configuration management

### `mazeconnect`

Hardware simulation and network communication for distributed maze-solving applications.

- **Files:** `hardware.go`, `mazeconnect.go`, `simulator.go`
- **Purpose:** Simulates hardware components and manages connections between distributed maze solvers

### `mazeio`

Input/output operations for maze data handling and testing.

- **Files:** `mazeio.go`, `mazeio_test.go`, `mazeio_test.json`
- **Purpose:** Provides interfaces for reading, writing, and testing maze configurations

### `playingcards`

Playing card utilities for card-based applications and simulations.

- **Files:** `cards.go`, `deck.go`, `doc.go`
- **Key Features:**
  - Card representation with suits (Clubs, Diamonds, Hearts, Spades) and ranks
  - Deck management with shuffle, create, and manipulation functions
  - Card-to-character conversion utilities

## Getting Started

This is a Go module project. To use these packages:

```go
import "github.com/ENGG-415/dc-tools/playingcards"
// or other packages as needed
```

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

This project is part of the ENGG-415 course at Dartmouth College.