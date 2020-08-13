# PROJETO: Utilizando K8S
Repositório para entrega do projeto Code Education

# Sessão da Etapa 3 - Desafio Go!
Crie um aplicativo Go que disponibilize um servidor web na porta 8000 que quando acessado seja exibido em HTML (em negrito) Code.education Rocks!
A exibição dessa string deve ser baseada no retorno de uma função chamada "greeting". Essa função receberá a string como parâmetro e a retornará entre as tags < b >  < / b >.
Como ótimo desenvolvedor(a), você deverá criar o teste dessa função.
Ative o processo de CI no Google Cloud Build para garantir que a cada PR criada faça com que os testes sejam executados.
Gere a imagem desse aplicativo de forma otimizada e publique-a no Docker Hub
Utilizando o Kubernetes, disponibilize o serviço do tipo Load Balancer que quando acessado pelo browser acesse a aplicação criada em Go.

# Endereço da imagem gerada no Docker Hub
https://hub.docker.com/repository/docker/samysilveira/gowebsrv