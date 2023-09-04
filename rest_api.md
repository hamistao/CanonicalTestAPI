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

HTTP code must be 200 or 201.

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

    host:port/books?authorFilter=Author&from=06-10-22&to=09-01-23

    host:port/books?authorFilter=Other%20Author&titleFilter=Sample%20Book%20Title

## Database

The Virtual Library API only supports PostgreSQL database. if you have Docker installed, you can create a container running an instance of PostgreSQL compatible with the standard `config.yml` by running:
```console
docker run --rm --name postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=canonical_api -v ./postgresql.conf:/etc/postgresql/postgresql.conf -d -p 5432:5432 postgres
docker cp database_structure.sql postgres:/database_structure.sql
docker exec -it postgres bash
psql -U postgres -d canonical_api -a -f /database_structure.sql # inside the container's shell
```
Additionally, you can populate it for testing purposes with populate.sql by running:
```console
docker cp populate.sql postgres:/populate.sql
docker exec -it postgres bash
psql -U postgres -d canonical_api -a -f /populate.sql # inside the container's shell
```

. Run `docker stop postgres` to stop it.

## Testing

The testing is limited but represents what should be done for the entire API.

## PATCH

The Virtual Library API supports only PATCH methods to alter entities in the database.

## Methods

### POST /book

Create a new book instance.

```console
$curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"title":"Sample Book Title", "author":"Sample Author", "genre":"Fiction", "publisher":"Sample Publisher"}' \
  http://localhost:8080/book
{"data":"431a68ea-f537-4597-82a8-8903be2ca597","status":"created"}
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"title":"", "author":"Other Author", "genre":"Horror", "publisher":"Sample Publisher", "pub_date":"10-12-2002"}' \
  http://localhost:8080/book
{"error":"book must have a title"}
```

### GET /books

Makes a query on the book database based on passed filters.

```console
$curl -X GET "http://localhost:8080/books"
[{"id":"431a68ea-f537-4597-82a8-8903be2ca597","title":"Sample Book Title","description":"","author":"Sample Author","edition":"","genre":"Fiction","publisher":"Sample Publisher","pub_date":""},{"id":"9923475d-aafb-4168-98fa-a1cf81ea7520","title":"Sample Book Title 2","description":"","author":"Sample Author","edition":"","genre":"Fiction","publisher":"Sample Publisher","pub_date":"10-07-2002"},{"id":"901d72e4-68a4-42c5-a067-40d69f91d56d","title":"Other Book Title","description":"","author":"Other Author","edition":"","genre":"Horror","publisher":"Sample Publisher","pub_date":"10-12-2002"},{"id":"e447d6e2-666e-42c3-9078-7e6b3e13a876","title":"Other Book Title 2","description":"","author":"Other Author","edition":"","genre":"Horror","publisher":"Sample Publisher","pub_date":"10-12-2001"},{"id":"20c76a62-9ddc-4757-9185-841b61f8b1cf","title":"titulo","description":"","author":"Other Author","edition":"","genre":"","publisher":"Sample Publisher","pub_date":"10-12-2002"},{"id":"7d1fd740-1bc6-4afc-adbe-dd5ba7729ec5","title":"Other Book Title 2","description":"","author":"Other Author","edition":"","genre":"Horror","publisher":"Sample Publisher","pub_date":"10-12-2001"},{"id":"81e715a2-7d68-4031-b134-88dd2b4f2e7c","title":"Wrong Book","description":"","author":"Other Author","edition":"","genre":"Essay","publisher":"Sample Publisher","pub_date":""}]

$curl -X GET "http://localhost:8080/books?authorFilter=Sample%20Author"
[{"id":"431a68ea-f537-4597-82a8-8903be2ca597","title":"Sample Book Title","description":"","author":"Sample Author","edition":"","genre":"Fiction","publisher":"Sample Publisher","pub_date":""},{"id":"9923475d-aafb-4168-98fa-a1cf81ea7520","title":"Sample Book Title 2","description":"","author":"Sample Author","edition":"","genre":"Fiction","publisher":"Sample Publisher","pub_date":"10-07-2002"}]
```

### GET /books/{id}

Get a specific book instance by its ID.

```console
$curl -X GET "http://localhost:8080/book/9923475d-aafb-4168-98fa-a1cf81ea7520"
{"data":{"id":"9923475d-aafb-4168-98fa-a1cf81ea7520","title":"Sample Book Title","description":"","author":"Sample Author","edition":"","genre":"Fiction","publisher":"Sample Publisher","pub_date":"10-07-2002"},"status":"sucess"}
```

### PATCH /books/{id}

Changes any field in a specific book instance.

```console
$curl -X PATCH \
  -H "Content-Type: application/json" \
  -d '{"title": "Right Title"}' \
  http://localhost:8080/book/81e715a2-7d68-4031-b134-88dd2b4f2e7c
{"data":{"id":"","title":"Right Title","description":"","author":"","edition":"","genre":"","publisher":"","pub_date":""},"status":"success"}
```

### DELETE /books/{id}

Removes a book from the database.

```console
$curl -X DELETE "http://localhost:8080/book/20c76a62-9ddc-4757-9185-841b61f8b1cf"
{"status":"success"}
$curl -X DELETE "http://localhost:8080/book/20c76a62-9ddc-4757-9185-841b61f8b1cf"
{"error":"Book doesn't exist"}
```

### POST /collection

Create a new collection instance.

```console
$curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"name":"colecao"}' \
  http://localhost:8080/collection
{"status":"collection created"}
```

### PATCH /collection/{name}

Changes the description of a specific collection instance.

```console
$curl -X PATCH \
  -H "Content-Type: application/json" \
  -d '{"description": "a collection of books"}' \
  http://localhost:8080/collection/colecao
{"data":{"name":"","description":"a collection of books"},"status":"success"}
```

### DELETE /collection/{name}

Removes a collection from the database.

```console
$curl -X DELETE "http://localhost:8080/collection/colecao"
{"status":"Collection deleted successfully"}
```

### GET /collections

List all collections in the database.

```console
$curl -X GET "http://localhost:8080/collections"
{"data":[{"name":"colecao","description":"a collection of books"}],"status":"created"}
```

### POST /collection/{name}/{id}

Adds a book with id {id} to a collection with name {name}. Is defined as DELETE because it creates a 'CollectionBooks' instance in the database.

```console
$curl -X POST "http://localhost:8080/collection/colecao/431a68ea-f537-4597-82a8-8903be2ca597"
{"status":"success"}
$curl -X POST "http://localhost:8080/collection/colecao/431a68ea-f537-4597-82a8-8903be2ca597"
{"error":"pq: duplicate key value violates unique constraint \"collectionbook_pkey\""}
```

### DELETE /collection/{name}/{id}

Removes a book with id {id} from a collection with name {name}. Is defined as DELETE because it deletes a 'CollectionBooks' instance from the database.

```console
$curl -X DELETE "http://localhost:8080/collection/colecao/431a68ea-f537-4597-82a8-8903be2ca597"
{"status":"success"}
$curl -X DELETE "http://localhost:8080/collection/colecao/431a68ea-f537-4597-82a8-8903be2ca597"
{"error":"book isn't a part of this collection"}
```
