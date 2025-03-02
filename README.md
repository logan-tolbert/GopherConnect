# GopherConnect 🐹🔗

_A simple, lightweight CLI-based contact manager built with Go._

## 📌 Overview

GopherConnect is a fast and minimalistic **CLI contact manager** that allows users to **add, list, update, and delete** contacts.
It is designed as a foundation for future enhancements, keeping things simple while ensuring structured, maintainable code.

## 🚀 Features

✅ Add, view, update, and delete contacts  
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

```sh
1. List Contacts
2. Add Contacts
3. Get Contact By Name
4. Update Phone Number
5. Update Email Address
6. Delete Contact
6. Exit
```

###🧪 **Running Tests**

```sh
go test ./storage
```

## 📌 Future Enhancements

- Implement file/database storage
- Add a REST API (api/, handlers/ folder placeholder)

## 📜 License

This project is licensed under the MIT License.
