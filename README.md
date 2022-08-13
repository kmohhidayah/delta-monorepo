## delta-monorepo

## Prerequisite

- go1.19
- Node
- PostgreSQL
- Docker

## Getting Started

### Clone Project

```bash
git clone https://github.com/kmohhidayah/delta-monorepo.git
```

### Docker

```bash
# give access to exec dir db/
chmod +x ./db/init-user-db.sh

# docker build & start
docker-compose up -d

# docker stop & remove
docker-compose down -v
```

### Manualy

#### auth-app

```bash
cd delta-monorepo
cd auth-app
cp .env.example .env
make postgres
make createdb
```

#### fetch-app

```bash
cd delta-monorepo
cd fetch-app
cp .env.example .env
yarn
yarn dev
```

#### API Documentation
**Auth-app**
- https://documenter.getpostman.com/view/17968129/VUjSF3hX

**Fetch-App**
- https://documenter.getpostman.com/view/17968129/VUjSF3mq

#### Document C4 (Context and Deployment)

<img title="a title" alt="Alt text" src="./docs/Screenshot 2022-08-12 20:07:47.png">

<img title="a title" alt="Alt text" src="./docs/Screenshot 2022-08-12 20:02:17.png">
