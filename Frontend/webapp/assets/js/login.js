$('#login').on('submit', fazerLogin)

function fazerLogin(evento) {
    evento.preventDefault();

    $.ajax({
        url: "/login",
        method: "POST",
        contentType: "application/json",
        data: JSON.stringify({
            email: $('#email').val(),
            senha: $('#senha').val(),
        })
    }).done(function() {
        window.location = "/home"
    }).fail(function() {
        Swal.fire("Ops...", "Usuário ou senhas incorretos!", "error");
    });
}