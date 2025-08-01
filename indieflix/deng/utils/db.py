import psycopg2
from dotenv import load_dotenv
import os

class PostgresClient:
    def __init__(self):
        load_dotenv()
        self.conn = psycopg2.connect(
            host=os.getenv("DATABASE_HOST"),
            user=os.getenv("DATABASE_USR"),
            password=os.getenv("DATABASE_PWD"),
            port=os.getenv("DATABASE_PORT"),
            dbname=os.getenv("DATABASE_DB"),
        )
        self.cursor = self.conn.cursor()

    def __enter__(self):
        return self

    def __exit__(self, exc_type, exc_value, traceback):
        self.cursor.close()
        self.conn.close()


if __name__ == "__main__":
    with PostgresClient() as engine:
        engine.cursor.execute("SELECT * FROM public.metrograph_provider;")
        result = engine.cursor.fetchone()
        print(result) 
