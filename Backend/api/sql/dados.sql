INSERT INTO usuarios(nome, nick, email, senha)
VALUES
("Usuário 1", "usuario_1", "usuario1@gmail.com", "$2a$10$h0.tQrjN1ae1lJpmm2fBF.xYwJtNPXUoC9DI5x9ibYvY/mm303n9i"), -- usuario 1
("Usuário 2", "usuario_2", "usuario2@gmail.com", "$2a$10$h0.tQrjN1ae1lJpmm2fBF.xYwJtNPXUoC9DI5x9ibYvY/mm303n9i"), -- usuario 2
("Usuário 3", "usuario_3", "usuario3@gmail.com", "$2a$10$h0.tQrjN1ae1lJpmm2fBF.xYwJtNPXUoC9DI5x9ibYvY/mm303n9i"); -- usuario 3
("Usuário 4", "usuario_4", "usuario4@gmail.com", "$2a$10$h0.tQrjN1ae1lJpmm2fBF.xYwJtNPXUoC9DI5x9ibYvY/mm303n9i"); -- usuario 3

INSERT INTO seguidores(usuario_id, seguidor_id)
VALUES
(1, 2),
(3, 1),
(1, 3),
(2, 4);

INSERT INTO publicacoes(titulo, conteudo, autor_id)
VALUES
("Publicação do usuário 1", "Essa é a publicação do usuário 1! Oba!", 1),
("Publicação do usuário 2", "Essa é a publicação do usuário 2! Oba!", 2),
("Publicação do usuário 3", "Essa é a publicação do usuário 3! Oba!", 3);