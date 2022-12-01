## _Environnement de developpement Go 1.17_

## Contruire le contener
```bash
[Windows] docker build -t docker-go:1.00 .
[Linux] sudo docker build -t docker-go:1.00 .
 ```

 ## Lancer le conteneur
```bash 
[Windows] docker run -it --rm --name "docker-go_1.00" -p 80:80 -v "%cd%/go":/go docker-go:1.00
[Linux] sudo docker run -it --rm --name "docker-go_1.00" -p 80:80 -v "`pwd`/go":/go docker-go:1.00
```