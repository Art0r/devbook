document.getElementById('signup-form').addEventListener("submit", function(event){
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
    
    const url = "http://localhost:3000/user";
    const data = {
        name,        
        email,        
        nick,        
        password     
    }
    
    const request = {
        method: "POST",
        body: JSON.stringify(data),
        headers: {
            'Content-type': 'application/json'
        }
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

    /*
    fetch(url, request)
    .then(response => response.json())
    .then((json) => {
        if (json.error) {
            alert(json.error);
        }
    })
    .catch(err => {
        console.log(err);
        alert("Erro ao cadastrar o usuário");
    });*/
});

