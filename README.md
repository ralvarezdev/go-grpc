# go-grpc

gRPC interceptors and helpers for Go projects: server and client authentication interceptors (JWT and API key), metadata helpers, error-detail generation, a request validation service and an error-handling interceptor. Requires Go 1.25.1 (per `go.mod`).

## Installation

```bash
go get github.com/ralvarezdev/go-grpc
```

Direct dependencies include `google.golang.org/grpc`, `connectrpc.com/connect`, `genproto/googleapis/rpc`, and `go-api-key`, `go-flags`, `go-jwt`, `go-reflect`, `go-validator` from `github.com/ralvarezdev`.

## Packages

- **`gogrpc`** (root) — `ErrorDetailsGenerator` and `NewDefaultErrorDetailsGenerator(logger)` (builds `errdetails.BadRequest` field violations), shared constants and errors.
- **`metadata`** — get/set/clear helpers for bearer, authorization, gcloud authorization, refresh and access tokens on `metadata.MD` and incoming/outgoing contexts.
- **`status`** — `ExtractErrorFromStatus(mode, err)`.
- **`server/interceptor/auth/jwt`** — `NewInterceptor(validator, interceptions)` and `Authenticate()`, a unary server interceptor validating JWTs for the configured full method names.
- **`server/interceptor/auth/apikey`** — `NewInterceptor(apiKeyService, methodsToIntercept)` and `Authenticate()`.
- **`server/interceptor/errorhandler`** — `NewInterceptor(modeFlag, logger)`.
- **`server/validator`** — `Service` interface for field validations (email, username, birthdate) and `NewService`.
- **`server/context`** — `GetClientIP(ctx)`.
- **`client/interceptor/auth/{apikey,gcloud}`** — unary client interceptors adding an API key or gcloud authorization.
- **`client/interceptor/auth/verifier/{apikey,jwt}`** — client-side verifier interceptors.
- **`client/interceptor/context/outgoing`** — outgoing context handling.
- **`client/net/http`** — `SetOutgoingCtxMetadataAuthorizationToken` for HTTP-originated calls.

## Usage

```go
// jwtValidator: a go-jwt validator; interceptions maps gRPC full method
// names to the token kind required for that method.
i, err := jwtinterceptor.NewInterceptor(jwtValidator, interceptions)
if err != nil {
    return err
}
s := grpc.NewServer(grpc.UnaryInterceptor(i.Authenticate()))
```

Methods missing from `interceptions` pass through without authentication.

## Development

```bash
go build ./...
go vet ./...
```

There are no tests.

## License

GNU General Public License v3.0 (see [LICENSE](LICENSE)).
