# RSA keygen

practice

## Using tool

* openssl
    * https://www.openssl.org
    * https://wiki.openssl.org/index.php/Binaries

```powershell
$ openssl genrsa -out key.pem 1024
$ openssl rsa -in pub.pem -pubout -out key.pem
# dkim
# v=DKIM1;k=rsa;p=public_key_value_in_pub.pem
```


## DMARC

```
v=DMARC1;p=none;sp=none;pct=100;rua=mailto:iam@domain.com;ruf=mailto:iam@domain.com;ri=86400;aspf=r;adkim=r;fo=1

v=DMARC1;p=none;rua=mailto:iam@domain.com
```
