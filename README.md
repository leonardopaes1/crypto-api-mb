# 🚀 Projeto: Crypto API - Aplicação Go + Kubernetes

## 🎯 Objetivo

Desenvolver uma API em Go para simular consulta de pares de criptomoedas, conteinerizar a aplicação, criar Helm Chart e realizar deploy automatizado via GitHub Actions.

---

## 📦 Estrutura

- **Aplicação:** Go 1.22+
- **Container:** Docker
- **Orquestração:** Helm + Kubernetes
- **CI/CD:** GitHub Actions

---

## 🔗 Repositório da Infraestrutura

Provisionamento da infraestrutura está disponível aqui:
- [leonardopaes1/infra-gke-mb](https://github.com/leonardopaes1/infra-gke-mb)

---

## 📈 Funcionalidades da API

- Consulta de preço de criptomoedas em tempo real utilizando a [Mercado Bitcoin API](https://api.mercadobitcoin.net/api/v4/docs).
- Healthcheck da aplicação (`/healthz`) para monitoramento básico.

### Exemplos de Pares de Criptomoedas (Sugestões para Testes)

Aqui estão algumas sugestões de siglas de criptomoedas que você pode consultar na API:

| **Par** | **Descrição**             |
|:--------|:---------------------------|
| BTC     | Bitcoin                    |
| ETH     | Ethereum                   |
| ADA     | Cardano                    |
| SOL     | Solana                     |
| XRP     | Ripple                     |
| DOGE    | Dogecoin                   |
| DOT     | Polkadot                   |
| LTC     | Litecoin                   |
| MATIC   | Polygon                    |
| AVAX    | Avalanche                  |

### Exemplo de Requisição

```bash
https://SEU-ENDERECO-API/BTC
```
Resposta esperada:
```bash
{
  "environment": "prod",
  "ticker": [
    {
      "high": "545000.00000000",
      "low": "535651.00000000",
      "vol": "18.60714195",
      "last": "538183.00000000",
      "buy": "538455.00000000",
      "sell": "538907.00000000",
      "open": "540127.00000000",
      "date": 1745777208,
      "pair": "BTC"
    }
  ]
}

```

---

## 🔒 Secrets Necessárias (GitHub)

| Secret             | Descrição                                                    |
|--------------------|---------------------------------------------------------------|
| `APPLICATION_NAME` | Nome da aplicação e Helm release                             |
| `DOCKER_PASSWORD`  | Senha do DockerHub para fazer push da imagem                  |
| `DOCKER_REPO`      | Nome do repositório no DockerHub (ex: `usuario/crypto-api`)    |
| `DOCKER_USERNAME`  | Usuário do DockerHub                                           |
| `GCP_CREDENTIALS`  | JSON de autenticação da Service Account                       |
| `GCP_PROJECT_ID`   | ID do projeto na Google Cloud                                 |
| `REPLICA_COUNT`    | Número de réplicas para o deploy                              |
| `SERVICE_PORT`     | Porta que o serviço expõe a aplicação                         |

---

## ⚙️ Pipelines CI/CD

- **Build & Push da Imagem Docker** para o DockerHub
- **Deploy Automatizado via Helm** em GKE:
  - Branch `staging`: deploy automático em ambiente staging
  - Tag `vX.X.X`: deploy automático em ambiente de produção

---

## 🌐 Observabilidade

- Uptime Check para o endpoint `/healthz`
- Dashboard para pods da aplicação no Google Monitoring
- Alertas para falha de disponibilidade

---

## 🚀 Caso Queira Executar Local

Para executar localmente precisa estar logado no google com o ``gcloud auth login`` e executar o comando abaixo substituindo os valores que estão entre <> pelos valores desejados.

```bash
helm upgrade --install <Nome para o helm> ./helm/crypto-api \
          --namespace <Namespace> --create-namespace \
          --set environment=<enviroument> \
          --set replicaCount=<Número de replicas> \
          --set image.repository=<Caminho da imagem no Docker hub> \
          --set image.tag=<Tag da imagem> \
          --set service.type=LoadBalancer \
          --set service.port=<Porta de exposição do serviço> \
          --set nodeSelector.pool=<prod ou staging>-pool
```

---

## 📋 Decisões Técnicas

- Go mod para gestão de dependências
- Healthcheck incluído para facilitar observabilidade
- Imagens versionadas pelo SHA ou pela Tag
- Deploy seguro e controlado por Branch ou Tag

---

## ✅ Status

| Item | Status |
|:---|:---|
| Build e Push funcionando | ✅ |
| Deploy com Helm funcional | ✅ |
| Healthcheck integrado | ✅ |
| Monitoramento configurado | ✅ |

---

## 📄 Licença

Este projeto é open-source para fins de estudo e demonstração de habilidades DevOps/SRE.

---