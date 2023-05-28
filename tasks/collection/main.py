from objects import Scraper, output
from constants import METROGRAPH, ANGELIKA, IFC_CENTER, LINCOLN_CENTER


def main():
        sources = [METROGRAPH, ANGELIKA, IFC_CENTER, LINCOLN_CENTER]
        for source in sources:
                scraper = Scraper(source)
                print(source.upper())
                output(scraper)

if __name__ == '__main__':
        main()