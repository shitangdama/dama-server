
// var myChart = echarts.init(document.getElementById('main'));

var table_type = "5min"
let elelist = document.querySelectorAll('[data-id="self"]');
console.log(elelist)
// console.log(elelist[0].getAttribute("info"))

elelist.forEach(function(dom){
    var myChart = echarts.init(dom);
    var data = $.parseJSON(dom.getAttribute("info_5min"))
    console.log(dom.getAttribute("info_5min"))
    var option = {
        xAxis: {
            show: false,
            type: 'category'
        },
        yAxis: {
            show: false,
            type: 'value'
        },
        series: [{
            data: data["close"],
            markPoint: {
                data: [
                    {type: 'max', name: '最大值'},
                    {type: 'min', name: '最小值'}
                ]
            },
            // markLine: {
            //     data: [
            //         {type: 'average', name: '平均值'}
            //     ]
            // },
            type: 'line',
            showSymbol: false,
        }]
    };

    myChart.setOption(option);

})

let contrastlist = document.querySelectorAll('[data-id="contrast"]');

contrastlist.forEach(function(dom){
    var data = $.parseJSON(dom.getAttribute("info_5min"))
    if(data["contrast"]) {
        var myChart = echarts.init(dom);
    
        console.log(dom.getAttribute("info_5min"))
        var option = {
            xAxis: {
                show: false,
                type: 'category'
            },
            yAxis: {
                show: false,
                type: 'value'
            },
            series: [{
                data: data["close"],
                markPoint: {
                    data: [
                        {type: 'max', name: '最大值'},
                        {type: 'min', name: '最小值'}
                    ]
                },
                // markLine: {
                //     data: [
                //         {type: 'average', name: '平均值'}
                //     ]
                // },
                type: 'line',
                showSymbol: false,
            }]
        };
        myChart.setOption(option);
    }
})

let vollist = document.querySelectorAll('[data-id="vol"]');

vollist.forEach(function(dom){
    var data = $.parseJSON(dom.getAttribute("info_5min"))

    var myChart = echarts.init(dom);

    console.log(dom.getAttribute("info_5min"))
    var option = {
        xAxis: {
            show: false,
            type: 'category'
        },
        yAxis: {
            show: false,
            type: 'value'
        },
        series: [{
            data: data["vol"],
            markPoint: {
                data: [
                    {type: 'max', name: '最大值'},
                    {type: 'min', name: '最小值'}
                ]
            },
            // markLine: {
            //     data: [
            //         {type: 'average', name: '平均值'}
            //     ]
            // },
            type: 'bar',
            showSymbol: false,
        }]
    };
    myChart.setOption(option);
})