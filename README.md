# ASCII Art Web

## Description

ASCII Art Web is a web-based implementation of the ASCII Art project written in Go. The application allows users to enter text through a graphical user interface (GUI), select an ASCII banner style, and generate character-based ASCII artwork directly in the browser.

The project is built using Go's standard library and follows a client-server architecture. Users interact with a web page, submit text through a form, and the server processes the input using ASCII banner files before returning the generated ASCII art.

Supported banners:

* standard
* shadow
* thinkertoy

---

## Authors

**Innocent Ekwueme**

Aspiring Full Stack Software Engineer and Learn2Earn Fellow.

---

## Usage: How to Run

### Clone the repository

```bash
git clone <repository-url>
cd ascii-art-web
```

### Run the application

```bash
go run .
```

The server will start on:

```text
http://localhost:8080
```

Open your browser and visit:

```text
http://localhost:8080
```

### Using the application

1. Enter text into the input field.
2. Select a banner style.
3. Click **Generate Art Now**.
4. The generated ASCII art will be displayed on the page.

---

## HTTP Endpoints

### GET /

Returns the main HTML page containing:

* Text input field
* Banner selection
* Submit button

Response:

```http
200 OK
```

---

### POST /ascii-art

Receives:

* User text
* Selected banner

Processes the input and generates ASCII art.

Response:

```http
200 OK
```

if generation is successful.

---

## HTTP Status Codes

The server returns appropriate HTTP status codes:

| Status Code               | Description                               |
| ------------------------- | ----------------------------------------- |
| 200 OK                    | Request completed successfully            |
| 400 Bad Request           | Invalid or empty input                    |
| 404 Not Found             | Route, template, or banner file not found |
| 405 Method Not Allowed    | Invalid HTTP method                       |
| 500 Internal Server Error | Unexpected server error                   |

---

## Project Structure

```text
ascii-art-web/
│
├── main.go
├── handlers.go
├── utils.go
├── go.mod
│
├── banners/
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
│
├── template/
│   └── index.html
│
└── static/
    └── style.css
```

---

## Implementation Details

### Overview

The application follows a request-response workflow.

```text
Browser
   │
   ├── GET /
   │
   ▼
homeHandler
   │
   ▼
HTML Page
   │
   ▼
User enters text
   │
   ▼
POST /ascii-art
   │
   ▼
asciiHandler
   │
   ▼
Read banner file
   │
   ▼
Parse banner file
   │
   ▼
Generate ASCII Art
   │
   ▼
Render template
   │
   ▼
Display result
```

---

### Banner Processing Algorithm

#### Step 1: Read Banner File

The selected banner file is loaded from disk.

```go
content, err := readBannerFile(banner)
```

---

#### Step 2: Parse Banner File

The file is converted into a map where each printable ASCII character is associated with its 8-line ASCII representation.

```go
map[rune][]string
```

Example:

```text
'A' -> [8 lines]
'B' -> [8 lines]
'C' -> [8 lines]
```

---

#### Step 3: Generate ASCII Art

For each character in the user input:

1. Locate the character in the map.
2. Retrieve its 8-line representation.
3. Combine all characters row by row.
4. Build the final ASCII art output.

---

### Template Rendering

Go templates are used to send generated ASCII art back to the webpage.

Example:

```go
data := Data{
    Message: result.String(),
}
```

The template accesses the value using:

```html
{{.Message}}
```

---

## Technologies Used

* Go
* HTML5
* CSS3
* Go Templates
* net/http Package

---

## Features

* Web-based ASCII art generator
* Multiple banner styles
* Form submission using POST requests
* Template rendering
* Error handling
* HTTP status code management
* Modular code structure

---

## Learning Outcomes

This project demonstrates understanding of:

* HTTP servers in Go
* Request handling
* Routing
* HTML templates
* Form processing
* File handling
* Error management
* ASCII art generation algorithms
* Client-server architecture

---

## Future Improvements

Potential enhancements include:

* Custom error pages
* Dark mode UI
* Download generated ASCII art
* Copy-to-clipboard functionality
* Additional banner styles
* Responsive mobile design
* Docker deployment

---

## Example

Input:

```text
Hello
```

Banner:

```text
standard
```

Output:

```text
 _   _          _
| | | |   ___  | |
| |_| |  / _ \ | |
|  _  | |  __/ | |
|_| |_|  \___| |_|
```

---

## License

This project was created for educational purposes as part of the ASCII Art Web project.
