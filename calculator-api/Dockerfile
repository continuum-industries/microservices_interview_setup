FROM python:3.8.10-slim AS services

RUN apt-get update \
    && apt-get install -y make netcat wget supervisor \
    && apt-get clean

WORKDIR /calculator-api

COPY ./calculator-api/requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
COPY shared ./shared

RUN python shared/setup.py install
RUN pip install shared/

COPY ./calculator-api .

ENTRYPOINT [ "python", "-u", "./src/main.py" ]
