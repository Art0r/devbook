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
        data: data,
        success: function(){
            Swal.fire('Sucesso', 'Usuário cadastrado com sucesso', 'success')
            .then(() => {
                $.ajax({
                    url: '/login',
                    method: 'POST',
                    data: {
                        email,
                        password
                    },
                    success: function(){
                        window.location = "/home";
                    },
                    error: function(){
                        Swal.fire('Erro', 'Erro ao autenticar um usuário', 'error');
                    }
                })
            })
        },
        error: function(){
            Swal.fire('Sucesso', 'Usuário cadastrado com sucesso', 'success')
            .then(() => {
                $.ajax({
                    url: '/login',
                    method: 'POST',
                    data: {
                        email,
                        password
                    },
                    success: function(){
                        window.location = "/home";
                    },
                    error: function(){
                        Swal.fire('Erro', 'Erro ao autenticar um usuário', 'error');
                    }
                })
            })
        }
    });
});

