var Bar = (function() {
  var startDate = new Date(new Date().setDate(new Date().getDate()-1)),
      endDate = new Date(),
      rangeText = document.getElementById('range-text'),
      registeredCallbacks = [],
      rangeUpdatedCallback = function() {
        console.log('range updated, redraw charts');
        registeredCallbacks.forEach(function(cb) {
          cb(startDate, endDate);
        });
      },
      updateRangeText = function() {
        rangeText.value = startDate + ' - ' + endDate;
      },
      updateStartDate = function() {
        startPicker.setStartRange(startDate);
        endPicker.setStartRange(startDate);
        updateRangeText();
      },
      updateEndDate = function() {
        startPicker.setEndRange(endDate);
        endPicker.setEndRange(endDate);
        if (endDate < startDate) {
          var tmp = startDate;
          startDate = endDate;
          endDate = tmp;
          updateStartDate();
          return updateEndDate();
        }
        updateRangeText();
      },
      startPicker = new Pikaday({
        numberOfMonths: 2,
        field: document.getElementById('start'),
        trigger: document.getElementById('custom-range'),
        setDefaultDate: true,
        defaultDate: startDate,
        minDate: new Date(2000, 1, 1),
        maxDate: new Date(2020, 12, 31),
        incrementMinuteBy: 15,
        onSelect: function() {
          startDate = this.getDate();
          updateStartDate();
          endPicker.show()
        }
      }),
      endPicker = new Pikaday({
        numberOfMonths: 2,
        field: document.getElementById('range-text'),
        setDefaultDate: true,
        defaultDate: endDate,
        minDate: new Date(2000, 1, 1),
        maxDate: new Date(2020, 12, 31),
        incrementMinuteBy: 15,
        onSelect: function() {
          endDate = this.getDate();
          updateEndDate();
          rangeUpdatedCallback();
        }
      });

  updateStartDate();
  updateEndDate();
  rangeUpdatedCallback();

  return {
    getDateRange: function() {
      return [startDate, endDate];
    },
    registerCallback: function(cb) {
      registeredCallbacks.push(cb);
    }
  };
})();
