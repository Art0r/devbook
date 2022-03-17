$('#follow').on('click', follow);
$('#unfollow').on('click', unfollow);
$('#edit-user').on('submit', edit);
$('#update-password').on('submit', updatePassword);
$('#delete-user').on('click', deleteUser)

function unfollow(event){
    event.preventDefault();

    const uid = $(this).data('user-id');
    $(this).prop('disabled', true);

    $.ajax({
        type: "POST",
        url: `/users/${uid}/unfollow`,
        success: function(response) {
            window.location = `/users/${uid}`
        },
        fail: function (response){
            Swal.fire("Erro ao parar de seguir um usuário", "error");
            $('#unfollow').prop('disabled', false);
        }
    });
}

function follow(event){
    event.preventDefault();

    const uid = $(this).data('user-id');
    $(this).prop('disabled', true);

    $.ajax({
        type: "POST",
        url: `/users/${uid}/follow`,
        success: function(response) {
            window.location = `/users/${uid}`
        },
        fail: function (response){
            Swal.fire("Erro ao seguir um usuário", "error");
            $('#follow').prop('disabled', false);
        }
    });
}

function edit(event){
    event.preventDefault();

    const name = document.getElementById("name").value;
    const email = document.getElementById("email").value;
    const nick = document.getElementById("nick").value;

    $.ajax({
        type: "PUT",
        url: "/edit-user",
        data: {
            name,
            email,
            nick
        },
        success: function (response) {
            Swal.fire("Sucesso!", "Usuário cadastrado com sucesso!", "success").then(() => {
                window.location = "/profile";
            });
        },
        fail: function () {
            Swal.fire("Ops...!", "Erro ao cadastrar o usuário!", "error");
        }
    });
}

function updatePassword(event) {
    event.preventDefault();

    const oldPassword = document.getElementById('this-password').value;
    const newPassword = document.getElementById('new-password').value;
    const confirmPassword = document.getElementById('confirm-password').value;

    if (confirmPassword != newPassword){
        Swal.fire('Ops....', 'As senhas não coincidem', 'warning');
        return;
    }

    $.ajax({
        url: "/update-password",
        method: 'POST',
        data: {
            newPassword,
            oldPassword,
        },
        success: function(){
            Swal.fire('Sucesso!', "A senha foi atualizada com sucesso", "success").then(() => {
                window.location = "/profile";
            });
        },
        fail: function(){
        }
    })
}

function deleteUser(){
    Swal.fire({
        title: "Atenção",
        text: "Tem certeza que deseja apagar sua conta? Essa ação é irreversível!",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then((confirm) => {
        if (confirm.value) {
            $.ajax({
                type: "DELETE",
                url: "/delete-user",
                success: function () {
                    Swal.fire("Sucesso!", "Seu usuário foi excluído com sucesso!", "success").then(() => {
                        window.location = "/logout";
                    })
                },
                fail: function () {
                    Swal.fire('Ops...', "Ocorreu um erro ao excluir o usuário", "error")
                }
            });
        }
    })
}