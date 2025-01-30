// $(document).ready(function () {
// });
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
        data: {
           nome: $('#nome').val(), 
           email: $('#email').val(),
           nick: $('#nick').val(),
           senha: $('#senha').val()
        }
    })
}