const VALLIMIT = 200
const DECIMAL_POINTS_ZOND = 6
const DECIMAL_POINTS_CURRENCY = 3
var csrfToken = ""
var currency = ""
var subsTable = null
// let validators = []

function create_typeahead(input_container) {
  var timeWait = 0
  var debounce = function (context, func) {
    var timeout, result

    return function () {
      var args = arguments,
        later = function () {
          timeout = null
          result = func.apply(context, args)
        }
      clearTimeout(timeout)
      timeout = setTimeout(later, timeWait)
      if (!timeout) {
        result = func.apply(context, args)
      }
      return result
    }
  }

  var bhValidators = new Bloodhound({
    datumTokenizer: Bloodhound.tokenizers.whitespace,
    queryTokenizer: Bloodhound.tokenizers.whitespace,
    identify: function (obj) {
      return obj.index
    },
    remote: {
      url: "/search/indexed_validators/%QUERY",
      // use prepare hook to modify the rateLimitWait parameter on input changes
      // NOTE: we only need to do this for the first function because testing showed that queries are executed/queued in order
      // No need to update `timeWait` multiple times.
      prepare: function (_, settings) {
        var cur_query = $(input_container).val()
        timeWait = 4000 - Math.min(cur_query.length, 5) * 500
        // "wildcard" can't be used anymore, need to set query wildcard ourselves now
        settings.url = settings.url.replace("%QUERY", encodeURIComponent(cur_query))
        return settings
      },
    },
  })
  bhValidators.remote.transport._get = debounce(bhValidators.remote.transport, bhValidators.remote.transport._get)
  var bhExecutionAddresses = new Bloodhound({
    datumTokenizer: Bloodhound.tokenizers.whitespace,
    queryTokenizer: Bloodhound.tokenizers.whitespace,
    identify: function (obj) {
      return obj.execution_address
    },
    remote: {
      url: "/search/indexed_validators_by_execution_addresses/%QUERY",
      wildcard: "%QUERY",
    },
  })
  bhExecutionAddresses.remote.transport._get = debounce(bhExecutionAddresses.remote.transport, bhExecutionAddresses.remote.transport._get)
  var bhName = new Bloodhound({
    datumTokenizer: Bloodhound.tokenizers.whitespace,
    queryTokenizer: Bloodhound.tokenizers.whitespace,
    identify: function (obj) {
      return obj.name
    },
    remote: {
      url: "/search/indexed_validators_by_name/%QUERY",
      wildcard: "%QUERY",
    },
  })
  bhName.remote.transport._get = debounce(bhName.remote.transport, bhName.remote.transport._get)
  var bhGraffiti = new Bloodhound({
    datumTokenizer: Bloodhound.tokenizers.whitespace,
    queryTokenizer: Bloodhound.tokenizers.whitespace,
    identify: function (obj) {
      return obj.graffiti
    },
    remote: {
      url: "/search/indexed_validators_by_graffiti/%QUERY",
      wildcard: "%QUERY",
    },
  })
  bhGraffiti.remote.transport._get = debounce(bhGraffiti.remote.transport, bhGraffiti.remote.transport._get)

  $(input_container).typeahead(
    {
      minLength: 1,
      highlight: true,
      hint: false,
      autoselect: false,
    },
    {
      limit: 5,
      name: "validators",
      source: bhValidators,
      display: "index",
      templates: {
        header: "<h3>Validators</h3>",
        suggestion: function (data) {
          return `<div class="text-monospace text-truncate high-contrast">${data.index}</div>`
        },
      },
    },
    {
      limit: 5,
      name: "addresses",
      source: bhExecutionAddresses,
      display: "address",
      templates: {
        header: "<h3>Validators by QRL Addresses</h3>",
        suggestion: function (data) {
          var len = data.validator_indices.length > VALLIMIT ? VALLIMIT + "+" : data.validator_indices.length
          return `<div class="text-monospace high-contrast" style="display:flex"><div class="text-truncate" style="flex:1 1 auto;">${data.execution_address}</div><div style="max-width:fit-content;white-space:nowrap;">${len}</div></div>`
        },
      },
    },
    {
      limit: 5,
      name: "graffiti",
      source: bhGraffiti,
      display: "graffiti",
      templates: {
        header: "<h3>Validators by Graffiti</h3>",
        suggestion: function (data) {
          var len = data.validator_indices.length > VALLIMIT ? VALLIMIT + "+" : data.validator_indices.length
          return `<div class="text-monospace high-contrast" style="display:flex"><div class="text-truncate" style="flex:1 1 auto;">${data.graffiti}</div><div style="max-width:fit-content;white-space:nowrap;">${len}</div></div>`
        },
      },
    },
    {
      limit: 5,
      name: "name",
      source: bhName,
      display: "name",
      templates: {
        header: "<h3>Validators by Name</h3>",
        suggestion: function (data) {
          var len = data.validator_indices.length > VALLIMIT ? VALLIMIT + "+" : data.validator_indices.length
          return `<div class="text-monospace high-contrast" style="display:flex"><div class="text-truncate" style="flex:1 1 auto;">${data.name}</div><div style="max-width:fit-content;white-space:nowrap;">${len}</div></div>`
        },
      },
    }
  )

  $(input_container).on("focus", function (event) {
    if (event.target.value !== "") {
      $(this).trigger($.Event("keydown", { keyCode: 40 }))
    }
  })
  $(input_container).on("input", function () {
    $(".tt-suggestion").first().addClass("tt-cursor")
  })
  $(input_container).on("blur", function () {
    $(this).val("")
    $(input_container).typeahead("val", "")
  })
  $(input_container).on("typeahead:select", function (ev, sug) {
    validators = $("#validator-index-view").val().split(",")
    if (sug.validator_indices) {
      validators = validators.concat(sug.validator_indices)
    } else {
      validators.push(sug.index)
    }

    validators = Array.from(new Set(validators))
    if (validators.length > VALLIMIT) {
      validators = validators.slice(0, VALLIMIT)
      alert(`No more than ${VALLIMIT} validators are allowed`)
    }

    $("#validator-index-view").val(validators)
    if ($("#validator-index-view").val().charAt(0) === ",") {
      $("#validator-index-view").val($("#validator-index-view").val().slice(1))
    }
    $(input_container).typeahead("val", "")
  })
}

function getValidatorQueryString() {
  return window.location.href.slice(window.location.href.indexOf("?"), window.location.href.length)
}

function hideSpinner() {
  $("#loading-div").addClass("d-none")
  $("#loading-div").removeClass("d-flex")
}

function showSpinner() {
  $("#loading-div").removeClass("d-none")
  $("#loading-div").addClass("d-flex")
}

function showTable(data) {
  $("#tax-table").DataTable({
    searchDelay: 0,
    processing: true,
    serverSide: false,
    ordering: true,
    searching: true,
    pagingType: "full_numbers",
    pageLength: 100,
    lengthChange: false,
    data: data.history,
    dom: "Bfrtip",
    buttons: ["copyHtml5", "excelHtml5", "csvHtml5", "pdfHtml5"],
    drawCallback: function (settings) {
      $("#form-div").removeClass("d-flex").addClass("d-none")
      $("#table-div").removeClass("d-none")
      $("#subscriptions-div").addClass("d-none")
      $("#total-income-quanta-span").html("Quanta " + data.total_quanta)
      $("#totals-div").removeClass("d-none")
      $(".dt-button").addClass("ml-2 ")
      hideSpinner()
    },
    order: [[0, "desc"]],
    language: {
      searchPlaceholder: "Enter Date",
    },
    columnDefs: [
      {
        targets: 0,
        data: "0",
        orderable: true,
        render: function (data, type, row, meta) {
          if (type === "filter" || type === "display") return data
          return moment(data).unix()
        },
      },
      {
        targets: 1,
        data: "1",
        orderable: true,
        render: function (data, type, row, meta) {
          // return (parseFloat(data).toFixed(DECIMAL_POINTS_ZOND))
          return data
        },
      },
      {
        targets: 2,
        data: "2",
        orderable: true,
        render: function (data, type, row, meta) {
          // return (parseFloat(data).toFixed(DECIMAL_POINTS_ZOND))
          return data
        },
      },
      // {
      //   targets: 3,
      //   data: "3",
      //   orderable: false,
      //   render: function (data, type, row, meta) {
      //     // return `${currency} ${addCommas(parseFloat(data).toFixed(DECIMAL_POINTS_CURRENCY))}`
      //     return data
      //   },
      // },
      // {
      //   targets: 4,
      //   data: "4",
      //   orderable: false,
      //   render: function (data, type, row, meta) {
      //     //    return `${currency} ${addCommas(parseFloat(data).toFixed(DECIMAL_POINTS_CURRENCY))}`
      //     return data
      //   },
      //   // }, {
      //   //     targets: 5,
      //   //     data: '5',
      //   //     "orderable": false,
      //   //     visible: false,
      //   //     render: function (data, type, row, meta) {
      //   //         return data.toUpperCase()
      //   //     }
      // },
    ],
  })
}

function updateSubscriptionTable(data, container) {
  if (data.data.length === 0) {
    $("#subscriptions-div").addClass("d-none")
    hideSpinner()
    return
  }
  if (subsTable !== null) {
    // subsTable.clear()
    // subsTable.empty()
    // subsTable.destroy()
    $("#" + container)
      .DataTable()
      .clear()
      .destroy()
    // $('#'+container+" tbody").empty()
    // $('#'+container+" thead").empty()
  }
  subsTable = $("#" + container).DataTable({
    searchDelay: 0,
    processing: true,
    serverSide: false,
    ordering: true,
    searching: true,
    pagingType: "full_numbers",
    pageLength: 100,
    lengthChange: false,
    data: data.data,
    drawCallback: function (settings) {
      $("#subscriptions-table-art").removeClass("d-flex").addClass("d-none")
      $("#subscriptions-table-div").removeClass("invisible")
      $("#subscriptions-div").removeClass("d-none")
      $("#subs-header").html(`Subscriptions (${data.count}/5)`)
      hideSpinner()
    },
    language: {
      searchPlaceholder: "Enter Date, Currency",
    },
    columnDefs: [
      {
        targets: 0,
        data: "0",
        orderable: true,
        render: function (data, type, row, meta) {
          if (type === "filter" || type === "display") {
            let date = data.split(" ")
            if (date.length >= 1) {
              return `${date[0]}`
            }
            return data
          }
          return moment(data).unix()
        },
      },
      {
        targets: 1,
        data: "1",
        orderable: true,
        render: function (data, type, row, meta) {
          return data.toUpperCase()
        },
      },
      {
        targets: 2,
        data: "2",
        orderable: false,
        render: function (data, type, row, meta) {
          if (type === "display") {
            l = data.split(",")
            l.sort((a, b) => parseInt(a) - parseInt(b))
            data = ""
            for (i of l) {
              data += `<li style="flex: 1 0 8%; list-style-type : none;" class="p-1"><a href="/validator/${i}"><i class="fas fa-male mr-1"></i>${i}</a></li>`
            }
            return `<ul style="display: flex; flex-wrap: wrap; height: 50px; width: 98%; overflow: auto; background-color: rgba(0, 0, 0, 0);" class="nice-scroll text-dark pl-0 ml-0">${data}</ul>`
          }
          return data
        },
      },
      // {
      //   targets: 3,
      //   data: "3",
      //   orderable: false,
      //   render: function (data, type, row, meta) {
      //     downloadQueryUrl = `${window.location.origin}/rewards/hist/download?validators=${row[2]}&currency=${row[1]}&days=${moment().subtract(1, "month").startOf("month").unix()}-${moment().subtract(1, "month").endOf("month").unix()}`
      //     return `
      //                   <div class="d-flex justify-content-between align-item-center">
      //                       <i class="far fa-clone mr-2" style="cursor: pointer;" onClick='loadValInForm("${row[2]}")' data-toggle="tooltip" data-placement="top" title="Load validators in the form"></i>
      //                       <a href="${downloadQueryUrl}" download><i class="fas fa-file-download mr-2" style="cursor: pointer;" data-toggle="tooltip" data-placement="top" title="Download the last month report"></i></a>
      //                   </div>
      //                   `
      //   },
      // },
    ],
  })
}

function loadValInForm(val) {
  $("#validator-index-view").val(val.replace(/([a-zA-Z ])/g, ""))
}

$(document).ready(function () {
  if (document.getElementsByName("CsrfField")[0] !== undefined) {
    csrfToken = document.getElementsByName("CsrfField")[0].value
  }

  if (JSON.parse(localStorage.getItem("load_dashboard_validators"))) {
    const dashBoardValidators = JSON.parse(localStorage.getItem("dashboard_validators"))
    const dashBoardValidatorIndices = dashBoardValidators.filter((entry) => {
      return !(entry.startsWith("0x") && entry.length === 98)
    })

    $("#validator-index-view").val(dashBoardValidatorIndices)
    localStorage.setItem("load_dashboard_validators", false)
  }

  $("#validator-index-view").on("keyup", function () {
    $(this).val(
      $(this)
        .val()
        .replace(/([a-zA-Z ])/g, "")
    )
  })

  $("#days").val(`${moment().startOf("month").unix()}-${moment().unix()}`)

  $('input[id="datepicker"]').daterangepicker(
    {
      pens: "left",
      minDate: moment.unix(MIN_TIMESTAMP),
      maxDate: moment(),
      maxSpan: {
        days: 365,
      },
      ranges: {
        "This Month to date": [moment().startOf("month"), moment()],
        "Last Month to date": [moment().subtract(1, "month").startOf("month"), moment()],
        "This Year to date": [moment().startOf("year"), moment()],
      },
      locale: {
        format: "DD/MM/YYYY",
      },
      singleDatePicker: false,
      alwaysShowCalendars: false,
      startDate: moment().startOf("month"),
      endDate: moment(),
    },
    function (start, end, label) {
      // let end_d = moment()
      $("#days").val(`${moment(start).unix()}-${moment(end).unix()}`)
    }
  )

  create_typeahead(".typeahead-validators")
  let qry = getValidatorQueryString()
  // console.log(qry, qry.length)

  if (qry.length > 1) {
    fetch(`/rewards/hist${qry}`, {
      method: "GET",
    })
      .then((res) => {
        if (res.status !== 200) {
          alert("Request failed :(")
          hideSpinner()
        }
        res.json().then((data) => {
          showTable(data)
        })
      })
      .catch(() => {
        alert("Failed to fetch the data :(")
        hideSpinner()
      })
  } else {
    hideSpinner()
  }
})
