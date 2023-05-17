# Documentação
> https://flatbuffers.dev/

# Instalação

```bash
    brew install flatbuffers
```

# Comandos uteis

## Gera arquivo binário para utilizar em qualquer linguagem de programação
```bash
    flatc -b cursos.fbs cursos.txt
```

## Pegar arquivo binário e gerar um JSON
```bash
    flatc --raw-binary -t --strict-json cursos.fbs -- cursos.bin
```
## Gerar o arquivo GO corresponde ao table
```bash
    flatc --go cursos.fbs
```