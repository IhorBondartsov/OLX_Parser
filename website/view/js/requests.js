function GetAuthorization(login, password){

var json = JSON.stringify({
    Method:"Login",
    Data:{
        Login:login,
        Password:password
    }
  });

    // 1. Создаём новый объект XMLHttpRequest
    //  var xhr = new XMLHttpRequest();

    var XHR = ("onload" in new XMLHttpRequest()) ? XMLHttpRequest : XDomainRequest;

    var xhr = new XHR();

// 2. Конфигурируем его: GET-запрос на URL 'phones.json'
    xhr.open('POST', "http://127.0.0.1:8000/userms/rpc",true);  //method, URL, async, user, password

    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.setRequestHeader('Accept', 'application/json');

    xhr.onreadystatechange = function() {//Вызывает функцию при смене состояния.
        if(xhr.readyState == XMLHttpRequest.DONE && xhr.status == 200) {
            // Запрос завершен. Здесь можно обрабатывать результат.
            var tokens = JSON.parse(xhr.responseText);
            if (tokens.RefreshToken !== undefined) {
                sessionStorage.setItem('Mail', login);
                sessionStorage.setItem('RefreshToken', tokens.RefreshToken);
                sessionStorage.setItem('AccessToken', tokens.AccessToken);
                document.location.href='http://127.0.0.1:8000/index2.html'
                 }else{
                alert(xhr.responseText)
            }
        }
    };

// 3. Отсылаем запрос
    xhr.send(json);

    var tokens = JSON.parse(xhr.responseText);

    // if (tokens.RefreshToken !== undefined) {
    //     alert(tokens.RefreshToken );
    // } else {
    //     alert("nothing");
    // }
}

function MakeOrder(form){
    var elems = form.elements;
    var url = elems.url.value;

    var login = sessionStorage.getItem('Mail');
    var accessToken = sessionStorage.getItem('AccessToken');
    var json = JSON.stringify({
        Method:"MakeOrder",
        Data:{
            Token: accessToken,
            Mail:login,
            DateTo: (Date.now() / 1000) + 1000,
            Frequency:      60,
            PageLimit:      1,
            URL:            url,
            UserID:         1
        }
    });
    console.log(json);

    // 1. Создаём новый объект XMLHttpRequest
    //  var xhr = new XMLHttpRequest();

    var XHR = ("onload" in new XMLHttpRequest()) ? XMLHttpRequest : XDomainRequest;

    var xhr = new XHR();

// 2. Конфигурируем его: GET-запрос на URL 'phones.json'
    xhr.open('POST', "http://127.0.0.1:8000/parser/rpc",true);  //method, URL, async, user, password

    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.setRequestHeader('Accept', 'application/json');

    xhr.onreadystatechange = function() {//Вызывает функцию при смене состояния.
        if(xhr.readyState == XMLHttpRequest.DONE && xhr.status == 200) {
            // Запрос завершен. Здесь можно обрабатывать результат.
            var tokens = JSON.parse(xhr.responseText);
            if ({} === tokens) {
                alert("Check your Post")
            }else{
                alert(xhr.responseText)
            }
        }
    };

// 3. Отсылаем запрос
    xhr.send(json);

    var tokens = JSON.parse(xhr.responseText);

    // if (tokens.RefreshToken !== undefined) {
    //     alert(tokens.RefreshToken );
    // } else {
    //     alert("nothing");
    // }
}
















//
// function GetAuthorization(){
//     var Request = `{
// 	"jsonrpc": "2.0",
// 	"id": 1,
//     "method": "API.Echo",
//     "params": {"Name":"string"}
// }`
//     var json = JSON.stringify({
//         jsonrpc: "2.0",
//         id: 1,
//         method: "API.Echo",
//         params: {Name:"string"}
//     });
//
//     // 1. Создаём новый объект XMLHttpRequest
//     //  var xhr = new XMLHttpRequest();
//
//     var XHR = ("onload" in new XMLHttpRequest()) ? XMLHttpRequest : XDomainRequest;
//
//     var xhr = new XHR();
//
// // 2. Конфигурируем его: GET-запрос на URL 'phones.json'
//     xhr.open('POST', "http://127.0.0.1:8001/rpc",true);  //method, URL, async, user, password
//
//     xhr.setRequestHeader('Content-Type', 'application/json');
//     xhr.setRequestHeader('Accept', 'application/json');
//
//     xhr.onreadystatechange = function() {//Вызывает функцию при смене состояния.
//         if(xhr.readyState == XMLHttpRequest.DONE && xhr.status == 200) {
//             // Запрос завершен. Здесь можно обрабатывать результат.
//             alert( xhr.responseText );
//         }
//     }
//
// // 3. Отсылаем запрос
//     xhr.send(json); //xhr.send([body])
//
// // 4. Если код ответа сервера не 200, то это ошибка
//     if (xhr.status != 200) {
//         // обработать ошибку
//         alert( xhr.status + ': ' + xhr.statusText ); // пример вывода: 404: Not Found
//     } else {
//         // вывести результат
//         alert( xhr.responseText ); // responseText -- текст ответа.
//     }
// }
