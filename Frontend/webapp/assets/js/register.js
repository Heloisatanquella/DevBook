
$('#register-form').on('submit', criarUsuario);

function criarUsuario(evento) {
    evento.preventDefault();

    if ($('#senha').val() != $('#confirmar-senha').val()) {
        alert("Ops...as senhas não coincidem!");
        return;
    }
    $.ajax({
        url: "http://localhost:5000/usuarios",
        method: "POST",
        contentType: "application/json",
        data: JSON.stringify({
            nome: $('#nome').val(), 
            email: $('#email').val(),
            nick: $('#nick').val(),
            senha: $('#senha').val()
        })
    }).done(function() {
        alert("Usuário cadastrado com sucesso!");
    }).fail(function() {
        alert("Erro ao cadastrar o usuário!");
    });
}