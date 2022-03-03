$('#login').on('submit', Signin);

function Signin(event) {
    event.preventDefault();

    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;
    const data = {
        email, 
        password
    }
    const method = "POST";
    const url = "http://localhost:3000/login";

    $.ajax({
        type: method,
        url: url,
        data: data,
        success: function(data){
            window.location = "/home";
        },
        error: function(req, status, err){
            window.location = "/home";
            //alert("Email ou senha inválidos");
        }
    });
}