# Submission

The project hosted in heroku free tier which can be see at https://secure-springs-22460.herokuapp.com/

## Changes 

- Show Chapter Title in search result
- Show Line Number in search result
- Fix index-out-bound error if any result show less/more than character 250 
- Tidy up html page (remove comma between row, add padding, add border)
- Tidy up project layout
- Add unit testing and benchmark
- Add makefile for build shortcut

## Metadata

We use the metadata to store more information to improve the search process. Currently, the metadata only contains a chapter list to identify and mapping chapter character index in the book.
```json
{
    "chapters":[]
}
```

## Further development

- pagination (with the current implementation, the pagination implementation is better on the Frontend side since the backend has the same cost to either return full or paged result)
- generate html version of the book (with text highlight and jump to specific line)
- high availability
  - caching
  - move book data to database/search-engine (stateless)
- production support
  - add observability (logger, request id, etc)
  - health-check
- security
  - rate limit

## Makefile

Run the project in local machine
```bash
make 
```

Run the test
```bash
make test
```

Run the benchmarking
```bash
make benchmarking
```

## Project Layout

- `internal/app` the golang source
- `data` the data which is the text file and its metadata
- `api/rest_api` http client the api ([rest-client vscode plugin](https://marketplace.visualstudio.com/items?itemName=humao.rest-client))

## Third Party Library

- `github.com/labstack/echo/v4 v4.1.17` web framework 
- `github.com/kelseyhightower/envconfig` environment variable config using struct
- `github.com/stretchr/testify v1.7.0` helper for unit testing

