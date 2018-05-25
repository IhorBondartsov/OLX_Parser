

function GetAuthorization(){
var Request = `{
    "method": "Echo", 
    "jsonrpc": "2.0",
    "params": {"name":"string"}, 
    "id": 1
}`

    // 1. Создаём новый объект XMLHttpRequest
    var xhr = new XMLHttpRequest();

// 2. Конфигурируем его: GET-запрос на URL 'phones.json'
    xhr.open('POST', "https://127.0.0.1:8001/rpc",false );  //method, URL, async, user, password

    xhr.setRequestHeader('Content-Type', 'application/json');

// 3. Отсылаем запрос
    xhr.send(Request); //xhr.send([body])

// 4. Если код ответа сервера не 200, то это ошибка
    if (xhr.status != 200) {
        // обработать ошибку
        alert( xhr.status + ': ' + xhr.statusText ); // пример вывода: 404: Not Found
    } else {
        // вывести результат
        alert( xhr.responseText ); // responseText -- текст ответа.
    }
}
