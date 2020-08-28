#!/bin/bash

#Baixando e instalando pre-requisitos
sudo apt-get install apt apt-transport-https wget curl -y;

#Adicionando os repositorios do PostgreSQL e do pgAdmin
sudo sh -c 'echo "deb [arch=amd64] https://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list'
sudo sh -c 'echo "deb [arch=amd64] https://ftp.postgresql.org/pub/pgadmin/pgadmin4/apt/$(lsb_release -cs) pgadmin4 main" > /etc/apt/sources.list.d/pgadmin4.list && apt update'

#Adicionando a chave do PostgreSQL e do pgAdmin
wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -
curl https://www.pgadmin.org/static/packages_pgadmin_org.pub | sudo apt-key add

#Atualizando os repositorios apt
sudo apt update

#Instalando o PostgreSQL e do pgAdmin
sudo apt install postgresql pgadmin4 pgadmin4-desktop pgadmin4-web -y;

#Configurando o servidor web do pgAdmin
sudo /usr/pgadmin4/bin/setup-web.sh
