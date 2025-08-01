from prefect import flow
from indieflix.deng.batch.metrograph.metrograph import main as metrograph_ingestion


@flow
def daily():
    metrograph_ingestion()

if __name__ == "__main__":
    daily()