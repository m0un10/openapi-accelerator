# openapi-accelerator

This is a tactical tool for creating a skeleton for CRUD-based OpenAPI Definitions in literally seconds.

Say we needed CRUD APIs for 5 different resources; users, groups, advocates, families and kittens...

We could simply run 

```
./openapi-accelerator /community/{advocates,families,kittens}/catalog /admin/{users,groups} 
```

And it will generate an OpenAPI specification skeleton like [this](sample-output.yml)

You'll notice that it determines the tags for grouping automatically and generates descriptions with awareness of plural and singular words.

## Building the tool

Nothing special, just the standard `go build`.

## Viewing the docs

Nothing special, just copy the output from the command and paste it at https://editor.swagger.io/


