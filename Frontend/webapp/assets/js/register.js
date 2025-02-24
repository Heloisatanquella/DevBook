
$('#register-form').on('submit', criarUsuario);

function criarUsuario(evento) {
    evento.preventDefault();

    if ($('#senha').val() != $('#confirmar-senha').val()) {
        Swal.fire("Ops...", "As senhas não coincidem!", "error");
        return;
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        contentType: "application/json",
        data: JSON.stringify({
            nome: $('#nome').val(), 
            email: $('#email').val(),
            nick: $('#nick').val(),
            senha: $('#senha').val()
        })
    }).done(function() {
        Swal.fire("Sucesso!", "O usuário foi cadastrado com sucesso!", "success")
            .then(function() {  
                $.ajax({
                    url: "/login",
                    method: "POST",
                    contentType: "application/json",
                    data: JSON.stringify({
                        email: $('#email').val(),
                        senha: $('#senha').val()
                    }).done(function() {
                        window.location = "/home";
                    }).fail(function() {
                        Swal.fire("Ops...", "Erro ao autenticar o usuário!", "error");
                    })
                })
            })
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao cadastrar o usuário!", "error");
    });
}