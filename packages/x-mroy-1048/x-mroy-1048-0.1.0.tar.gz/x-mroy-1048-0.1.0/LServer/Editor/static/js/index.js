
function JsonPost(url, data, callback){
    $.ajax({
        type: "POST",
        url: url,
        data: JSON.stringify(data),
        success: function( data ) {
            console.log(data);
            if (callback != null){
                callback(data);
            }
        },
        dataType: 'json'
    });
}

function web_client(url){
                var ws = new WebSocket(url); //"ws://localhost:8080/socket"
                
                ws.onopen = function() {
                   

                };

                this.send = function(msg){
                    ws.send(msg);
                }

                this.on_msg = function(callback){
                    ws.onmessage = function (evt) {
                    console.log(evt.data);
                        if (callback){
                            callback(JSON.parse(evt.data)); 
                        }
                   
                    };      
                }
                
            }

