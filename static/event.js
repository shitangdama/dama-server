
$(document).ready(function(){
    $("button").click(function(d){
        var symbol = $(d.target).attr("symbol")
        if(symbol) {
            axios.post('/subscribe', {
                symbol: symbol,
            })
            .then(function (response) {
                console.log(response);
            })
            .catch(function (error) {
                console.log(error);
            });
        }
    });
})