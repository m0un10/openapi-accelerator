# openapi-accelerator

This is a tactical tool for creating a skeleton for CRUD-based OpenAPI Definitions in literally seconds.

Say we needed CRUD APIs for 5 different resources (users, groups, advocates, families and kittens), we could simply run 

```
./openapi-accelerator /admin/{users,groups} /community/{advocates,families,kitten}/catalog
```

And it will generate an OpenAPI specification skeleton like this

```
 families
 /home/families
   POST - List families
   GET - Add to families
 /home/families/{id}
   POST - Delete a specific family
   GET - Create a specific family
   PUT - Update a specific family
   DELETE - Delete a specific family
```

Notice that is actually determined the tags for grouping automatically and generated the descriptions with awareness of plural and singular words.

## Building the tool

Nothing special, just the standard `go build`.