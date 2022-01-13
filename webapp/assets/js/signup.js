const submitBt = document.getElementById("submitBt");
submitBt.addEventListener("click", function(event){
    data(event);
});

function data(event){
    event.preventDefault();
    const name = document.getElementById("name").value;
    const nick = document.getElementById("nick").value;
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;
    const repassword =  document.getElementById("repassword").value;

    if (password != repassword) {
        alert("As senhas devem ser iguais");
        return
    }

    $.ajax({
        type: "POST",
        url: "/user",
        data: {
            name: name,
            email: email,
            nick: nick,
            password: password
        },
    });
}