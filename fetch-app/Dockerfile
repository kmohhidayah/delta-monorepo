FROM node:latest as builder
WORKDIR /usr/app
COPY package*.json ./
RUN npm install

COPY . .
COPY .env.example .

EXPOSE 3000
CMD ["npm", "run", "dev"]
