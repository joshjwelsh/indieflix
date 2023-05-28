from constants import *
import requests
from bs4 import BeautifulSoup
from datetime import datetime, timedelta


def get_dates():

    current_date = datetime.now()
    dates = []
    for i in range(7):
        dates.append((current_date + timedelta(days=i)).strftime('%Y-%m-%d'))

    return dates

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

		# # disable antholgy film archives for now
		# elif source == ANTHOLOGY_FILM_ARCHIVES:
		# 	self.url = 'http://anthologyfilmarchives.org/film_screenings/calendar?view=list#day-24'


	def scrape_website(self):
		if self.source == METROGRAPH:
			response = requests.get(self.url)
			soup = BeautifulSoup(response.text, 'html.parser')
			data = []
			movies = soup.find_all('div', class_='col-sm-12 homepage-in-theater-movie')
			for movie in movies:
				movie_data = {}
				name = movie.find('h3', class_='movie_title').find('a').text
				showtime_day =movie.find('div', class_='film_day')
				showtime_time = showtime_day.find('a').text
				showtime = (showtime_day.text, showtime_time)
				metadata = {'h5':movie.find_all('h5'),'desc':movie.find('p', class_='synopsis').text}
				movie_data['name'] = name
				movie_data['showtime'] = showtime
				movie_data['metadata'] = metadata
				data.append(movie_data)
			return data
		elif self.source == IFC_CENTER:
			# Problem: This won't read the current days showtimes.
			# Shouldnt be a problem if this is run daily 
			response = requests.get(self.url)
			soup = BeautifulSoup(response.text, 'html.parser')
			weekdays = ['mon', 'tue', 'wed', 'thu', 'fri', 'sat', 'sun']
			data = []
			movie_data = {} 

			for weekday in weekdays:
				daily_schedule = soup.find_all('div', class_=f"daily-schedule {weekday}")
				for day_schedule in daily_schedule:
					date = day_schedule.find('h3').text
					details = day_schedule.find_all('div', class_='details')
					for detail in details:
						movie_name = detail.find('h3').text
						if movie_name in movie_data:
							times = detail.find('ul', class_='times').find('li').find('a').text
							movie_data[movie_name].append((date,times, {}))
						else:
							movie_data[movie_name] = []
			
					data.append(movie_data)
						
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
			
			data = []
			for day in get_dates():
				url = f'{self.url}/{day}'
				response = requests.get(url)
				soup = BeautifulSoup(response.text, 'html.parser')
				films = soup.find_all('div', class_='film status-now_playing')
		

				for film in films:
					name = film.find('div', class_='text').find('a').text
					showtimes_html = film.find('div', class_='movietimes').find_all('form')
					
					showtimes = (day, [showtime.find('input', class_='showtime reserved-seating')['value'] for showtime in showtimes_html])
					metadata = {}
					data.append({'name': name, 'showtimes': showtimes, 'metadata': metadata})
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
			
def output(s: Scraper):
	metadata = s.scrape_website()
	for movie in metadata:
		print(movie)


