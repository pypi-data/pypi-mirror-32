# Riddle 7: Motif
from riddles.riddle import Riddle

class Motif(Riddle):
	
	def __init__(self):
		self.description = 'One small letter, sourrounded by exactly three bodyguards on each of its sides'
		self.hints = 'Go read about "regex".\nList all matches.'
		self.path += "motif\\"
		self.file = self.get_path("bodyguards.txt")
		self.photo = self.get_path("candles.jpg")
		self.status_phrases ={
			'Kingsguard': "Lowercase please.",
			'kings': 'Good, but incomplete.',
			'king': 'Good, but incomplete.',
			'guard': 'Good, but incomplete.',
			'kings guard': 'Omit the space!'
			}
		self.solution = 'kingsguard'
		self.topic = 'Data Processing'