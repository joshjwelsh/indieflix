#!/usr/bin/env bash 

pip install -r requirements.txt
airflow db init --email admin@example.com --firstname admin --lastname admin --password admin --role Admin --username admin 
