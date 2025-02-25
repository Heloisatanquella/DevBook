$('#parar-de-seguir').on('click', pararDeSeguir);
$('#seguir').on('click', seguir);
$('#editar-usuario').on('submit', editar);

function pararDeSeguir() {
    const usuarioId = $(this).data('usuario-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/usuarios/${usuarioId}/parar-de-seguir`,
        method: "POST"
    }).done(function() {
        window.location = `/usuarios/${usuarioId}`;
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao parar de seguir o usuário!", "error");
        $('parar-de-seguir').prop('disabled', false);
    });
}

function seguir() {
    const usuarioId = $(this).data('usuario-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/usuarios/${usuarioId}/seguir`,
        method: "POST"
    }).done(function() {
        window.location = `/usuarios/${usuarioId}`;
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao seguir o usuário!", "error");
        $('seguir').prop('disabled', false);
    });
}

function editar(evento) {
    evento.preventDefault();

    $.ajax({
        url: "/editar-usuario",
        method: "PUT",
        contentType: "application/json",
        data: JSON.stringify({
            nome: $('#nome').val(), 
            email: $('#email').val(),
            nick: $('#nick').val(),
        })
    }).done(function() {
        Swal.fire("Sucesso!", "Dados do usuário foram atualizados com sucesso", "success")
            .then(function() {
                window.location = "/perfil";
            })
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao atualizar os dados do usuário!", "error");
    });
}
