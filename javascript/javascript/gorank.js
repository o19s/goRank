var GORANK = (function () {
	var my = {};

	my.log_search = function(search, item) {
        var data = {
            search: search,
            item: item
        }
        // construct an HTTP request
        var xhr = new XMLHttpRequest();
        xhr.open("POST", "http://localhost:8000/events", true);
        xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');

        // send the collected data as JSON
        xhr.send(JSON.stringify(data));
    }

    my.search_stats = function(search) {
        // construct an HTTP request
        var xhr = new XMLHttpRequest();
        xhr.open("GET", "http://localhost:8000/searches/" + search, false);
        xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');

        xhr.send();
        console.log(xhr.status);
        console.log(xhr.responseText);
        return(JSON.parse(xhr.responseText));
    }
	return my;
}());