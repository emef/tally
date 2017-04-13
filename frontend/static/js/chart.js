Chart = (function() {
  var margin = {top: 20, right: 20, bottom: 30, left: 50};
  var parseTime = d3.timeParse("%d-%b-%y");

  return {
    newChart: function(name, source) {
      var name = name;
      var source = source;
      var div = document.createElement("div");
      div.className = "chart";
      document.body.appendChild(div);


      var svg = d3.select(div).append("svg:svg").attr("width", "100%").attr("height", "100%");

      var width = div.offsetWidth - margin.left - margin.right;
      var height = div.offsetHeight - margin.top - margin.bottom;
      var g = svg
          .append("g")
          .attr("transform", "translate(" + margin.left + "," + margin.top + ")");

      var x = d3.scaleTime()
          .rangeRound([0, width]);

      var y = d3.scaleLinear()
          .rangeRound([height, 0]);

      var line = d3.line()
          .x(function(d) { return x(d.date); })
          .y(function(d) { return y(d.sum); });

      function fetchAndUpdate(startDate, endDate) {
        var url =
            "/csv/" + name + "/" + source +
            "?startEpochMinute=" + Math.floor(startDate.getTime() / (1000 * 60)) +
            "&endEpochMinute=" + Math.floor(endDate.getTime() / (1000 * 60));
        console.log(url);
        d3.csv(url, function(d) {
          d.date = new Date(d.epochMinute * 60 * 1000);
          d.sum = +d.sum;
          return d;
        }, function(error, data) {
          if (error) throw error;

          x.domain(d3.extent(data, function(d) { return d.date; }));
          y.domain(d3.extent(data, function(d) { return d.sum; }));

          g.selectAll("*").remove();

          g.append("g")
            .attr("transform", "translate(0," + height + ")")
            .call(d3.axisBottom(x))
            .select(".domain")
            .remove();

          g.append("g")
            .call(d3.axisLeft(y))
            .append("text")
            .attr("fill", "#000")
            .attr("transform", "rotate(-90)")
            .attr("y", 6)
            .attr("dy", "0.71em")
            .attr("text-anchor", "end")
            .text("New Items");

          g.append("path")
            .datum(data)
            .attr("fill", "none")
            .attr("stroke", "steelblue")
            .attr("stroke-linejoin", "round")
            .attr("stroke-linecap", "round")
            .attr("stroke-width", 1.5)
            .attr("d", line);
        });
      }

      return {
        name: name,
        source: source,
        fetchAndUpdate: fetchAndUpdate
      };
    }
  };
})();
