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

## Running the service

Go to the `infrastructure`-folder with
```bash
cd infrastructure
```

### Using Vagrant
```bash
vagrant up
```

### Run locally with Docker

Pull the Docker Image
```bash
docker pull ghcr.io/hfxbse/dhbw-devops/server:latest
```

Run the Docker Image
```bash
docker run -p 8080:8080 ghcr.io/hfxbse/dhbw-devops/server:latest
```

### Run locally with Docker Compose
```bash
docker compose up
```