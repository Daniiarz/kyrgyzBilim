FROM python:3.9

WORKDIR /usr/src/admin

COPY requirements.txt .

RUN pip install -r requirements.txt

COPY . .

RUN chmod 755 start-server.sh

CMD ["/usr/src/admin/start-server.sh"]