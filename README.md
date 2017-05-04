# What's my IP?

Tiny script to run on a server returning the public IP of a requesting client. Exposes a minimal endpoint reachable via HTTP. Allows for checking one's own public IP with little overhead. This is useful to quickly determine one's external IP in the common case one is usually working behind a NAT router. Deploy behind a TLS-encrypted route.


## Usage

```
Usage of ./whats-my-ip:
  -ip string
        Define this service's IP to operate on. (default "0.0.0.0")
  -port string
        Define the port this service is supposed to listen on. (default "27300")
```


## License

This project is [GPLv3](https://github.com/numbleroot/whats-my-ip/blob/master/LICENSE) licensed.
