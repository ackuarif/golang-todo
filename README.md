## API Todo Application (By : Achmad Kumail Arif)
### Demo Todo Application
```bash
https://todoapp.osc-fr1.scalingo.io/swagger/index.html
```

### How to Run from local computer :
1. Clone Project - 

```bash
git clone https://github.com/ackuarif/golang-todo.git
```
2. Go to the project drectory by `cd golang-todo`.
3. Create `.env` file & Copy `.env.example` file to `.env` file
4. Create a database called - `todo`.
5. Run to update swagger in local computer - 
``` bash
swag init
```
6. Run the server - 
``` bash
go run main.go
```
7. Open Browser - http://localhost:8000 & 
go to API Documentation - http://localhost:8000/swagger
8. You'll see a Swagger Panel.
