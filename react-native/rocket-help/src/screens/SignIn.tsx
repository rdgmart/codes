import { useEffect, useState } from "react";
import { Alert } from "react-native";
import { VStack, Heading, Icon, useTheme, HStack, IconButton } from "native-base";
import { Envelope, Key, GoogleLogo, FacebookLogo } from "phosphor-react-native";

import auth from '@react-native-firebase/auth'
import { GoogleSignin, statusCodes } from '@react-native-google-signin/google-signin';

import Logo from '../assets/logo_primary.svg';

import { Button } from "../components/Button";
import { Input } from "../components/Input";

export function SignIn() {
    const { colors } = useTheme();

    const [isLoading, setIsLoading] = useState(false);
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    GoogleSignin.configure({
        webClientId:"YOUR_WEB_CLIENT_ID_TYPE_3",
    });

    function handleSignIn() {
        if (!email || !password) {
            return Alert.alert('Entrar', 'Informe e-mail e senha');
        }

        setIsLoading(true);

        auth()
            .signInWithEmailAndPassword(email, password)
            .catch((error) => {
                console.log(error);
                setIsLoading(false);

                if (error.code === 'auth/invalid-email') {
                    return Alert.alert('Entrar', 'Email inválido.');
                }

                if (error.code === 'auth/user-not-found' || error.code === 'auth/wrong-password') {
                    return Alert.alert('Entrar', 'E-mail e/ou senha inválidos.');
                }

                return Alert.alert('Entrar', 'Não foi possível acessar.');
            });
    }

    const handleGoogleLogin = async () => {

        const {idToken} = await GoogleSignin.signIn();

        const googleCredential = auth.GoogleAuthProvider.credential(idToken);

        setIsLoading(true);
        await auth()
            .signInWithCredential(googleCredential)
            .catch((error) => {
                console.log(error);
                setIsLoading(false);

                if (error.code === statusCodes.SIGN_IN_CANCELLED) {
                    return Alert.alert('Entrar Google', 'Processo cancelado pelo usuário');
                } else if (error.code === statusCodes.IN_PROGRESS) {
                    return Alert.alert('Entrar Google', 'Processo de autenticação em andamento...');
                } else if (error.code === statusCodes.PLAY_SERVICES_NOT_AVAILABLE) {
                    return Alert.alert('Entrar Google', 'Serviço de autenticação indisponível. Verifique sua conexão ou tente mais tarde.');
                } 

                return Alert.alert('Entrar Google', 'Não foi possível acessar.');
            });
    }

    const handleFacebookLogin = async () => {
        return Alert.alert('Entrar Facebook', 'Ainda está pendente...');
    }

    return (
        <>
            <VStack flex={1} alignItems="center" bg="gray.600" px={8} pt={24} >
                <Logo />
                <Heading color="gray.100" fontSize="xl" mt={20} mb={6}>
                    Acesse sua conta
                </Heading>

                <Input
                    mb={4}
                    placeholder="E-mail"
                    InputLeftElement={<Icon as={<Envelope color={colors.gray[300]} />} ml={4} />}
                    onChangeText={setEmail}
                />

                <Input
                    mb="8"
                    placeholder="Senha"
                    InputLeftElement={<Icon as={<Key color={colors.gray[300]} />} ml={4} />}
                    secureTextEntry
                    onChangeText={setPassword}
                />
                <HStack>
                    <Button title="Entrar" w="full" onPress={handleSignIn} isLoading={isLoading} />

                </HStack>
                <HStack>
                    <VStack flex={1} alignItems="center" bg="gray.600" px={8} pt={5} >
                        <Heading color="gray.100" fontSize="sm" mb={6}>
                            Ou entre utilizando
                        </Heading>
                        <HStack>
                        <IconButton
                            mr={20}
                            icon={<GoogleLogo size={48} color={colors.gray[300]} />}
                            onPress={handleGoogleLogin}
                        />
                        <IconButton
                            icon={<FacebookLogo size={48} color={colors.gray[300]} />}
                            onPress={handleFacebookLogin}
                        />

                        </HStack>
                    </VStack>

                </HStack>

            </VStack>



        </>
    )

}