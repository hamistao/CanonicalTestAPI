# User Experience

The Virtual Library CLI allows the user to add, manage and query books on the Virtual Library platform. For that the following commands and options are offered:

## Commands

### vl_add

Add a book or list of books to the Virtual Library. Must be used with a book title(-t) or the path to a csv file with a list of books to be added.

#### Usage

vl_add (-t/--title book_title | -f/--file csv_file_path) [OPTIONS]

#### Options:

**-t, --title:** The title of the book being added to the system between quotes or double quotes. This parameter can include the official title or an informal name for identification. Obligatory parameter.

**-f, --file:** The path to a csv file containing information about the books being added.

**-a, --author:** The author of the book being added to the system.

**-g, --genre:** The genre of the book being added to the system.

**-T, --pub-date:** The publication date of the book being added to the system.

**-e, --edition:** The edition of the book being added to the system.

**-d, --description:** The description of the book being added to the system.

**-p, --publisher** The publisher of the book being added to the system.

#### Examples:

```console
$vl_add -t "Sample Book" -a "Sample Author" -g "Fiction" -p "Sample Publisher" -T "10-06-2002"
af0eac3e-6bf6-4548-a177-32e335891d43

# equivalent with csv
$echo "Sample_Book_Title,,Sample_Author,,Fiction,Sample_Publisher,10-06-2002" > ./book.csv

$vl_add -f library.csv
```

### vl_remove

Removes books to the Virtual Library. Automatically removes it from all the collections it is a part of.

#### Usage

vl_remove (-t/--title book_title | -i/--id book_id)

#### Options:

**-t, --title:** The name of the book being removed from the system.

**-i** The id of the book being removed from the system.

#### Examples:

```console
$vl_remove -t "Good Book"
there are many books with that title:
9ef46af8-1265-47b4-a4e4-0d1dfc011358, Good Book, John Doe
397e6802-85d7-4b4b-9fbd-ec646c32c5eb, Good Book, Mary Doe

$vl_remove -i 9ef46af8-1265-47b4-a4e4-0d1dfc011358
```

### vl_edit
Changes selected fields of a specific book in the Virtual Library.

#### Usage

vl_edit (-t/--title book_title | -i/--id book_id) [OPTIONS]

#### Options:

**-t, --title:** The current title of the book.

**-i:** The id of the book.

**-N, --new-title:** The new title of the book.

**-a, --author:** The new author of the book.

**-g, --genre:** The new genre of the book.

**-T, --pub-date:** The new publication date of the book.

**-e, --edition:** The new edition of the book.

**-d, --description:** The new description of the book.

**-p, --publisher** The new publisher of the book.

#### Examples:

```console
$vl_edit -t "Good Book" -a "Jane Doe" -g "Fiction"
there are many books with that title:
9ef46af8-1265-47b4-a4e4-0d1dfc011358, Good Book, John Doe
397e6802-85d7-4b4b-9fbd-ec646c32c5eb, Good Book, Mary Doe

vl_edit -i 9ef46af8-1265-47b4-a4e4-0d1dfc011358 -a "Jane Doe" -g "Fiction"
```

### vl_query

Makes queries within the Virtual Library. The queries can use multiple filters and can include either books or collections. The output has format: 
title | description | author | edition | genre | publisher | published_date

#### Usage

vl_query [OPTIONS]

#### Options:

**--from:** Filter books published after a certain date. The date format is mm-dd-yyyy.

**--to:** Filter books published before a certain date. The date format is mm-dd-yyyy.

**-a, --author:** Filter query by a specific author.

**-c, --collection:** Filter query by a specific collection.

**-g, --genre:** Filter query by a specific genre.

**-p, --publisher** Filter query by a specific publisher.

**-m, --max:** Set max items returned by query.

**-o** Set file for query output in .csv format.

#### Examples:

```console
$vl_add -t "Sample Book" -a "Sample Author" -T "10-06-2002"
$vl_add -t "Other Book" -a "Sample Author" -T "05-02-1999"
$vl_add -t "One More Book" -a "Other Author" -T "01-03-2002"

$vl_query -from 01-01-2000 -a "Sample Author"
Sample Book | | Sample Author | | | | 10-06-2002
```

### vl_collect new

Creates a new empty collection. If the name provided already exists the collection won't be created.

#### Usage

vl_collect new (-c/--collection collection_name) [OPTIONS]

#### Options:

**-c, --collection:** Name of the new collection. Obligatory parameter.

**-d, --description** Description of the collection.

#### Examples:

```console
$vl_collect new -c "Favorite Books" -d "My favorite books"

$vl_collect new -c "Favorite Books" -d "Her favorite books"
there is already a collection named "Favorite Books"
```

### vl_collect remove

Removes a book from a collection.

#### Usage

vl_collect new (-c/--collection collection_name -t/--title book_title)

#### Options:

**-c, --collection:** Name of the desired collection.

**-t, --title:** Title of the book being removed from the collection.

#### Examples:

```console
$vl_collect remove -c "Favorite Books" -t "Bad Book"
Favorite Books doesn't contain Bad Book

$vl_collect remove -c "Favorite Boks" -t "Bad Book"
collection "Favorite Boks" doesn't exist
```

### vl_collect edit

Sets a new name and/or a new description for a collection.

#### Usage

vl_collect new -c/--collection collection_name [OPTIONS]

#### Options:

**-c, --collection:** Actual name of the desired collection. Obligatory parameter.

**-n, --name:** New name of the desired collection.

**-d, --description** New description of the desired collection.

#### Examples:

```console
$vl_collect edit -c "Favorite Books" -n "Old Favorite Books" -d "My old favorite books"

$vl_collect remove -c "Old Favorite Boks" -n "Older Favorite Books"
collection "Favorite Boks" doesn't exist
```

### vl_collect list

Lists all existing collections.

#### Usage

vl_collect list

#### Examples:

```console
$vl_collect list -c "Favorite Books" -n "Old Favorite Books" -d "My old favorite books"
Favorite Books - My favorite books - "Book 1", "Book 2"
Old Favorite Books - My old favorite books - "Old Book 1", "Old Book 2"
```

### vl_collect add

Add a book or list of books into a collection.

#### Usage

vl_collect add -c/--collection collection_name (-t/--title book_title | -f/--file csv_file_path) [OPTIONS]

#### Options:

**-c, --collection:** Name of the desired collection. Obligatory parameter.

**-t, --title:** Title of the book being added to the collection.

**-b, --book:** Id of the book being added to the collection.

**-f, --file:** The path to a csv file containing information about the books being added to the collection.

**-F, --force** Flag that detemermines the creation of a collection with the provided name if one does not exist.

#### Examples:

```console
$vl_collect add -c "Favorite Books" -t "Good Book"
there are many books with that title:
9ef46af8-1265-47b4-a4e4-0d1dfc011358, Good Book, John Doe
397e6802-85d7-4b4b-9fbd-ec646c32c5eb, Good Book, Mary Doe

$vl_collect add -c "Favorite Books" -b 397e6802-85d7-4b4b-9fbd-ec646c32c5eb

$vl_collect add -c "Favorite Boks" -b 9ef46af8-1265-47b4-a4e4-0d1dfc011358
collection "Favorite Boks" doesn't exist

$vl_collect add -c "Favorite Boks" -b 9ef46af8-1265-47b4-a4e4-0d1dfc011358 -F
```
