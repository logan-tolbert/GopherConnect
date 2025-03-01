# GopherConnect 🐹🔗

_A simple, lightweight CLI-based contact manager built with Go._

## 📌 Overview

GopherConnect is a fast and minimalistic **CLI contact manager** that allows users to **add, list, update, and delete** contacts.
It is designed as a foundation for future enhancements, keeping things simple while ensuring structured, maintainable code.

## 🚀 Features

✅ Add, list, update, and delete contacts  
✅ In-memory storage for quick access  
✅ Phone number and email validation  
✅ Fully interactive **CLI interface**  
✅ Confirmation prompts to prevent accidental deletions

## 📂 Project Structure

```plaintext
gopher-connect/
│── api/            # (Future API integration)
│── cli/            # Handles CLI interactions
│── models/         # Defines data structures (Person)
│── storage/        # Handles in-memory contact storage
│── handlers/       # (Future API integration)
│── main.go         # Entry point (CLI startup)
│── go.mod          # Go module definition
│── README.md       # Project documentation
```

## 🛠 Installation & Usage

### **Prerequisites**

- Go 1.24

### **Clone the Repository**

```sh
git clone https://github.com/yourusername/gopher-connect.git
```

### **Run the Application**

```sh
cd gopher-connect

go run main.go
```

#### 🛠 **CLI Commands**

1: Add a new contact
2: List all contacts
3: Update a contact’s phone number
4: Update a contact’s email
5: Delete a contact (with confirm)
6: Exit the application

###🧪 **Running Tests**

```sh
go test ./storage
```

## 📌 Future Enhancements

- Implement file/database storage
- Add a REST API (api/, handlers/ folder placeholder)

## 📜 License

This project is licensed under the MIT License.
