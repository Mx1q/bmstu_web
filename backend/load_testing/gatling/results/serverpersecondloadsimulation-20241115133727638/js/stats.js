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
        "total": "13349",
        "ok": "13349",
        "ko": "-"
    },
    "meanResponseTime": {
        "total": "6417",
        "ok": "6417",
        "ko": "-"
    },
    "standardDeviation": {
        "total": "4377",
        "ok": "4377",
        "ko": "-"
    },
    "percentiles1": {
        "total": "8014",
        "ok": "8014",
        "ko": "-"
    },
    "percentiles2": {
        "total": "10099",
        "ok": "10099",
        "ko": "-"
    },
    "percentiles3": {
        "total": "11818",
        "ok": "11818",
        "ko": "-"
    },
    "percentiles4": {
        "total": "12726",
        "ok": "12726",
        "ko": "-"
    },
    "group1": {
    "name": "t < 800 ms",
    "count": 1311,
    "percentage": 22
},
    "group2": {
    "name": "800 ms < t < 1200 ms",
    "count": 85,
    "percentage": 1
},
    "group3": {
    "name": "t > 1200 ms",
    "count": 4604,
    "percentage": 77
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
"req_backend-metrics-aa2a0": {
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
        "total": "13349",
        "ok": "13349",
        "ko": "-"
    },
    "meanResponseTime": {
        "total": "6390",
        "ok": "6390",
        "ko": "-"
    },
    "standardDeviation": {
        "total": "4353",
        "ok": "4353",
        "ko": "-"
    },
    "percentiles1": {
        "total": "8011",
        "ok": "8011",
        "ko": "-"
    },
    "percentiles2": {
        "total": "10087",
        "ok": "10087",
        "ko": "-"
    },
    "percentiles3": {
        "total": "11818",
        "ok": "11818",
        "ko": "-"
    },
    "percentiles4": {
        "total": "12719",
        "ok": "12719",
        "ko": "-"
    },
    "group1": {
    "name": "t < 800 ms",
    "count": 655,
    "percentage": 22
},
    "group2": {
    "name": "800 ms < t < 1200 ms",
    "count": 42,
    "percentage": 1
},
    "group3": {
    "name": "t > 1200 ms",
    "count": 2303,
    "percentage": 77
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
    },"req_nginx-metrics-ff665": {
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
        "total": "13331",
        "ok": "13331",
        "ko": "-"
    },
    "meanResponseTime": {
        "total": "6444",
        "ok": "6444",
        "ko": "-"
    },
    "standardDeviation": {
        "total": "4400",
        "ok": "4400",
        "ko": "-"
    },
    "percentiles1": {
        "total": "8014",
        "ok": "8014",
        "ko": "-"
    },
    "percentiles2": {
        "total": "10132",
        "ok": "10132",
        "ko": "-"
    },
    "percentiles3": {
        "total": "11818",
        "ok": "11818",
        "ko": "-"
    },
    "percentiles4": {
        "total": "12813",
        "ok": "12813",
        "ko": "-"
    },
    "group1": {
    "name": "t < 800 ms",
    "count": 656,
    "percentage": 22
},
    "group2": {
    "name": "800 ms < t < 1200 ms",
    "count": 43,
    "percentage": 1
},
    "group3": {
    "name": "t > 1200 ms",
    "count": 2301,
    "percentage": 77
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
