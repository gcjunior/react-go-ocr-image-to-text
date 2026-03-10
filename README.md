# setup

Frontend & backend URLs are configurable per environment. frontend + backend using environment variables for URL and ports.

Inside Docker Compose, frontend can call http://backend:8080 because services are on the same Docker network.

For local npm start, .env will point to http://localhost:8080.

root dir
docker-compose up --build

Frontend: http://localhost:3000

Backend: http://localhost:8080

React calls backend automatically via environment variable.

# Deployment on Render

Backend service: Docker (port 8080), set FRONTEND_URL=https://your-frontend.onrender.com

Frontend service: Static site, build folder, set REACT_APP_BACKEND_URL=https://your-backend.onrender.com

✅ This setup is environment-variable driven, so you can switch URLs and ports per environment.