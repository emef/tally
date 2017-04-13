(function() {
  var charts = [];

  charts.push(Chart.newChart("reddit_new_links_frontpage", ""));

  var dateRange = Bar.getDateRange();
  charts.forEach(function(chart) {
    chart.fetchAndUpdate(dateRange[0], dateRange[1]);
  });

  Bar.registerCallback(function(startDate, endDate) {
    charts.forEach(function(chart) {
      chart.fetchAndUpdate(startDate, endDate);
    });
  });

})();
