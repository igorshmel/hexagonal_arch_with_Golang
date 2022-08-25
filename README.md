# hex-arch-template
This is a full-stack Go template for Hexagonal Architecture.
Includes backend with different kinds of API, a few flavors of DB, Docker and docker-compose.
Front end is React with TypeScript using Apollo client for GraphQL consumption as well as Auth0 for authentication.
There is a simple frontend with Vecty.
The structure is inspired by [hexArchGoGRPC](https://github.com/selikapro/hexArchGoGRPC) repo as well as the [video](https://t.co/QaN1cAzDmu?amp=1) itself.

## Basic idea
The basic idea is to create business/application logic in the Application/Core layers.
Declare ports interfaces and arrange them to the left and right side.
The right side is the controlled systems by the Application while the left side is the controller of the Application.
Such things as APIs go to the left and DB, FS and external APIs/systems go to the right.
One important rule is that the core must not depend on Application, while Application must not depend on APIs or DBs.
The communication with external layers should go through ports.