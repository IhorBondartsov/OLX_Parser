

function GetAuthorization(){
var Request = `{ 
	"jsonrpc": "2.0",
	"id": 1,
    "method": "API.Echo", 
    "params": {"Name":"string"}
}`
var json = JSON.stringify({
    jsonrpc: "2.0",
	id: 1,
    method: "API.Echo", 
    params: {Name:"string"}
  });

    // 1. Создаём новый объект XMLHttpRequest
    //  var xhr = new XMLHttpRequest();

    var XHR = ("onload" in new XMLHttpRequest()) ? XMLHttpRequest : XDomainRequest;

    var xhr = new XHR();

// 2. Конфигурируем его: GET-запрос на URL 'phones.json'
    xhr.open('POST', "http://127.0.0.1:8001/rpc",true);  //method, URL, async, user, password

    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.setRequestHeader('Accept', 'application/json');

    xhr.onreadystatechange = function() {//Вызывает функцию при смене состояния.
        if(xhr.readyState == XMLHttpRequest.DONE && xhr.status == 200) {
            // Запрос завершен. Здесь можно обрабатывать результат.
            alert( xhr.responseText );
        }
    }

// 3. Отсылаем запрос
    xhr.send(json); //xhr.send([body])

// 4. Если код ответа сервера не 200, то это ошибка
    if (xhr.status != 200) {
        // обработать ошибку
        alert( xhr.status + ': ' + xhr.statusText ); // пример вывода: 404: Not Found
    } else {
        // вывести результат
        alert( xhr.responseText ); // responseText -- текст ответа.
    }
}
