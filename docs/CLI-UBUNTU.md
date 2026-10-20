# BedrockTool CLI en Ubuntu (RyzeMCBE)

Base upstream: [BedrockTool 1.62.0](https://github.com/bedrock-tool/bedrocktool), licencia GPL-3.0. Este fork repara la autenticación por código de dispositivo en la CLI, sin solicitar contraseñas en el VPS.

## Descargar y ejecutar

Cuando exista una release validada en https://github.com/RyzeMCBE/bedrocktool/releases:

```bash
mkdir -p ~/bedrocktool
cd ~/bedrocktool
# Descarga el archivo Linux amd64 desde los assets de la release
chmod +x bedrocktool-linux-amd64-1.62.0-ryze.1
./bedrocktool-linux-amd64-1.62.0-ryze.1 packs -address=play.mercurymc.club:19132 -capture=false
```

El primer uso sin una cuenta guardada imprime una dirección de Microsoft y un código de dispositivo. Abre esa dirección **en tu teléfono**, inicia sesión con la cuenta Microsoft/Xbox autorizada y escribe el código allí. No introduzcas contraseñas en el VPS ni compartas los tokens de autenticación.

Para ver la ayuda de flags, usa `./bedrocktool-linux-amd64-1.62.0-ryze.1 packs -h` (los flags del subcomando van después de `packs`).

## Diagnóstico

Si se interrumpe antes de conectarse:

```bash
./bedrocktool-linux-amd64-1.62.0-ryze.1 packs -address=play.mercurymc.club:19132 -capture=false
echo "Código de salida: $?"
tail -n 60 bedrocktool.log
```

Los logs y archivos de autenticación pueden contener detalles sensibles: elimina códigos, tokens y datos personales antes de compartirlos.

## Compilación y releases

- Los workflows deben pasar pruebas de autenticación y compilación para Linux amd64, Linux arm64 y Windows amd64.
- El workflow solo publica una release con un tag `v*-ryze.*`; un push normal a `master` no crea ni reescribe tags.
- El fork puede compilar sin el material privado de `git-crypt`. La equivalencia funcional con todos los binarios upstream, especialmente las funciones privadas relacionadas con packs, **no está verificada**.
- `update` consulta releases de RyzeMCBE, no sobrescribe esta versión con una release del proyecto original.
- Los tests de CLI no sustituyen las pruebas de inicio de sesión Microsoft y descarga de resource packs en un servidor real.
