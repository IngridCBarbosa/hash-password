# Hash Password
>Hashing e checando senhas utilizando Bcypt.

>O Bcrypt foi desenvolvido para o armazenamento seguro de senhas e é "lento", o que o faz resistente a ataques de força bruta.

## Processo de Hashing
* Pegar a senha string e converte-la em um slice de bytes.
* Gerar um salto aleatório de 16 bytes.
* concatenar o salto com o slice de senha.
* Hash a senha combinada com o salto.
* Armazenar o hash e o salto no banco de dados.

### 1- Por que fazer o hash de senhas ?
> É uma prática segura para garantir que o seu a senha dos seus usuários não seja exposta em texto se o banco de dados for comprometido.
### 2- Por que utilizar o salto (salt) ?
>Salt é um valor aleatório __único__ concatenado a uma senha. Se uma senha for armazenada sem o salto, senhas iguais poderão produzir o mesmo hash com a ameaça de serem quebrados por um dicionário de hashes pré-computados de senhas hashes. Se a senha do dicionário bater com o hash no banco de dados, eles saberão a senha original. Então o salt anula o risco, pois, a senha hash não pode ser pré-calculada se o invasor não souber qual o salto foi utilizado. Se cada senha tiver um salt diferente, cada hash será será único. Além disso, quanto mais longo o salt, mais tempo leva para o hacker gerar todos os salt's possíveis e hash com todas as senhas possíveis.
### 3- Por que só armazenar o hash e o salt
>A senha original não é mais necessária. Para verificar se a senha inserida está correta, basta seguir o processo de hash original, pegando a senha inserida e fazendo hash com o salta armazenado. Se o hash gerado corresponder ao hash da senha original, eles são iguais.

