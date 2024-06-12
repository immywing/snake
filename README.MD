# Snake

This repository contains the source code for the Snake game, divided into two directories: `snake-back-end` and `snake-front-end`.

## Prerequisites

Before running the project, ensure you have the following installed on your system:

- Go (version 1.22)
- Node.js (with npm)
- TypeScript

## Directory Structure

- `snake-back-end`: Contains the Go backend server.
- `snake-front-end`: Contains the React/TypeScript frontend application.

## Getting Started

### Running the Backend

1. Navigate to the `snake-back-end` directory:
   ```
   cd snake-back-end
   ```
1. Install the necessary Go dependencies:
    ```
    go mod tidy
    ```
1. Start the backend server:
    ```
    go run .
    ```
The backend server will be running on port 8080.

### Running the Frontend

1. Navigate to the `snake-front-end` directory:
    ```
    cd snake-back-end
    ```
1. Install the necessary dependencies:
    ```
    npm install
    ```
1. Start the frontend application:
    ```
    npm start
    ```
The frontend application will be running on port 3000.

### Playing the game
The game is controllowed with `W`   `A`   `S`   `D` for directional inputs.

### Known issues

1. Directional inputs from a player in quick succession can allow the player to turn the snake back on itself causing a game over state.

1. Game over states that occur when the snake collides with a wall currently cause an error on the backend server due to an index out of range error on the array.