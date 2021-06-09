FROM python:3.9-buster

RUN mkdir /app

WORKDIR /app

COPY . /app

RUN pip install -r /app/requirements.txt

CMD python -m flask run --host=0.0.0.0

