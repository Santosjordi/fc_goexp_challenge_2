
## Server structure

```
project/
├── main.go             // Entry point
├── handler.go          // HTTP handlers
├── db/
│   ├── repository.go   // Database logic
│   └── schema.sql      // Database schema
├── quote/
│   └── quote.go        // Logic to fetch and process quotes
├── go.mod              // Go module file
```