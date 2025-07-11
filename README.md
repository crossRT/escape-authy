# escape-authy

A CLI tool to convert `decrypted_tokens.json` into formats compatible with other MFA applications.

## Intention

While researching how to migrate from Authy to other MFA apps, I found many helpful resources—but the amount of information was overwhelming and the format differences were confusing.

I created this project to help others like me who want a simple, reliable way to convert `decrypted_tokens.json` into a usable format.

## Supported Formats

- [Raivo](https://raivo-otp.com/) - `escape-auth toRaivo`

- [Aegis](https://getaegis.app/) - `escape-auth toAegis`

- [Ente](https://ente.io/auth/) - `escape-auth toEnte`

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
