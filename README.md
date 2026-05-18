# 🚀 Social Backend API (Go + Postgres + Redis)

A modular backend project built with **Golang**, featuring user management, friendships, and posts (with image support).  
Designed for learning and scalability — fully containerized with Docker & Docker Compose.  

---

## 📌 Features
- 👤 **User APIs** – Add, list, and delete users.  
- 🤝 **Friendship APIs** – Send/accept friendships, list friends.  
- 📝 **Post APIs** – Create posts (text/images), view all posts.  
- ⚡ **Redis Caching** – Speeds up repeated queries.  
- 🗄️ **Postgres Database** – Reliable persistence.  
- 🐳 **Dockerized** – Easy setup & deployment.  

---

## 📂 Project Structure
```bash
.
├── cmd/                # App entrypoint(s)
├── controllers/        # API controllers (handle requests)
├── services/           # Business logic
├── models/             # Database models
├── routes/             # API routes
├── utils/              # Helpers & utilities
├── init_db_table/               # Database migrations/init scripts
├── docker-compose.yaml # Multi-service setup
├── Dockerfile          # Go app container
├── go.mod              # Go dependencies
└── README.md           # This file

1. ⚙️ Setup & Run
 -  Clone the repo
 -  git clone https://github.com/Nikhiliitg/SocialMediaPlateform.git
 - cd <repo-name>
2. Run with Docker Compose
docker-compose up --build
This will start:
-  Go app on localhost:3015
-  Postgres on localhost:5432
-  Redis on localhost:6379

3. 🛠️ API Endpoints (Sample)
 - Users
 - POST /users – Create a user
 - GET /users – Get all users
 - DELETE /users/:id – Delete user
 - Friends
 - POST /friends/:userId/:friendId – Add friend
 - GET /friends/:userId – Get all friends

Posts
 - POST /posts – Create a post (text/image)
 - GET /posts – Get all posts
 - 🗄️ Database Schema (Simplified)

```bash 
 CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT,
    email TEXT,
    password TEXT,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS posts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    content TEXT,
    user_id UUID REFERENCES users(id),
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS friendships (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id),
    friend_id UUID REFERENCES users(id),
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);
```

🐳 Deployment

Push your image to DockerHub:
docker build -t <dockerhub-username>/social-backend .
docker push <dockerhub-username>/social-backend

🎯 Roadmap

 - ✅ Basic CRUD for users/friends/posts
 - ✅ Redis caching
 - ✅ Dockerized environment

**  Contributing **
Contributions welcome! Fork, branch, and PR.



docker run --name database -p 5432:5432 -e POSTGRES_USER=<user-name> -e POSTGRES_PASSWORD=<password> -e POSTGRES_DB=<db-name>  postgres:14.23-alpine

docker run --name redis -p 6379:6379 redis:alpine3.23