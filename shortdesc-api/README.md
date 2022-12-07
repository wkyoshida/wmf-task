# shortdesc-api

## Starting the API server:

```bash
$ go run .
```

**Note:** The server makes use of a [env.yaml](./env.yaml) config file.
The default config file stored in the repo configures the API
for the https://en.wikipedia.org/ instance (this can be configured).

## API Details

The API currently has only a single endpoint:

```
GET /shortdesc
```

### Query parameter:

`titles` string (Required)

The titles of MediaWiki pages. The title is case sensitive.

For the example, https://en.wikipedia.org/wiki/Abraham_Lincoln,
the title would be `Abraham_Lincoln`.

**Note:** The API reflects the same limit on the maximum number of `titles`
values in a single request allowed by the parameter of the same name for the
[MediaWiki API](https://www.mediawiki.org/wiki/API:Query#API_documentation), i.e. maximum is 50 values.

### Response schema

The API response follows this schema:

```json
{
    "pages": [
        {
            "pageid": "1234",
            "title": "Page_Title",
            "shortdescription": "Short description of the wiki page",
            "shortdescriptionraw": "{{Short description|Short description of the wiki page}}",
            "timestamp": "2012-12-12T12:12:12Z"
        },
        ...
    ]
}
```

For each page returned, the following fields may be populated:  

`pageid`: the identifier of the page in its given MediaWiki instance

`title`: the title of the page

`shortdescription`: the short description of the page

`shortdescriptionraw`: the raw Wikitext annotation of the short description of the page

`timestamp`: the timestamp of the last revision for the page

## Usage example

1. `titles` can include multiple pages

Making a request for the `Abraham_Lincoln` and `Yoshua_Bengio` pages in a single request:

```
GET /shortdesc?titles=Abraham_Lincoln&titles=Yoshua_Bengio
```

Receives the following response: 

```json
{
    "pages": [
        {
            "pageid": "307",
            "title": "Abraham_Lincoln",
            "shortdescription": "President of the United States from 1861 to 1865",
            "shortdescriptionraw": "{{Short description|President of the United States from 1861 to 1865}}",
            "timestamp": "2022-11-29T21:11:42Z"
        },
        {
            "pageid": "47749536",
            "title": "Yoshua_Bengio",
            "shortdescription": "Canadian computer scientist",
            "shortdescriptionraw": "{{Short description|Canadian computer scientist}}",
            "timestamp": "2022-11-13T20:09:37Z"
        }
    ]
}
```

2. Alternative use of `titles`

The below request receives the same response as the above example.

```
GET /shortdesc?titles=Abraham_Lincoln|Yoshua_Bengio
```

3. Request for invalid title returns nothing

Making a request for invalid title, `asdfasdf`:

```
GET /shortdesc?titles=Abraham_Lincoln&titles=asdfasdf
```

Does not receive an entry in the response:

```json
{
    "pages": [
        {
            "pageid": "307",
            "title": "Abraham_Lincoln",
            "shortdescription": "President of the United States from 1861 to 1865",
            "shortdescriptionraw": "{{Short description|President of the United States from 1861 to 1865}}",
            "timestamp": "2022-11-29T21:11:42Z"
        }
    ]
}
```
