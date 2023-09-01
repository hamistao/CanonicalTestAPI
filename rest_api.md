# Virtual Library REST API

All communication between the Virtual Library servers and its clients happens using a RESTful API over HTTP.

## API versioning

This API has only one version(v1.0).

## Return values

There are two standard return types:

* Standard return value
* Error

### Standard return value

For a standard synchronous operation, the following JSON object is returned:

```js
{
    "status": "Success",
    "status_code": 200,
    "data": {}
}
```

HTTP code must be 200.

### Error

There are various situations in which something may immediately go
wrong, in those cases, the following return value is used:

```js
{
    "error": "Error message"
}
```

HTTP code must be one of of 400, 404 or 500.

## Query Filtering

To filter your results on GET "/books" requests, filter is implemented for id, collections, genre, title, author, publish date and edition.

There is no default value for filter which means that all results found will
be returned. The following is the language used for the filter argument:

    ?fieldfilter=field_name&fieldfilter=other_field_name

The language follows the OData conventions for structuring REST API filtering
logic.
Values with spaces can be surrounded with quotes. Nesting filtering is also supported.
For instance, to filter on a field in a configuration you would pass:

    ?filter=config.field_name

For filtering on publish date you would use 'from' and 'to'. One can be used independently on the other:

    ?from=book.field_name&to=book.field_name

Here are a few GET query examples of the different filtering methods mentioned above:

    host:port/books?id=id

    host:port/books?authorFilter=author&from=06-10-22&to=09-01-23

## Database

The Virtual Library API supports only PostgreSQL database. Running 
```console
docker run --rm --name postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=canonical_api -d -p 5432:5432postgres
docker cp database_structure.sql postgres:/database_structure.sql
docker exec -it postgres bash
psql -U postgres -d canonical_api -a -f /database_structure.sql
```
you can create a container running an instance of PostgreSQL compatible with the standard `config.yml`. Run `docker stop postgres` to stop it.

## Testing

The testing is limited but represents what should be done for the entire API.

## PATCH

The Virtual Library API supports only PATCH methods to alter entities in the database.

## Methods

### POST /book

Create a new book instance.

### GET /books

Makes a query on the book database based on passed filters.

### PATCH /books/{id}

Changes any field in a specific book instance.

### DELETE /books/{id}

Removes a book from the database.

### POST /collection

Create a new collection instance.

### PATCH /collection/{name}

Changes any field in a specific collection instance.

### DELETE /collection/{name}

Removes a collection from the database.

### GET /collections

List all collections in the database.

### POST /collection/{name}/{id}

Adds a book with id {id} to a collection with name {name}. Is defined as DELETE because it creates a 'CollectionBooks' instance in the database.

### DELETE /collection/{name}/{id}

Removes a book with id {id} from a collection with name {name}. Is defined as DELETE because it deletes a 'CollectionBooks' instance from the database.
