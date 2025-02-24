# devops-lecture-project-2025

Project for the lecture "Einführung in DevOps, Continuous Delivery Tools und Mindset"

## Description of the web-shop

The products offered by this web-shop are as follows:

| Product Identifier | Name        | Price |
|--------------------|-------------|-------|
| 4                  | Hammer      | 9.99  |
| 5                  | Wrench      | 14.99 |
| 6                  | Screwdriver | 7.99  |
| 7                  | Drill       | 49.99 |
| 8                  | Pliers      | 12.99 |

## Running the service locally

### Using the Docker CLI

Either pull the Docker image
```sh
docker pull ghcr.io/hfxbse/dhbw-devops/server:latest
```

Or build the Docker image
```sh
docker build -f infrastructure/Dockerfile . -t ghcr.io/hfxbse/dhbw-devops/server:latest
```

Then run the Docker image
```sh
docker run -p 8080:8080 ghcr.io/hfxbse/dhbw-devops/server:latest
```

### Using Docker Compose

The setup above can be simplified with Docker Compose
```sh
docker compose -f infrastructure/docker-compose.yml up
```

### Using Vagrant

```sh
vagrant up
```
