var mail, passwd, form;


function Authorization(form){
    var elems = form.elements;
    mail = elems.email.value;
    passwd = elems.password.value;
    
    if (!isValidMail(mail)){
        return
    }
    GetAuthorization(mail, passwd)
}
