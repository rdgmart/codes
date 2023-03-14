
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

### React Native Firebase
- https://firebase.google.com/?hl=pt-br
- https://rnfirebase.io/
```bash
    yarn add @react-native-firebase/app
```
- Configuration Expo:
> Open file app.json and add:
```
    "plugins":[
      "@react-native-firebase/app"
    ],
    "android": {
      "googleServicesFile": "./google-services.json"
    },
    "ios": {
      "googleServicesFile": "./GoogleService-Info.plist"
    },
```
- Prebuild 
```bash
    npx expo prebuild --npm
```

# Install the authentication and firestore module
```bash
  yarn add @react-native-firebase/app
  yarn add @react-native-firebase/auth
  yarn add @react-native-firebase/firestore
```

# If you're developing your app using iOS:
- Open xcode > build phases > Link binary with Libraries > add hermes.xcframework
- xcode > general > Frameworks, Libraries and Embedded Content > hermes.xcframework > change to  Embed & Sign 
```bash
  yarn install 
  cd ios/ 
  rm -rf Podfile.lock
  pod install --repo-update
  npx react-native run-ios
```
# If you're developing your app using Android:
## Config
```bash
npx react-native info
export ANDROID_HOME=/Users/yourname/Library/Android/sdk
export PATH=$ANDROID_HOME/tools:$PATH
export PATH=$ANDROID_HOME/platform-tools:$PATH
```
## Run Android
```bash
npx react-native run-android
```

# Social Authentication
## Google
- https://github.com/react-native-google-signin/google-signin#project-setup-and-initialization 
- https://docs.expo.dev/guides/google-authentication/
```bash
  yarn add @react-native-google-signin/google-signin
  cd android/app 
  keytool -keystore debug.keystore -list -v
```
- https://console.cloud.google.com/apis/credentials?project=YOUR_PROJECT_FIREBASE
> Configure a chave SHA-1 criando uma credential 'OAuth Client ID'. 

- https://console.cloud.google.com/apis/credentials/consent?project=YOUR_PROJECT_FIREBASE
> HABILITE os e-mails que poderão autenticar na sua aplicação