FROM node:latest

WORKDIR /app

COPY package*.json ./

RUN npm install
RUN npm install pm2@latest -g

COPY . .

COPY .env.example .env

CMD ["pm2-runtime", "./src/server.js"]
