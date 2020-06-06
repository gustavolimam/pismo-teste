# Documentação

## Execução do projeto

Para o correto funcionamento da aplicação, na raiz do projeto devemos executar o seguinte comando:

```sh
docker-compose up
```

Esse comando será responsável por criar a imagem com o código compilado e ao mesmo tempo subir o banco de dados.

## Detalhes das APIs

Para a correta execução das APIs, deverá ser utilizado como base da url da requisição `localhost:3003`.

### createAccount

API responsável pela criação de uma nova _account_, gerando um _id_ dinamicamente.

**URL:** `base_url/accounts`

**Method:** `POST`

**Body:**

```json
{
  "document_number": "1111"
}
```

**Response OK**

```json
{
  "ID": 1,
  "Message": "Sucesso"
}
```

### getAccount

API responsável por retornar os dados de uma _account_ específica.

**URL:** `base_url/accounts/:account_id`

**Method:** `GET`

**Response OK**

```json
{
  "account_id": 1,
  "document_number": "111"
}
```

### createTransaction

API responsável pela criação de uma nova _transaction_, gerando um _id_ dinamicamente.

**URL:** `base_url/transactions`

**Method:** `POST`

**Body:**

```json
{
  "account_id": 1,
  "operation_type_id": 1,
  "amount": 199.9
}
```

**Response OK**

```json
{
  "ID": 1,
  "Message": "Sucesso"
}
```
