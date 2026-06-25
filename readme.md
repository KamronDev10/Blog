# Blog API 📝

A blog backend built with Go.

## Tech Stack
- Go
- PostgreSQL
- JWT Authentication
- Swagger UI

## Getting Started

Fill in the .env file:
```
DB_URL=postgres://...
JWT_SECRET=your-secret-key
```

```bash
go run src/main/main.go
```

Swagger: `http://localhost:8080/swagger/index.html`

## API Endpoints

### Auth
- `POST /auth/sign-up` — Register
- `POST /auth/sign-in` — Login

### Articles
- `GET /articles` — Get all articles
- `GET /articles/get` — Get single article
- `POST /articles/create` — Create article 🔒
- `PUT /articles/update` — Update article 🔒
- `DELETE /articles/delete` — Delete article 🔒


### Comments
- `GET /comments` — Get comments by article
- `POST /comments/create` — Create comment 🔒
- `DELETE /comments/delete` — Delete comment 🔒



### Tags
- `GET /tags` — Get all tags
- `GET /tags/get` — Get single tag
- `POST /tags/create` — Create tag 🔒
- `DELETE /tags/delete` — Delete tag 🔒


> 🔒 — Requires token

yana qoshish mumkin boladi 

api qosh 