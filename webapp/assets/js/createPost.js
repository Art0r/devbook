$('#new-post').on('click', createPost);

$(document).on('click', '.post-liked', likePost);
$(document).on('click', '.disliked-post', dislikePost);

$('#update-post').on('click', updatePost);
$('.delete-post').on('click', deletePost);

function createPost(){

    const title = document.getElementById('title').value;
    const content = document.getElementById('content').value;
    const data = {
        title,
        content
    }

    $.ajax({
        url: "/posts",
        method: "POST",
        data: data,
        success: function(data){
            document.location.reload();
        },
        error: function(req, status, err){
            document.location.reload();
        }
    });
}

function likePost(event){
    event.preventDefault();

    const clicked = $(event.target);
    const pid = clicked.closest('div').data('post-id');

    clicked.prop('disabled', true);
    $.ajax({
        url: `/posts/${pid}/like`,
        method: 'POST'
    }).done(function() {
        const likesCounter = clicked.next('span');
        const qntLikes = parseInt(likesCounter.text());

        likesCounter.text(qntLikes + 1);

        clicked.addClass('.dislike-post');
        clicked.addClass('text-danger');
        clicked.removeClass('post-liked');
    }).fail(function() {
        alert('Erro ao curtir');
    }).always(function() {
        clicked.prop('disabled', true);
    })
}

function dislikePost(event){
    event.preventDefault();

    const clicked = $(event.target);
    const pid = clicked.closest('div').data('post-id');

    clicked.prop('disabled', true);
    $.ajax({
        url: `/posts/${pid}/dislike`,
        method: 'POST'
    }).done(function() {
        const likesCounter = clicked.next('span');
        const qntLikes = parseInt(likesCounter.text());

        likesCounter.text(qntLikes - 1);

        clicked.removeClass('.dislike-post');
        clicked.removeClass('text-danger');
        clicked.addClass('post-liked');

    }).fail(function() {
        alert('Erro ao curtir');
    }).always(function() {
        clicked.prop('disabled', true);
    });
}

function updatePost(){
    $(this).prop('disabled', true);

    const pid = $(this).data('post-id');
    
    const title = document.getElementById('title').value;
    const content = document.getElementById('content').value;

    const data = {
        title,
        content
    }

    const method = "PUT";
    const url = `http://localhost:3000/posts/${pid}`;

    $.ajax({
        type: method,
        url: url,
        data: data,
        success: function(data){
            Swal.fire(
                'Sucesso',
                'Publicação criada com sucesso',
                'success'
            ).then(() => {
                window.location = "/home";
            })
        },
        error: function(req, status, err){
            alert("Erro ao atualizar a publicação");
            //alert("Email ou senha inválidos");
        }
    }).always(function() {
        $('#update-post').prop('disabled', false);
    });
}

function deletePost(event){
    event.preventDefault();

    Swal.fire({
        title: "Atenção",
        text: "Tem certeza que deseja apagar?",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then((confirm) => {
        if (!confirm.value) return;

        const clicked = $(event.target);
        const post = clicked.closest('div')
        const pid = post.data('post-id');
        
        clicked.prop('disabled', true);

        $.ajax({
            type: "DELETE",
            url: `/posts/${pid}`,
            success: function (response) {
                

                post.fadeOut('slow', function(){
                    $(this).remove();
                });
            },
            error: function(req, status, err){
                alert("Erro ao excluir a publicação");
            }
        });
    })
}