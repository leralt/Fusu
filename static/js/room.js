window.onload = function () {
    var Words = document.getElementById("talkShowId");
    var Input = document.getElementById("talkInputId");
    var Send = document.getElementById("sendButtonId");
    var socket = new WebSocket("ws://localhost:8080/room");


    function sendMessage(){
        var str = "";
        if (Input.value === "") {
            // 消息为空时弹窗
            alert("消息不能为空");
            return;
        }
        var name= "long";
        var d = new Date();
        var time = d.toLocaleTimeString();
        str = '<div class="MNameTime"><span>' + time + '</span><span> ' + name + '</span></div>'
        str = str + '<div class="myTalk"><span>' + Input.value + '</span></div>';

        var message = {name: name, time: time, msg: Input.value}
        console.log(message);
        var json = JSON.stringify(message)
        console.log(json)
        socket.send(json);
        Input.value = "";
        // 将之前的内容与要发的内容拼接好提交
        Words.innerHTML = Words.innerHTML + str;
        Words.scrollTop = Words.scrollHeight;
    }
    Send.onclick = function () {
        sendMessage();
    }

    Input.onkeypress = function (e) {
        if (e.code === "Enter") {
            sendMessage();
        }
    }

    socket.onmessage = function (e) {
        var str = ""
        data = JSON.parse(e.data);
        console.log(data);
        console.log(data.date)
        str = '<div class="ONameTime"><span>' + data.name + '</span><span> ' + data.time + '</span></div>'
        str = str + '<div class="otherTalk"><span>' + data.msg + '</span></div>';
        Words.innerHTML = Words.innerHTML + str;
        Words.scrollTop = Words.scrollHeight;
    }

    socket.onclose = function () {

    }
};

