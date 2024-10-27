# Sistema de Stress test

Uma ferramenta de linha de comando em Go para realizar testes de carga em serviços web. Permite executar múltiplas requisições HTTP concorrentes e gera relatórios detalhados sobre o desempenho do serviço.

## 🚀 Funcionalidades

- ✨ Execução de requisições HTTP concorrentes
- 📊 Geração de relatórios detalhados
- ⚡ Controle preciso do número de requisições simultâneas
- 🔍 Análise da distribuição de códigos de status HTTP
- 🐳 Suporte a Docker para fácil execução

## 📋 Pré-requisitos

- Go 1.21 ou superior (apenas para build local)
- Docker (para execução via container)

## 🛠️ Instalação

### Usando Docker (Recomendado)

1. Clone o repositório:
```bash
git clone https://github.com/wanderlei2583/stressTest.git
cd stressTest
```

2. Construa a imagem Docker:
```bash
docker build -t load-tester .
```

### Build Local

1. Clone o repositório:
```bash
git clone https://github.com/wanderlei2583/stressTest.git
cd stressTest
```

2. Compile o projeto:
```bash
go build -o load-tester
```

## 💻 Uso

### Via Docker

```bash
docker run load-tester --url=http://exemplo.com --requests=1000 --concurrency=10
```

### Localmente

```bash
./load-tester --url=http://exemplo.com --requests=1000 --concurrency=10
```

### Parâmetros

| Parâmetro | Descrição | Obrigatório | Padrão |
|-----------|-----------|-------------|---------|
| --url | URL do serviço a ser testado | Sim | - |
| --requests | Número total de requisições | Sim | - |
| --concurrency | Número de requisições simultâneas | Não | 1 |

## 📊 Exemplo de Saída

```
Iniciando teste de carga para http://exemplo.com
Total de requests: 1000
Concorrência: 10

Relatório do Teste de Carga
==========================
Tempo total de execução: 5.2s
Total de requests: 1000
Requests com erro: 0

Distribuição de Status HTTP:
--------------------------
HTTP 200: 985 requests (98.5%)
HTTP 404: 15 requests (1.5%)
```

## 🔧 Configurações Avançadas

- Timeout de requisições: 10 segundos (configurado no código)
- Distribuição automática de carga entre workers
- Tratamento de erros e timeouts

## 📝 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## ✨ Autor

Wanderlei Pereira - [wanderlei2583@gmail.com](mailto:wanderlei2583@gmail.com)

Project Link: [https://github.com/wanderlei2583/stressTest](https://github.com/wanderlei2583/stressTest)
