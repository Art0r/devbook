$('#follow').on('click', follow);
$('#unfollow').on('click', unfollow);
$('#edit-user').on('submit', edit);

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
