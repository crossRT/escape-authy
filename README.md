# escape-authy

A CLI tool to convert `decrypted_tokens.json` into formats compatible with other MFA applications.

## Intention

I found many helpful resources while researching how to migrate from Authy to other MFA apps.

Especially to [this guide](https://github.com/AlexTech01/Authy-iOS-MiTM) created by @AlexTech01.

I managed to retrieve the `decrypted_tokens.json`, but the format differences on different MFA applications were confusing for me.

And some of the scripts in the thread are outdated, and some requires the Authy Desktop app which is not installed on my machine.

So I created this project to help others like me who want a simple, reliable way to convert `decrypted_tokens.json` into a usable format without the need of Authy Desktop app.

## Supported Formats

- [Raivo](https://raivo-otp.com/) - `escape-authy toRaivo`

- [Aegis](https://getaegis.app/) - `escape-authy toAegis`

- [Ente](https://ente.io/auth/) - `escape-authy toEnte`

## How to use

> ⚠️ Make sure you've already retrieved the `decrypted_tokens.json` file by following [this guide](https://github.com/AlexTech01/Authy-iOS-MiTM).


1. Download the binary from the [releases](https://github.com/crossRT/escape-authy/releases) page and extract locally.

   If you prefer commands:
   ```bash
   VERSION=v0.0.1
   curl -L https://github.com/crossRT/escape-authy/releases/download/$VERSION/escape-authy_Darwin_arm64.tar.gz | tar -xz
   ```

2. Run the command below (replace with your actual path):
   ```bash
   ./escape-authy --decrypted-tokens-file-path ~/Downloads/decrypted_tokens.json toRaivo
   ```

3. A converted file will be exported to the same directory.

4. Import the exported file into your preferred MFA app.

## Notes

- 2FSA supports importing from `raivo` and `aegis` formats.


## References

- https://github.com/AlexTech01/Authy-iOS-MiTM

- https://www.tygertec.com/aegis-raivo-otp-translator/

- https://help.ente.io/auth/migration-guides/authy/
