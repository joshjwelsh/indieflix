import requests
from requests.adapters import HTTPAdapter, Retry

class CustomHTTPAdaptor(HTTPAdapter):
        def __init__(self, timeout=None, *args, **kwargs):
                self.timeout = timeout
                super().__init__(*args, **kwargs)

        def send(self, *args, **kwargs):
                kwargs["timeout"] = self.timeout
                return super().send(*args, **kwargs)


def create_metrograph_session(
                timeout=10, 
                num_retries=1,
                backoff_factor=0.1,
                status_forcelist=(404, 500, 503, 503, 504)
):
    session = requests.Session()
    retries = Retry(
                total=num_retries,
                backoff_factor=backoff_factor, 
                status_forcelist=status_forcelist,
        )
    session.mount("http://", CustomHTTPAdaptor(timeout=timeout, max_retries=retries))
    session.mount("https://", CustomHTTPAdaptor(timeout=timeout, max_retries=retries))

    return session 