
# Ignite Lab - 3ª edição, Evento da Rocketseat ocorrido entre 17 e 20 de julho de 2022.
## Professor Rodrigo Goncalves - https://github.com/rodrigorgtic

<p align="center">
  <img src="https://github.com/rdgmart/Test/blob/develop/react-native/rocket-help/src/assets/layout-rocket-help.png" width="700" alt="Layout rocket-help">
</p>

<div style="display: inline_block"><br>
  <img align="center" alt="Rdgmart-Ts" height="30" width="40" src="https://raw.githubusercontent.com/devicons/devicon/master/icons/typescript/typescript-plain.svg">
  <img align="center" alt="Rdgmart-React" height="30" width="40" src="https://raw.githubusercontent.com/devicons/devicon/master/icons/react/react-original.svg">
</div>


# Configuração Ambiente
## Node
```bash
    nvm install -lts
```

## Expo - https://expo.dev/
```bash
    yarn global add expo-cli 
```

 OU
        
```bash         
    npm install -g expo-cli
```

## Expo Go
- App utilizado para instalar/executar o app em desenvolvimento no device (fisico ou emulador). 
- No emulador é instalado automaticamente.
- Nos dispositivos fisicos instale através das lojas Android ou Apple.

## VSCODE
### Add Extensions
> R Components - https://marketplace.visualstudio.com/items?itemName=rodrigorgtic.rcomponent


## Para criar o projeto
```bash
    expo init NOME-DO-PROJETO
       bare and minimal, just the essentials to get you started
```

### Configurar o typescript no projeto
- Abrir no vscode e trocar o App.js para App.tsx
- Criar na raiz do projeto o arquivo tsconfig.json (sem conteudo) que o typescript utiliza para definir suas configurações
- Ao iniciar o projeto com expo vai identificar as configurações com typescript e vai perguntar se deseja utilizar.
```bash
   expo start
   Yes
```

### Nativebase 
- https://docs.nativebase.io/install-expo
- https://docs.nativebase.io/customizing-theme

### react-native-svg-transformer
- https://github.com/kristerkari/react-native-svg-transformer
```bash
   yarn add --dev react-native-svg-transformer
   config metro.config.js
   config declaration.d.ts
```

### React Navigation
- https://reactnavigation.org/docs/getting-started/
```bash
    yarn add @react-navigation/native
    npx expo install react-native-screens react-native-safe-area-context
    yarn add @react-navigation/native-stack
```