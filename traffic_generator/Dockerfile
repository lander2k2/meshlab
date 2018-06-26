FROM python:3

RUN pip install locustio

COPY locust_file.py /
COPY entrypoint.sh /
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]

