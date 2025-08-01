from indieflix.deng.batch.metrograph.dataclass.metrograph import Metrograph
from indieflix.deng.utils.db import PostgresClient
from indieflix.deng.utils.http_sessions import create_metrograph_session
from bs4 import BeautifulSoup
from datetime import datetime, timezone
from zoneinfo import ZoneInfo
from urllib.parse import urlparse, parse_qs
import pandas as pd 

CURRENT_YEAR = 2025

def extract_vista_film_id(html: str) -> str | None:
    """
    Extracts the vista_film_id from a given HTML snippet containing an <a> tag.

    Parameters:
        html (str): A string of HTML containing the <a> tag with the film ID.

    Returns:
        str or None: The extracted vista_film_id, or None if not found.
    """
    a_tag = html.find("a")

    if a_tag and a_tag.has_attr("href"):
        href = a_tag["href"]
        query_params = parse_qs(urlparse(href).query)
        return query_params.get("vista_film_id", [None])[0]

    return None


def parse_datetime(input_str: str, year: int = None) -> datetime:
    print(input_str)
    if year is None:
        year = datetime.now().year
    dt = datetime.strptime(input_str, "%A %B %d %I:%M%p")
    dt = dt.replace(year=year)
    dt = dt.replace(tzinfo=ZoneInfo("America/New_York"))
    return dt

def parse_html(movie):
    try:
        name = movie.find('h3', class_='movie_title').find('a').text
        showtime_datetime = movie.find('div', class_='film_day').text
        showtime_datetime = parse_datetime(showtime_datetime, 2025)
        date = showtime_datetime.date()
        start_time = showtime_datetime.time()
        datetime_utc = showtime_datetime.astimezone(ZoneInfo("UTC")).timestamp()
        h5s = movie.find_all('h5')
        director_name = h5s[1].text.split(":")[1]
        release_year = h5s[2].text.split("/")[0]
        runtime_mins = h5s[2].text.split("/")[1]
        video_format = h5s[2].text.split("/")[2]
        vista_id = extract_vista_film_id(h5s[3])
        desc = movie.find('p', class_='synopsis').text
        film = Metrograph(
            title=name,
            date=date,
            start_time=start_time,
            datetime_utc=datetime_utc,
            director_name=director_name,
            release_year=release_year,
            runtime_min=runtime_mins,
            video_format=video_format,
            vista_film_id=vista_id,
            desc=desc,
        )
        return film 
    except Exception as e:
        print("failed with error: ", e)
        return None 

def main():
    url = "https://metrograph.com/nyc/"
    session = create_metrograph_session()
    r = session.get(url)
    if r.status_code == 200:
        soup = BeautifulSoup(r.text, "html.parser")
        films = []
        movies = soup.find_all('div', class_='col-sm-12 homepage-in-theater-movie')
        for movie in movies:
            film = parse_html(movie)
            if film:
                films.append(film)

        films = pd.DataFrame(data=films)
        films["datetime_utc"] = films["datetime_utc"].apply(
            lambda x: datetime.fromtimestamp(x, tz=timezone.utc)
        )
        with PostgresClient() as engine:
            # Convert DataFrame rows to list of tuples
            film_records = list(films.itertuples(index=False, name=None))  # [(title1, year1), (title2, year2), ...]

            insert_query = """
                INSERT INTO public.metrograph_provider (title, date, start_time, datetime_utc, director_name, release_year, runtime_min, video_format, vista_film_id, "desc")
                VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s)
            """

            engine.cursor.executemany(insert_query, film_records)
            engine.conn.commit()
    else:
        print(f"failed req: {r.status_code}")


if __name__ == "__main__":
    try:
        main() 
    except Exception as e:
        raise 
