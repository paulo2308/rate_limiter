# 🛡️ Rate Limiter

Este projeto é uma implementação de um **Rate Limiter** desenvolvido em Go, com persistência no Redis. Ele permite controlar o número de requisições por **IP** ou por **token de acesso** (usando o header `API_KEY`).

---

## 🚀 Como Funciona

O rate limiter controla o número de requisições com base nos seguintes critérios:

- **Por IP**: limita requisições de um mesmo IP por segundo.
- **Por Token**: se o header `API_KEY` estiver presente, o controle será feito pelo token.
  - Tokens têm prioridade sobre IPs. Se um token tiver um limite maior, ele será usado.

Se o limite for ultrapassado, o IP ou token será **bloqueado por um tempo configurável**. Todas as requisições nesse período serão recusadas com:

- **HTTP 429 - Too Many Requests**
- Mensagem: `"you have reached the maximum number of requests or actions allowed within a certain time frame"`

---

## ⚙️ Configuração

As configurações são feitas por meio de variáveis de ambiente carregadas dinamicamente do seguinte arquivo:

- `.env` → usado no ambiente local

### ✅ Variáveis disponíveis

| Variável              | Descrição                                     | Exemplo         |
|-----------------------|-----------------------------------------------|-----------------|
| `REDIS_ADDR`          | Endereço do Redis                             | `localhost:6379` ou `redis:6379` |
| `REDIS_PASSWORD`      | Senha do Redis (se houver)                    | `""`            |
| `REDIS_DB`            | Banco Redis a ser utilizado                   | `0`             |
| `RATE_LIMIT_IP`       | Limite de requisições por segundo por IP      | `5`             |
| `RATE_LIMIT_TOKEN`    | Limite por segundo para tokens (API_KEY)      | `10`            |
| `BLOCK_TIME_SECONDS`  | Tempo de bloqueio após limite ser excedido    | `60`            |

---

### Comandos

1. **Clone o repositório**:

```bash
git clone https://github.com/paulo2308/rate_limiter
```

2. **Rodar a aplicação com docker-compose**
```bash
docker-compose up --build
```

3. **Rodar a aplicação com localmente**  
- **Subir o Redis:**
```bash
docker-compose up -d  
```  
- **Rodar a aplicação**
```bash
go run ./cmd 
```
