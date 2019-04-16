# documents-crud

- CRUD of documents(CPF or CNPJ), with flags type and isBlacklisted.
- Backend(server) done in Golang with Echo framework.
- Frontend(client) is SPA done in VueJs.
- Structured to run on Docker, instructions are below:

##Instructions to run:

## Method 1 - Docker
###Pre-requisites: docker and docker-compose installed
```bash
# Run docker compose:
sudo docker-compose up

#Then check on browser: http://localhost:8080

```

## Method 2 - Run locally
###Pre-requisites: docker installed or mongo standalone and NodeJs
```bash
# Run Server:
# mongo needs to be running:
sudo docker run -d -p 27017:27017 mongo:3.2-jessie

# and following env variables specified:
export PORT=1323
export MONGO_PORT=27017
export MONGO_HOST=localhost

cd server

./server

# Run Client:
cd ../client

npm install

npm run serve

```
