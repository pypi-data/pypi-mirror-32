# haproxysubdomains

Manages redirect rules in haproxy configuration based on subdomain acl.

Watch out, this will remove any comment in the haproxy config.

Uses [`pyhaproxy`](https://github.com/imjoey/pyhaproxy) under the
hood.

## Usage
```
$ cat test.cfg
```

```
frontend http
    bind *:80


frontend https
    bind *:443 ssl crt /etc/ssl/my_cert/
    default_backend default


backend default
    server default 127.0.0.1:8888
```

Add a redirect rule:
```
$ haproxysubdomains add test.cfg https mydomain.com subdomain 5432

$ cat test.cfg
```

```
frontend http
    bind *:80


frontend https
    bind *:443 ssl crt /etc/ssl/my_cert/
    default_backend default
    acl subdomain hdr(host) -i subdomain.mydomain.com
    use_backend subdomain if subdomain


backend default
    server default 127.0.0.1:8888


backend subdomain
    server subdomain 127.0.0.1:5432
```

Remove a redirect rule:
```
$ haproxysubdomains del test.cfg https subdomain
```
