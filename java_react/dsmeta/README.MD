# DevSuperior - Semana Java React - Evento ocorrido entre os dias 05 a 11 setembro 2022. 
## Professor Nelio Alves - https://github.com/devsuperior/sds-dsmeta

<div> 
  <a href="https://www.youtube.com/devsuperior" target="_blank"><img src="https://img.shields.io/badge/YouTube-FF0000?style=for-the-badge&logo=youtube&logoColor=white" target="_blank"></a>
  <a href="https://www.instagram.com/devsuperior.ig" target="_blank"><img src="https://img.shields.io/badge/-Instagram-%23E4405F?style=for-the-badge&logo=instagram&logoColor=white" target="_blank"></a>
 
</div>

<div style="display: inline_block"><br>
  <img align="center" alt="Rdgmart-Java" height="30" width="40" src="https://raw.githubusercontent.com/devicons/devicon/master/icons/java/java-original.svg">
  <img align="center" alt="Rdgmart-Ts" height="30" width="40" src="https://raw.githubusercontent.com/devicons/devicon/master/icons/typescript/typescript-plain.svg">
  <img align="center" alt="Rdgmart-React" height="30" width="40" src="https://raw.githubusercontent.com/devicons/devicon/master/icons/react/react-original.svg">
</div>


# Roteiro
## Criar o projeto frontend 
> yarn create vite frontend --template react-ts
## Criar o projeto backend
> Spring Initializr - https://start.spring.io/

# Heroku
> Install the Heroku CLI
## Comandos
```bash
heroku -v
heroku login
heroku git:remote -a <nome-do-app>
git remote -v
git subtree push --prefix java_react/dsmeta/backend heroku main
```

## URL teste:
- Sales (https://dsmeta-rdgmart.herokuapp.com/sales)
- Sales by date (https://dsmeta-rdgmart.herokuapp.com/sales?minDate=2022-01-01&maxDate=2022-03-31)
- Notification (https://dsmeta-rdgmart.herokuapp.com/sales/53/notification)

# Deploy no Netlify
- https://www.netlify.com/
- acrescente `window.React = React` no seu main.tsx
- Deploy básico
  - Base directory: java_react/dsmeta/frontend
  - Build command: yarn build
  - Publish directory: java_react/dsmeta/frontend/dist
  - Variáveis de ambiente:
    - VITE_BACKEND_URL --> coloque a url de acesso do backend (heroku nesse caso)

- Configurações adicionais
  - Site settings -> Domain Management: (colocar o nome que você quiser) -> Ex: dsmeta-rdgmart
  - Deploys -> Trigger deploy
