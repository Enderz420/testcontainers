# Testcontainers PoC

At work we need to update our DevOps pipeline before the summer. In our backend we use Testcontainers to do unit testing.

Right now we have no way to ensure that the frontend and backend work together.

This is a Proof of Concept (more of just a playground) to experiment with how Testcontainers works within a Nuxt environment.

The packages used are Vite 8 with Vitest 4.1.7, Testcontainers 12, Nuxt Test Utils 4 and Typescript 6

The following scripts are added in webpage/package.json.

```json
    "test:int": "vitest --config ./vitest.config.ts --project integrations run",
    "test:e2e": "vitest --config ./vitest.config.ts --project e2e run",
    "test:nuxt": "vitest --config ./vitest.config.ts --project nuxt run",
```

The docker file used for testing is in /deployments.

Migrations are defined in /migrations.

This project also uses Github Actions to verify that the frontend CI build works.
