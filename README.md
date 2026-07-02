# Go Library GraphQL API

A simple GraphQL backend engine built with **Go Fiber**, **`gqlgen`**, and **GORM** by a **Dockerized PostgreSQL** database container.

---

## Code Generation Workflow

Every time you modify your core GraphQL types or fields, you must regenerate the matching Go structural models.

### 1. Configure the Generator (`gqlgen.yml`)
Ensure your `gqlgen.yml` by generating it, 
```
go tool gqlgen init
```
Configuration mapping references your project structure accurately:

```yaml
schema:
  - graph/*.graphqls
exec:
  filename: graph/generated.go
  package: graph
model:
  filename: graph/model/models_gen.go
  package: model
resolver:
  layout: follow-schema
  dir: graph
  package: graph
```

### 2. Execute Code Generation
After editing your schema blueprint file (`graph/schema.graphqls`), run the native Go tool package command to synchronize your code layouts:
```bash
go tool gqlgen generate
```

---

## 🗄️ Inspect Data in PostgreSQL Container

You can step directly inside your active running database container to view tables, check seed records, or run manual SQL diagnostics.

### 1. Log in via CLI
Run this command in your terminal to open the PostgreSQL interactive terminal (`psql`):
```bash
docker exec -it library_postgres_db psql -U lib_admin -d library_db
```

### 2. Useful SQL Diagnostics
Once logged into the prompt (`library_db=#`), execute these commands to check your structures:
* **List all tables**: `\dt`
* **View seeded books list**: `SELECT * FROM books;`
* **View seeded categories list**: `SELECT * FROM categories;`
* **Exit the database workspace**: `\q`

---

## 🚀 Testing API Methods in Postman

Because GraphQL wraps operation execution inside a single data stream, **all requests (GET and POST equivalents) must be sent using the HTTP `POST` method.**

### 1. Endpoint Target Configurations
* **HTTP Method**: `POST`
* **Request URL**: `http://localhost:8080/books`
* **Payload Type**: Under the input URL bar, navigate to **Body** $\rightarrow$ select the **GraphQL** radio button.

### 2. Fetch All Books (GET Equivalent)
Paste this query text inside your Postman GraphQL body panel to pull all active book arrays:
```graphql
query {
  books {
    id
    title
    author
    category {
      name
    }
  }
}
```

### 3. Fetch Single Book by ID (GET Equivalent)
Pass an explicit argument target ID parameter string to query a singular entity record:
```graphql
query {
  book(id: "1") {
    id
    title
    author
  }
}
```

### 4. Create a New Book (POST Equivalent)
Switch your body syntax layout to a mutation block structure to write a new record into your live container:
```graphql
mutation {
  createBook(input: {
    title: "Designing Data-Intensive Applications"
    author: "Martin Kleppmann"
    isbn: "978-1449373320"
    categoryId: 2
  }) {
    id
    title
    category {
      name
    }
  }
}
```
