# jwt-kit

Simple fake idp and jwt generator for testing code that uses jwts.

### Quickstart

```bash
go build

# assumes workspace w/ gloo-mesh and default namespaces with workloads applied to mgmt-cluster
kubectl apply -f ./deploy

kubectl port-forward -n gloo-mesh-gateways $(kubectl get pods -n gloo-mesh-gateways -l app=istio-ingressgateway -oname) 8080:8080 &

curl localhost:8080/httpbin/ip -i -H "Authorization: Bearer $(./jwt-kit)"
```

should output

```
Handling connection for 8080
HTTP/1.1 200 OK
server: istio-envoy
date: Tue, 14 Mar 2023 15:52:44 GMT
content-type: application/json
content-length: 30
access-control-allow-origin: *
access-control-allow-credentials: true
x-envoy-upstream-service-time: 1

{
  "origin": "10.10.43.16"
}
```

### Usage

```
jwt-kit -h

Jwt-kit contains an embedded keypair used to sign jwts.

Public JWKS url: https://bensolo-io.github.io/jwt-kit/.well-known/jwks.json

Issuer name: https://bensolo-io.github.io/jwt-kit

Usage:
  jwt-kit [flags]

Flags:
  -a, --audiences stringArray   jwt audience (default [https://fake-resource.solo.io])
  -c, --claims stringArray      add jwt claims
  -e, --expires-in string       expires duration (uses https://pkg.go.dev/time#ParseDuration) (default "8766h")
  -h, --help                    help for jwt-kit
  -p, --pretty-print            pretty print the token
  -s, --scopes stringArray      add jwt scopes
  -u, --subject string          jwt subject (default "glooey@solo.io")
```

### Create a JWT

Create a new jwt with default values:

```bash
jwt-kit

eyJhbGciOiJSUzI1NiIsImtpZCI6ImZPcFdRT0pfc1NrQXp6djVOVTFqZnk5UDZuT3d6OXVkVWV3MjZmZXJZTXMiLCJ0eXAiOiJKV1QifQ.eyJhdWQiOlsiaHR0cHM6Ly9mYWtlLXJlc291cmNlLnNvbG8uaW8iXSwiYmVlcl9vZl90aGVfZGF5IjoiVHdvIEhlYXJ0ZWQgQWxlIiwiZXhwIjoxNjc4ODAyNjg2LCJpc3MiOiJodHRwczovL2JlbnNvbG8taW8uZ2l0aHViLmlvL2p3dC1raXQiLCJzY29wZXMiOltdLCJzdWIiOiJnbG9vZXlAc29sby5pbyJ9.MBlP2zLN4FANiyvRMXqsXtBmOfS8yOSZqEq6oRIWpnF8zt_JUIpVWHMbWlVGwDZ7dEqsup_LsAsvXXruLYKpxNLOZNqQFUPPQTmSMmPqwfe_tANsKEf8SkkNWV_emNW_cYnobe3QzsxAhG9Lg1xupN2jb8O97951mVIMavdImOwcvS5-xBD7ruT3WHX4w5lOoFnGLizuQO4lhAfFCwwdtx5jrhOADwhM4x_Lyd74poUpbyqtPWjQ-aslMbgSCcNwM6OHK9D0cgdTaGEhZg7KaooBvITb0DU46mXF-1vcWxkB7p2J7hLrQdndCsld6vdS2E0dl9ZB9hWi85VtIFVynQ
```

Pretty print to see the token details:

```bash
jwt-kit -p

{
  "Raw": "eyJhbG...",
  "Method": {
    "Name": "RS256",
    "Hash": 5
  },
  "Header": {
    "alg": "RS256",
    "kid": "fOpWQOJ_sSkAzzv5NU1jfy9P6nOwz9udUew26ferYMs",
    "typ": "JWT"
  },
  "Claims": {
    "aud": [
      "https://fake-resource.solo.io"
    ],
    "beer_of_the_day": "Westmalle Trappist Tripel",
    "exp": 1678802041,
    "iss": "https://bensolo-io.github.io/jwt-kit",
    "scopes": [],
    "sub": "glooey@solo.io"
  },
  "Signature": "KfF_...",
  "Valid": true
}
```

Add claims, subject, scopes, audiences, and set duration:

```bash
jwt-kit -e 24h -c my-claim=my-value -c color=orange -a https://some-resource -u ben@solo.io -s openid -s profile -p

{
  "Raw": "eyJh...",
  "Method": {
    "Name": "RS256",
    "Hash": 5
  },
  "Header": {
    "alg": "RS256",
    "kid": "fOpWQOJ_sSkAzzv5NU1jfy9P6nOwz9udUew26ferYMs",
    "typ": "JWT"
  },
  "Claims": {
    "aud": [
      "https://some-resource"
    ],
    "beer_of_the_day": "Trappistes Rochefort 8",
    "color": "orange",
    "exp": 1678888566,
    "iss": "https://bensolo-io.github.io/jwt-kit",
    "my-claim": "my-value",
    "scopes": [
      "openid",
      "profile"
    ],
    "sub": "ben@solo.io"
  },
  "Signature": "PYTNjUD...",
  "Valid": true
}
```

### Private Keys

Private keys are embedded in [./internal/idp/keys.go](./internal/idp/keys.go) ... it should go without saying that this idp should not be used for anything beyond local development!

### Public JWKS

Public JWKS is hosted at https://bensolo-io.github.io/jwt-kit/.well-known/jwks.json.

### TODO

Add a JSON claims file argument for complex / nested claims.