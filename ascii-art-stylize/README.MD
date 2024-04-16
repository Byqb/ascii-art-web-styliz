# Ascii Art Web

## Description
Ascii Art Web is a web application that allows users to generate ASCII art using different banners. The server is implemented in Go, providing a simple and interactive graphical user interface (GUI) for creating and displaying ASCII art.

## Authors
- ALI 
- Sayed Ali
- Yusuf  

## Usage: How to Run
1. Clone the repository:
   git clone [repository_url]
   cd ascii-art-web
   go run main.go
   and the website will run 


Implementation Details: Algorithm
The server-side logic for generating ASCII art is implemented in Go. Here's a high-level overview of the algorithm used:

HTTP Endpoints:

GET /:
Renders the main page with a form to input text and choose a banner.
POST /ascii-art:
Handles the form submission, validates input, and generates ASCII art based on the selected banner.
Text Processing:

Handles newline characters to ensure proper line breaks in the ASCII art.
Font Selection:

Reads the selected font file (e.g., static/standard.txt) to retrieve ASCII characters.
ASCII Art Generation:

Maps ASCII characters to their corresponding lines in the selected font.
Generates the final ASCII art by combining lines for each character in the input text.
Error Handling:

Returns appropriate HTTP status codes for different scenarios (OK, Not Found, Bad Request, Internal Server Error).