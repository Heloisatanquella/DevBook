$('#register-form').on('submit', criarUsuario);

function criarUsuario(evento) {
    evento.preventDefault();
    console.log("funcionou")

    if ($('#senha').val() != $('#confirmar-senha').val()) {
        alert("As senha não coincidem!")
        return
    }
    $.ajax({
        url: "/users",
        method: "POST",
        data: {
            name: ('#nome').val(),
            email: ('#email').val(),
            nick: ('#nick').val(),
            password: ('#senha').val(),
        }
    })
}