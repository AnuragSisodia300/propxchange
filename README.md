# Propxchange API

Propxchange is a Go-based backend system designed to manage users, properties, and KYC (Know Your Customer) processes. It includes features like user management, favorites, and file uploads to a GCP bucket.

## Features

- **User Management**: Create, update, and manage users.
- **Favorites**: Add, list, and remove properties from a user's favorites.
- **KYC Process**: Three-step KYC workflow including personal, financial, and document uploads.
- **File Uploads**: Upload KYC documents to a structured GCP bucket.

## Project Structure

```plaintext
.
├── config/                 # Configuration files (e.g., GCP credentials)
├── controllers/            # Business logic for handling requests
├── routes/                 # Route definitions
├── schema/                 # Data models and structs
├── test/                   # Test cases
├── main.go                 # Entry point of the application
├── go.mod                  # Go module dependencies
├── go.sum                  # Dependency checksum file
├── .gitignore              # Ignored files and directories
└── README.md               # Project documentation
