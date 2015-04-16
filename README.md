## PassGo

A Simple password storage solution using RSA and YAML.

Only a fool would use this, I don't know crypto this could be insecure.

PassGo acts as a key/value store backed by yaml, passwords are saved to ~/.password.yml set to 0600, values are base64 encoded.  
your private key is saved to ~/.passgo PEM encoded.

Values can be provided on the command line or piped in.


      Λ projects/passgo bin/passgo hello world
      Λ projects/passgo bin/passgo hello
      world%

      Λ projects/passgo echo "woo" | pbcopy
      Λ projects/passgo pbpaste | bin/passgo wee -
      Λ projects/passgo bin/passgo wee
      woo


build me via `make build`
