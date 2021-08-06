var Words = document.getElementById("talkShowIdId");
var Input = document.getElementById("talkInputId");
var Send = document.getElementById("sendButtonId");
var socket = new WebSocket("ws://localhost:8080/echo");

// function InputPress() {
//     if (Event.keyCode === 13) {
//         chatRoom();
//     }
// }

Send.onclick = function () {
    var str = "";
    if (Input.value === "") {
        // 消息为空时弹窗
        alert("消息不能为空");
        return;
    }
    str = '<div class="myTalk"><span>' + Input.value + '</span></div>';

    socket.send(Input.value);
    Input.value = "";
    // 将之前的内容与要发的内容拼接好 提交
    Words.innerHTML = Words.innerHTML + str;
}


socket.onmessage = function (e) {
    str = '<div class="otherTalk"><span>' + e.data + '</span></div>';
    Words.innerHTML = Words.innerHTML + str;
}
