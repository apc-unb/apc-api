![](https://i.imgur.com/LN5iFfC.jpg)

![GitHub top language](https://img.shields.io/github/languages/top/apc-unb/apc-api?style=plastic)
![GitHub language count](https://img.shields.io/github/languages/count/apc-unb/apc-api?color=yellow)
![GitHub last commit](https://img.shields.io/github/last-commit/apc-unb/apc-api)
[![GitHub stars](https://img.shields.io/github/stars/apc-unb/apc-api)](https://github.com/apc-unb/apc-api/stargazers)
![GitHub](https://img.shields.io/github/license/apc-unb/apc-api?color=blue)


# Visão Geral

DragonT é uma plataforma que centraliza o acesso à materiais do curso. Fornece relatórios automáticos do progresso de cada aluno individualmente e da turma como um todo.

  

# Objetivo

O objetivo desse software será para facilitar o trabalho do professor no acompanhamento de cada aluno, tendo relatório automático do progresso da turma, e também permitindo que você tenha uma noção do desenvolvimento da turma, pra ver se aquele aluno merece o famoso 0,1 no final do semestre. Também irá servir o conteúdo teórico que atualmente está neste [link](https://carlacastanho.github.io/Material-de-APC/).

  

# Features

- Login individual para o aluno

  

- Dashboard com progresso da turma e métricas do curso

  

- Relatório individual de progresso do aluno

  

- Integração com Code Forces (Uri para os exercícios extras)

  

- Relatórios por semestre

  

- Possível fórum de discussão alunos - monitores para cada assunto da matéria (talvez)
 

# Tente você mesmo

## Deploy Mode

Na raiz do projeto, execute os seguintes comandos para executar o projeto no modo de produção

```bash

docker-compose up --build

```

## Developer Mode

Na raiz do projeto, execute os seguintes comandos para executar o projeto no modo desenvolvedor
  
1. Subir as aplicações que o DraGonT precisa

  

```bash

docker-compose up -d local

```

  

2. Compile a versão local

  

```bash

go build

```

  

3. Servindo o DraGonT localmente

  

```

./apc-api serve --port 8080 \
  --mongo-host localhost \
  --mongo-port 27017 \
  --jwt-key SUPER_SECRET \
  --codeforces-key f3d968eea83ad8d5f21cad0365edcc200439c6f0 \
  --codeforces-secret b30c206b689d5ba004534c6780aa7be8e234a7f3 \
  --log-level debug

```


# Product backlog

- Modelagem do banco

  

- Login e senha

  

- Integração com Online Judges

  

- Dash board individual

  

- Processar dados do PJudge

  

- Análise de dados para gerar dash boards

  

- Dash board da turma

  

# Planejamento

  

## Semana 1~2 ( 1 ~ 15 abril)

  

- Estudar a linguagem (familiarizar)

  

- Pesquisar possíveis tecnologias (Frameworks tudo mais)

  

## Semana 3~4 ( 15 ~ 29 abril)

  

- Modelagem do banco

  

- Design das interfaces (esboço)

  

## Semana 5~6 (29 abril ~ 13 maio)

  

- Implementar testes das classes de BD

  

- Implementar BD e classes

  

- Fazer telas de login/senha

  

## Semana 7 (13 ~ 20 maio)

  

- Migrar as aulas que estão no GitHub

  

## Semana 8~9 (20 ~ 3 de junho)

  

- Analisar possíveis judges (CodeForces é certeza)

  

- Integrar as APIs de online Judges externos

  

## Semana 10 (3 ~ 10 junho)

  

- Processamento automático do resultado de provas do PJudge

  

- Integrar no sistema

  

> Neste ponto já temos um sistema funcional para os alunos

  

## Semana 11 (10 ~ 17 junho)

  

- Revisão da segurança do projeto

  

- Teste Alpha

  

## Semana 12~13 (17 junho ~ 1 julho)

  

- Depoly do projeto

  

- Testes Beta

  

> Fim do primeiro semestre

  

## Semana 1 (19 agosto ~ 26 agosto)

  

- Revisão do conteúdo / Questões (Pedido anteriormente)

- Refatoração da API

  

## Semana 2 (26 agosto ~ 2 setembro)

  

- Implementação servidor de autenticação

- Finalização CRUD alunos e monitores

  

## Semana 3~5 (2 ~ 23 setembro)

  

- Implementação envio de provas, trabalhos e exercícios extras (Exercícios de arquivos)

- Deploy aplicação para acompanhamento

  

## Semana 6 (23 setembro ~ 30 setembro)

  

- Semana Universitária

  

## Semana 7~8 (30 setembro ~ 7 outubro)

  

- Integração das aulas práticas

  

## Semana 8 (7 outubro ~ 14 outubro)

  

- Analise e estruturação dos dados da turma

  

## Semana 9~11 (14 outubro ~ 4 novembro)

  

- Implementação das analises dos dados e do dashboard

  

## Semana 12 (4 novembro ~ 11 novembro)

  

- Analise da estrutura do relatório automático

  

## Semana 13~15 (11 novembro ~ 2 dezembro)

  

- Implementando geração de relatório automático

  

## Semana 16 (2 dezembro ~ 9 dezembro)

  

- Análise da segurança

- Estudo sobre deploy oficial da aplicação

  
  

> Versão 100% funcional

  
  

.

  

.

  

> possível dilatação do tempo de projeto + novas ideias


# MIT License

Copyright (c) 2019 DraGonT

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

