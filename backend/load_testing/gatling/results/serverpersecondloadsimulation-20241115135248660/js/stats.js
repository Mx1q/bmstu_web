var stats = {
    type: "GROUP",
name: "Global Information",
path: "",
pathFormatted: "group_missing-name-b06d1",
stats: {
    "name": "Global Information",
    "numberOfRequests": {
        "total": "6000",
        "ok": "6000",
        "ko": "0"
    },
    "minResponseTime": {
        "total": "0",
        "ok": "0",
        "ko": "-"
    },
    "maxResponseTime": {
        "total": "16663",
        "ok": "16663",
        "ko": "-"
    },
    "meanResponseTime": {
        "total": "7752",
        "ok": "7752",
        "ko": "-"
    },
    "standardDeviation": {
        "total": "4976",
        "ok": "4976",
        "ko": "-"
    },
    "percentiles1": {
        "total": "9040",
        "ok": "9040",
        "ko": "-"
    },
    "percentiles2": {
        "total": "11745",
        "ok": "11745",
        "ko": "-"
    },
    "percentiles3": {
        "total": "15081",
        "ok": "15081",
        "ko": "-"
    },
    "percentiles4": {
        "total": "16192",
        "ok": "16192",
        "ko": "-"
    },
    "group1": {
    "name": "t < 800 ms",
    "count": 910,
    "percentage": 15
},
    "group2": {
    "name": "800 ms < t < 1200 ms",
    "count": 136,
    "percentage": 2
},
    "group3": {
    "name": "t > 1200 ms",
    "count": 4954,
    "percentage": 83
},
    "group4": {
    "name": "failed",
    "count": 0,
    "percentage": 0
},
    "meanNumberOfRequestsPerSecond": {
        "total": "193.548",
        "ok": "193.548",
        "ko": "-"
    }
},
contents: {
"req_nginx-metrics-ff665": {
        type: "REQUEST",
        name: "Nginx Metrics",
path: "Nginx Metrics",
pathFormatted: "req_nginx-metrics-ff665",
stats: {
    "name": "Nginx Metrics",
    "numberOfRequests": {
        "total": "3000",
        "ok": "3000",
        "ko": "0"
    },
    "minResponseTime": {
        "total": "0",
        "ok": "0",
        "ko": "-"
    },
    "maxResponseTime": {
        "total": "16663",
        "ok": "16663",
        "ko": "-"
    },
    "meanResponseTime": {
        "total": "7869",
        "ok": "7869",
        "ko": "-"
    },
    "standardDeviation": {
        "total": "5102",
        "ok": "5102",
        "ko": "-"
    },
    "percentiles1": {
        "total": "8997",
        "ok": "8997",
        "ko": "-"
    },
    "percentiles2": {
        "total": "12066",
        "ok": "12066",
        "ko": "-"
    },
    "percentiles3": {
        "total": "15158",
        "ok": "15158",
        "ko": "-"
    },
    "percentiles4": {
        "total": "16197",
        "ok": "16197",
        "ko": "-"
    },
    "group1": {
    "name": "t < 800 ms",
    "count": 455,
    "percentage": 15
},
    "group2": {
    "name": "800 ms < t < 1200 ms",
    "count": 76,
    "percentage": 3
},
    "group3": {
    "name": "t > 1200 ms",
    "count": 2469,
    "percentage": 82
},
    "group4": {
    "name": "failed",
    "count": 0,
    "percentage": 0
},
    "meanNumberOfRequestsPerSecond": {
        "total": "96.774",
        "ok": "96.774",
        "ko": "-"
    }
}
    },"req_backend-metrics-aa2a0": {
        type: "REQUEST",
        name: "Backend Metrics",
path: "Backend Metrics",
pathFormatted: "req_backend-metrics-aa2a0",
stats: {
    "name": "Backend Metrics",
    "numberOfRequests": {
        "total": "3000",
        "ok": "3000",
        "ko": "0"
    },
    "minResponseTime": {
        "total": "0",
        "ok": "0",
        "ko": "-"
    },
    "maxResponseTime": {
        "total": "16663",
        "ok": "16663",
        "ko": "-"
    },
    "meanResponseTime": {
        "total": "7635",
        "ok": "7635",
        "ko": "-"
    },
    "standardDeviation": {
        "total": "4844",
        "ok": "4844",
        "ko": "-"
    },
    "percentiles1": {
        "total": "9105",
        "ok": "9105",
        "ko": "-"
    },
    "percentiles2": {
        "total": "10837",
        "ok": "10837",
        "ko": "-"
    },
    "percentiles3": {
        "total": "14965",
        "ok": "14965",
        "ko": "-"
    },
    "percentiles4": {
        "total": "16182",
        "ok": "16182",
        "ko": "-"
    },
    "group1": {
    "name": "t < 800 ms",
    "count": 455,
    "percentage": 15
},
    "group2": {
    "name": "800 ms < t < 1200 ms",
    "count": 60,
    "percentage": 2
},
    "group3": {
    "name": "t > 1200 ms",
    "count": 2485,
    "percentage": 83
},
    "group4": {
    "name": "failed",
    "count": 0,
    "percentage": 0
},
    "meanNumberOfRequestsPerSecond": {
        "total": "96.774",
        "ok": "96.774",
        "ko": "-"
    }
}
    }
}

}

function fillStats(stat){
    $("#numberOfRequests").append(stat.numberOfRequests.total);
    $("#numberOfRequestsOK").append(stat.numberOfRequests.ok);
    $("#numberOfRequestsKO").append(stat.numberOfRequests.ko);

    $("#minResponseTime").append(stat.minResponseTime.total);
    $("#minResponseTimeOK").append(stat.minResponseTime.ok);
    $("#minResponseTimeKO").append(stat.minResponseTime.ko);

    $("#maxResponseTime").append(stat.maxResponseTime.total);
    $("#maxResponseTimeOK").append(stat.maxResponseTime.ok);
    $("#maxResponseTimeKO").append(stat.maxResponseTime.ko);

    $("#meanResponseTime").append(stat.meanResponseTime.total);
    $("#meanResponseTimeOK").append(stat.meanResponseTime.ok);
    $("#meanResponseTimeKO").append(stat.meanResponseTime.ko);

    $("#standardDeviation").append(stat.standardDeviation.total);
    $("#standardDeviationOK").append(stat.standardDeviation.ok);
    $("#standardDeviationKO").append(stat.standardDeviation.ko);

    $("#percentiles1").append(stat.percentiles1.total);
    $("#percentiles1OK").append(stat.percentiles1.ok);
    $("#percentiles1KO").append(stat.percentiles1.ko);

    $("#percentiles2").append(stat.percentiles2.total);
    $("#percentiles2OK").append(stat.percentiles2.ok);
    $("#percentiles2KO").append(stat.percentiles2.ko);

    $("#percentiles3").append(stat.percentiles3.total);
    $("#percentiles3OK").append(stat.percentiles3.ok);
    $("#percentiles3KO").append(stat.percentiles3.ko);

    $("#percentiles4").append(stat.percentiles4.total);
    $("#percentiles4OK").append(stat.percentiles4.ok);
    $("#percentiles4KO").append(stat.percentiles4.ko);

    $("#meanNumberOfRequestsPerSecond").append(stat.meanNumberOfRequestsPerSecond.total);
    $("#meanNumberOfRequestsPerSecondOK").append(stat.meanNumberOfRequestsPerSecond.ok);
    $("#meanNumberOfRequestsPerSecondKO").append(stat.meanNumberOfRequestsPerSecond.ko);
}
