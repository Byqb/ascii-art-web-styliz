# Ascii Art Web

## Description
Ascii Art Web is a web application that allows users to generate ASCII art using different banners. The server is implemented in Go, providing a simple and interactive graphical user interface (GUI) for creating and displaying ASCII art.

## Usage: 
How to Run: 
1. Start the server: 
```
go run main.go
```
2. Open a web browser and visit:
```
http://localhost:8080
```

## Implementation Details:
The Ascii-Art-Web application is built using the following technologies and concepts:
- Go programming language for the server-side implementation.
- HTML, CSS, for the frontend user interface.
- Go templates for rendering dynamic HTML pages.
- HTTP endpoints for handling different actions and requests.

The server provides the following HTTP endpoints:
- GET /: Returns the main HTML page, which includes a form for inputting text and selecting a banner style.
- POST /ascii-art: Accepts data from the client, including the user's text and selected banner style, and generates the corresponding ASCII art.

HTTP Status Codes:
- 200 OK: Returned when the request is successful and the ASCII art is generated without errors.
- 404 Not Found: Returned when necessary templates or banners are not found.
- 400 Bad Request: Returned when the request is incorrect or missing required parameters.
- 500 Internal Server Error: Returned when unhandled errors occur during the server-side processing.

Algorithm:
The algorithm for generating ASCII art banners involves the following steps:
1. Receive the user's text and selected banner style from the client.
2. Validate the input data to ensure it is correct and complete.
3. Retrieve the corresponding ASCII art banner template based on the selected style.
4. Replace specific characters in the template with the user's text to generate the final ASCII art.
5. Return the ASCII art as a response to the client.

### In the web:

https://github.com/user-attachments/assets/82550a5d-e23f-46eb-b90c-e6d0997009c0



## Authors
- ALI 
- Sayed Ali
- Yusuf  
