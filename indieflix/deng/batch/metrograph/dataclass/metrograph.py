from dataclasses import dataclass
from datetime import datetime


@dataclass
class Metrograph:
        title: str 
        date: datetime.date
        start_time: datetime.timestamp
        datetime_utc: datetime.timestamp 
        director_name: str 
        release_year: str 
        runtime_min: str 
        video_format: str 
        vista_film_id: str 
        desc: str 




        