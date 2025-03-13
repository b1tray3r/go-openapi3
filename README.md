# OpenAPI Restservice with Swagger-UI

This repository contains an almost minimal setup for developing a REST service using [Echo](https://echo.labstack.com/), [OpenAPI](https://www.openapis.org/) and [Swagger-UI](https://swagger.io/tools/swagger-ui/) integration.

The generation and build process is using [air](https://github.com/air-verse/air) for easy setup.

The OpenAPI defintion is defined with code using go in the `internal/openapi/definition.go` file.

During the build steps air is going to generate the `openapi3.json` file in the `embed/swagger` directory.
Additionally, the according server interfaces are generated with [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen)) into the `pkg/api/` folder.

The application will serve the Swagger-UI at `swagger/` endpoint.

The backend is implemented in the `internal/server/server.go`file.


## Setup

1. Clone the project

```sh
git clone https://github.com/b1tray3r/go-openapi3.git
```

2. Change the directory

```sh
cd go-openapi3
```

3. Install openapi-gen

```shell
go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
```

4. Install air

```sh
go install github.com/air-verse/air@latest
```

5. Launch

```sh
air
```

### Details

Many things are hardcoded on purpose to keep the repository as barebone as possible.

#### index.html

This contains the swagger-ui code and requires access to [unpkg](https://unpkg.com/) servers.

The file is requesting the generated `/swagger/openapi3.json` file to build the GUI.

#### .air.toml

Using air the complete build pipeline is executed on the fly.
When building the final application you have to execute all `pre_cmd` entries in the `.air.toml` configuration.

Please note that all `*.gen.go` files are ignored by git and air.
It is important for air to exclude this files, because generating these file(s) will cause air to loop indefinetly.

#### oapi-codegen.yml

This configuration is responsible to generate the api interfaces and types for echo.
I am not sure whether this is an optimal configuration, but it seems to work.

## Sources

- Thanks to [plutov](https://github.com/plutov/packagemain/tree/master/oapi-example) for his [Youtube](https://www.youtube.com/watch?v=87au30fl5e4) guide.

- Thanks to [MarioCarrion](https://github.com/MarioCarrion/todo-api-microservice-example/tree/074bbb9f4d0f79e5bced943c10c56013705969a9) for his [Youtube](https://www.youtube.com/watch?v=HwtOAc0M08o) guide.
