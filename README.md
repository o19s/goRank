# Go, Rank!

The begining of a click tracker for your clicks.

## run with docker-compose

```bash
docker-compose up
```

Wait for the Elasticsearch instace to get fully loaded then create the 'gorank' collection by going here:

http://localhost:8000/init

### sending clicks

Take a look at the index.html in the main project folder for an example of sending click data to go-rank

### reading clicks

The url to read clicks is "/searches/:search_name", ex. http://localhost:8000/searches/spoon

## building and testing

### add some packages used to the vendor folder

```bash
dep ensure
```
