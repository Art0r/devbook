$('#login').on('submit', Signin);

function Signin(event) {
    event.preventDefault();

    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;

    $.ajax({
        type: "POST",
        url: "http://localhost:3000/login",
        data: {
            email,
            password
        },
    }).done(function () {
        window.location = "/home";  
    }).fail(function () {
        alert("Email ou senha inválidos");
    });
}