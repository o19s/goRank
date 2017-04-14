var GORANK = (function () {
	var my = {};

	my.log_search = function(search, item) {
        var data = {
            search: search,
            item: item
        }
        // construct an HTTP request
        var xhr = new XMLHttpRequest();
        xhr.open("POST", "http://click.labs.o19s.com/events", true);
        xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');

        // send the collected data as JSON
        xhr.send(JSON.stringify(data));
    }

	return my;
}());