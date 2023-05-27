from constants import *
import requests
from bs4 import BeautifulSoup


class Scraper:
	def __init__(self, source):
		self.source = source
		if source == METROGRAPH:
			self.url = 'https://metrograph.com/nyc/'
		elif source == ANGELIKA:
			self.url = 'https://www.angelikafilmcenter.com/nyc/showtimes-and-tickets/now-playing'
		elif source == IFC_CENTER:
			self.url = 'https://www.ifccenter.com/'
		elif source == LINCOLN_CENTER:
			self.url = 'https://www.filmlinc.org/wp-content/themes/filmlinc/api-events.php?start=2023-04-30&end=2023-06-04'
		elif source == ANTHOLOGY_FILM_ARCHIVES:
			self.url = 'http://anthologyfilmarchives.org/film_screenings/calendar?view=list#day-24'


	def scrape_website(self):
		if self.source == METROGRAPH:
			response = requests.get(self.url)
			soup = BeautifulSoup(response.text, 'html.parser')
			data = []
			movies = soup.find_all('div', class_='col-sm-12 homepage-in-theater-movie')
			for movie in movies:
				movie_data = {}
				h5_tags = movie.find_all('h5')
				movie_title = movie.find('h3', class_='movie_title')
				movie_data['h5_tags'] = [tag.text for tag in h5_tags]
				movie_data['h5_links'] = [tag.a.get('href') for tag in h5_tags if tag.a]
				movie_data['movie_title'] = movie_title.text if movie_title else None
				data.append(movie_data)
			return data
		elif self.source == IFC_CENTER:
			response = requests.get(self.url)
			soup = BeautifulSoup(response.text, 'html.parser')
			weekdays = ['mon', 'tue', 'wed', 'thu', 'fri', 'sat', 'sun']
			data = []
			for weekday in weekdays:
				daily_schedule = soup.find_all('div', class_=f"daily-schedule {weekday}")
				for day_schedule in daily_schedule:
					h3_tags = day_schedule.find_all('h3')
					data.extend([h3_tag.text for h3_tag in h3_tags])

			return data
		elif self.source == LINCOLN_CENTER:
			
			headers = {
				"Accept": "application/json, text/javascript, */*; q=0.01",
				"Accept-Language": "en-GB,en-US;q=0.9,en;q=0.8",
				"Cookie": "BCwFZaUJYkVNOQvo=7%2Ar8.EgHypXs; zXRJOWcyp=FMpvhLWAz.yU65l; JiWcVHmyUqolnK=gKm0S1J_uYEfs; qRAaFTG=%5B4j%2A1s; kifcc_null=1",
				"Referer": "https://www.filmlinc.org/calendar/",
				"Sec-Ch-Ua": '"Google Chrome";v="113", "Chromium";v="113", "Not-A.Brand";v="24"',
				"Sec-Ch-Ua-Mobile": "?0",
				"Sec-Ch-Ua-Platform": "macOS",
				"Sec-Fetch-Dest": "empty",
				"Sec-Fetch-Mode": "cors",
				"Sec-Fetch-Site": "same-origin",
				"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36",
				"X-Requested-With": "XMLHttpRequest",
			}

			response = requests.get(self.url, headers=headers)

			# Here you can check the response
			return response.json()
		

		elif self.source == ANGELIKA:
			response = requests.get(self.url)
			soup = BeautifulSoup(response.text, 'html.parser')
			films = soup.find_all('div', {'class': 'film status-now_playing'})
			data = []
			for film in films:
				a_href = film.find('a')['href']
				name = film.find('h2', {'class': 'name'}).get_text(strip=True)
				desc = film.find('div', {'class': 'desc'}).p.get_text(strip=True)
				data.append({'name': name, 'desc': desc, 'a_href': a_href})
			return data

				
		elif self.source == ANTHOLOGY_FILM_ARCHIVES:
			response = requests.get(self.url)
			soup = BeautifulSoup(response.text, 'html.parser')
			films = soup.find_all('div', class_='film-showing clearfix')
			data = []
			for film in films:
				film_title = film.find('span', {'class': 'film-title'}).text
				br_siblings = film.find('span', {'class': 'film-title'}).find_next_siblings(text=True)
				br_text = [text for text in br_siblings if str(text).strip()]
				data.append({'film_title': film_title, 'br_text': br_text})
			return data
			
s = Scraper('anthology_film_archives')
metadata = s.scrape_website()
for movie in metadata:
	print(movie)

