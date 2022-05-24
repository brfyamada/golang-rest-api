* Conectando com o banco no docker:
    comando para achar o host do branco:


    ```
    $ docker-compose exec postgres sh
    $ hostname -i

    ```

    ou

    ```
    $ docker ps
    $ docker inspect 'cod of image' | grep IPAddress

    ```

