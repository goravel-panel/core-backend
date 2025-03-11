# Admin Platform Backend

### Local Development Environment

#### Installing Air
```bash
go install github.com/air-verse/air@latest
air -v
```

#### Starting Service
- Copy `.env.example` to `.env`.
- Import the SQL file from the `database/backups/` directory into your database manually.
- Update the `.env` file with the correct settings for the port, database, Redis, and any other configurations.

```bash
go mod tidy  # Install dependencies
./serve.sh   # Run the server
```

Visit http://127.0.0.1:3000/api to confirm the server is running.

