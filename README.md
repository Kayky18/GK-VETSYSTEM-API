# VetSystem API

## Descrição

A **VetSystem API** é uma API REST que permite gerenciar pacientes (animais) em um sistema veterinário. Ela suporta operações como listar, criar, atualizar, visualizar e excluir pacientes, além de fornecer respostas padronizadas para erros e resultados. A API está documentada com Swagger e foi construída com Golang.

## Endpoints

### 1. Listar Pacientes
- **Endpoint:** `GET /patients`
- **Descrição:** Lista todos os pacientes cadastrados no sistema.
- **Respostas:**
  - `200 OK`: Retorna a lista de pacientes.
  - `500 Internal Server Error`: Erro no servidor.

### 2. Visualizar Paciente
- **Endpoint:** `GET /patient/{id}`
- **Descrição:** Exibe os detalhes de um paciente específico.
- **Parâmetros:**
  - `id` (path): Identificador do paciente.
- **Respostas:**
  - `200 OK`: Retorna os detalhes do paciente.
  - `400 Bad Request`: Solicitação inválida.
  - `404 Not Found`: Paciente não encontrado.

### 3. Criar Paciente
- **Endpoint:** `POST /patient/create`
- **Descrição:** Cria um novo paciente no sistema.
- **Corpo da Requisição:**
  ```json
  {
    "age": 5,
    "breed": "Golden Retriever",
    "cpf": "12345678900",
    "name": "Rex",
    "nameTutor": "João Silva",
    "phone": "11999999999",
    "species": "Cachorro",
    "temperature": 38.5,
    "weight": 30
  }
  ```
- **Respostas:**
  - `200 OK`: Paciente criado com sucesso.
  - `400 Bad Request`: Erro nos dados fornecidos.
  - `500 Internal Server Error`: Erro no servidor.

### 4. Atualizar Paciente
- **Endpoint:** `PUT /patient/{id}`
- **Descrição:** Atualiza as informações de um paciente existente.
- **Parâmetros:**
  - `id` (path): Identificador do paciente.
- **Corpo da Requisição:**
  ```json
  {
    "age": 6,
    "breed": "Golden Retriever",
    "name": "Rex",
    "nameTutor": "João Silva",
    "phone": "11999999999",
    "species": "Cachorro",
    "state": "São Paulo",
    "temperature": 38.7,
    "weight": 32
  }
  ```
- **Respostas:**
  - `200 OK`: Paciente atualizado com sucesso.
  - `400 Bad Request`: Erro nos dados fornecidos.
  - `404 Not Found`: Paciente não encontrado.

### 5. Excluir Paciente
- **Endpoint:** `DELETE /patient/{id}`
- **Descrição:** Exclui um paciente específico.
- **Parâmetros:**
  - `id` (path): Identificador do paciente.
- **Respostas:**
  - `200 OK`: Paciente excluído com sucesso.
  - `404 Not Found`: Paciente não encontrado.

### 6. Bad Request Pacientes
- **Endpoint:** `GET /patient`
- **Descrição:** Simulação de requisição inválida (Bad Request).
- **Respostas:**
  - `200 OK`: Simula um Bad Request para pacientes.

## Estrutura das Respostas

As respostas da API seguem um formato padrão com os campos:

```json
{
  "data": { /* Dados do Paciente */ },
  "message": "Mensagem sobre a operação realizada"
}
```

## Execução com Docker Compose

Para executar a API utilizando Docker Compose, siga os seguintes passos:

1. **Criar a imagem Docker:**
   Execute o comando para buildar a imagem:
   ```bash
   docker build -t gk-vetsystem .
   ```

2. **Executar a API:**
   Utilizando o `docker-compose.yml` fornecido:
   ```bash
   version: '3'
   services:
     api_gk:
       image: gk-vetsystem
       ports:
         - "8080:8080"
       command: go run main.go
   ```
   Para iniciar a API:
   ```bash
   docker-compose up
   ```

3. **Acessar a API:**
   Acesse a API via `http://localhost:8080`.

## Tecnologias Utilizadas

- **Golang**: Linguagem de programação principal.
- **Docker**: Para containerização da aplicação.
- **Swagger**: Para documentação da API.

## Contribuição

1. Faça um fork do repositório.
2. Crie uma nova branch: `git checkout -b feature/nova-funcionalidade`.
3. Faça commit das suas alterações: `git commit -m 'Adiciona nova funcionalidade'`.
4. Envie suas alterações: `git push origin feature/nova-funcionalidade`.
5. Abra um Pull Request.

## Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

--- 

Esse README fornece as instruções básicas para executar o projeto e documenta os endpoints da API.
