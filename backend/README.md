kill -9 $(lsof -t -i:8080) // if port in use
open -a Docker
docker build -t go-project .
docker run -p 8080:8080 go-project
http://localhost:8080/ocr // it woll show 405 Method not allowed

docker ps
docker kill container_id


docker run -it go-project sh // it will give you ssh access
ls -la /app/uploads
touch /app/uploads/test.txt // test adding a file

curl -X POST http://localhost:8080/ocr -F "image=@images/form-filled.png"

curl -X OPTIONS http://localhost:8080/ocr -i

HTTP/1.1 200 OK
Access-Control-Allow-Headers: Content-Type
Access-Control-Allow-Methods: GET, POST, OPTIONS
Access-Control-Allow-Origin: *
Date: Mon, 09 Mar 2026 01:57:58 GMT
Content-Length: 0