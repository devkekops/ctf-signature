# ctf-signature

docker build -t ctf-signature .  

docker run -e SERVER_ADDRESS='0.0.0.0:80' -e SECRET_KEY='top_secret' -e FLAG='...' -p 80:80 ctf-signature
