# Note-Taker Backend

A backend system for a note-taking application, built using **Go (Golang)** and **DynamoDB**. This project enables authenticated users to securely create, update, delete, and view notes. It is optimized for fast access and scalability, leveraging AWS DynamoDB as a NoSQL datastore.

---

## Table of Contents

- [Features](#features)
- [Technologies Used](#technologies-used)
- [Branch Details](#branch-details)
- [Installation](#installation)

---

## Features

### User Authentication:

- Register with email and password
- Login with JWT-based session
- Logout securely

### Note Management:

- Create a new note
- View a list of notes or a specific note by ID
- Update the content or title of a note
- Delete a note
- Secure endpoints using JWT middleware

---

## Technologies Used

- **Go (v1.20 or higher)**: Backend logic, API endpoints
- **Chi Router**: Lightweight and idiomatic HTTP router for Go
- **DynamoDB**: Fast and scalable NoSQL database by AWS
- **Docker**: Containerized local development environment
- **JWT**: JSON Web Token for secure user authentication
- **AWS SDK for Go**: To interface with DynamoDB

---

## Branch Details

The project uses the following branches for streamlined development:

- **main**:  
  Production-ready, stable version of the application. Deployed versions are pulled from here.

- **dev**:  
  Active development branch. New features, testing, and bug fixes are done here before merging into `main`.

---

## Installation

Follow these steps to set up the project locally:

### Prerequisites

Ensure you have the following installed on your system:

- **Go**: 1.20 or higher
- **Docker**: Installed and running
- **Git**

---

### Steps

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/yash-gkmit/Note-Taker.git
   cd note-taker
   ```

2. **Install dependencies**:

   ```bash
   go mod download
   ```

3. Setup environment variables as given in .env.sample

4. **Run application with Docker**:
   ```bash
   docker-compose up -d
   ```
5. Verify the API is Running.
