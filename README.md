# Simple Startup Application Management

Before we start, this was the projects that me and my friend make for making an application startup management program using GOLANG as a backend,
A lightweight console-based application built with Go that allows users to manage startups, team members, and funding information efficiently. 
Ideal for small startup teams or academic purposes, this project demonstrates fundamental programming concepts including structs, arrays, sorting, and searching in Go.

---

## Features

- Add, view, and manage startups
- Assign team members with specific roles
- Search startups by name (Sequential Search)
- Sort startups by funding (Selection Sort)
- Generate category-based summary reports
- Console-based interface with clean prompts

---

## Technologies Used

- **Programming Language**: Go (Golang)
- **Data Structure**: Arrays and Structs
- **CLI Interface**: `fmt.Scan`, `fmt.Printf`
- **Algorithms**: 
  - Sequential Search
  - Selection Sort
  - Frequency counting

---

## Project Structure

```bash
Project.go     # Main Go application with routes, function, and logic
templates/     # HTML templates for Gin
layout.html    # Base layout (header/nav)
index.html     # Home page
add.html       # Add startup form
view.html      # View/search/sort/report/delete interface
README.md      # You're reading it!