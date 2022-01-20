$('#signup-form').on('submit', Signup)

function Signup(event){
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

    const url = "/user";
    const data = {
        name,        
        email,        
        nick,        
        password     
    }

    $.ajax({
        type: "POST",
        url: url,
        data: data
    }).done(() => {
        alert("Usuário cadastrado com sucesso")
    }).fail((err) => {
        console.log(err)
        alert("Erro ao cadastrar o usuário")
    });

}