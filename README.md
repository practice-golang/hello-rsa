# 키젠

## Using tool

* openssl
    * https://www.openssl.org
    * https://wiki.openssl.org/index.php/Binaries

```powershell
$ openssl genrsa -out key.pem 1024
$ openssl rsa -in key.pub -pubout -out key.pem
# dkim
# v=DKIM1;k=rsa;p=value_of_key.pem
```


## DMARC

```
v=DMARC1;p=none;sp=none;pct=100;rua=mailto:iam@domain.com;ruf=mailto:iam@domain.com;ri=86400;aspf=r;adkim=r;fo=1

v=DMARC1;p=none;rua=mailto:iam@domain.com
```
