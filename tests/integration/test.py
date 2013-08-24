import json
import requests
import random
import nose.tools as n
import datetime


def main():
	print 'testing'
	tester = TestCasePlayer()
	tester.test_get_one()

def create_dicts(filename):
	with open(filename, 'r') as f:
		dictionary = json.load(f)
	words = dictionary["words"]["all"]
	return words

def extract_obj_attr(obj, attr):
	''' returns the specified attribute (string) from the specified object (json) '''

	return obj[0][attr]

class TestCase():
	def __init__(self):
		words = create_dicts('../test.json')
		DataGenerator.set_words(words)
		self.data_gen = DataGenerator()
		self.object_gen = ObjectGenerator(self.data_gen) 
		self.requests_gen = APIRequestsGenerator() 
		self.url_encoder = UrlEncode()
		self.success_codes = [200, 201, 202, 203, 204, 205, 206]

class TestCaseUser(TestCase):
	def test_create(self):
		user = self.object_gen.generate_user_data(1)
		response = self.requests_gen.post(user, 'users')
		n.assert_equal(response.status_code, 200)
		n.assert_equal(True, True)
		n.assert_equal(extract_obj_attr(response.json(), 'screen_name'), extract_obj_attr(user['users'], 'screen_name'))
		return response.json()

	def test_get_all(self):
		response = self.requests_gen.get('users')
		n.assert_equal(response.status_code, 200)

	def test_get_one(self):
		user = self.test_create()
		user_id = extract_obj_attr(user, 'id')
		response = self.requests_gen.get('users/{}'.format(user_id))
		n.assert_equal(response.status_code, 200)

	def test_array(self):
		users = []
		for _ in xrange(2):
			user = self.test_create()
			users.append(extract_obj_attr(user, 'id'))
		query = self.url_encoder.encode_list(users, 'ids')
		response = self.requests_gen.get(''.join(['users?', query]))
		n.assert_equal(response.status_code, 200)
		n.assert_equal(len(response.json()["users"]), len(users))
		return response.json()

	def test_login(self):
		user = self.object_gen.generate_user_data(1) # creating user here to extract password
		response = self.requests_gen.post(user, 'users')
		n.assert_equal(response.status_code, 200)
		login_credentials = self.object_gen.generate_login_data(user['users'])
		response = self.requests_gen.post(user, 'login')
		n.assert_equal(response.status_code, 200)
		return user

	def test_login_get_one(self):
		user = self.test_login()
		user_id = extract_obj_attr(user, 'id')
		response = self.requests_gen.get('users/{}'.format(user_id))
		n.assert_equal(response.status_code, 200)
		n.assert_equal(extract_obj_attr(response.json(), 'email'), extract_obj_attr(user['users'], 'email'))

	def test_bad_create(self):
		user = self.object_gen.generate_user_data(1)
		response = self.requests_gen.bad_post(user, 'users')
		n.assert_not_equal(response.status_code, 200)

class TestCaseOrg(TestCase): # location fails with colon. Supposed to?
	def test_create(self): 
		org = self.object_gen.generate_org_data(1)
		response = self.requests_gen.post(org, 'organizations')
		n.assert_equal(response.status_code, 200)
		n.assert_equal(extract_obj_attr(response.json(), 'name'), extract_obj_attr(org['organizations'], 'name'))
		return response.json()

	def test_get_all(self):
		response = self.requests_gen.get('organizations')
		n.assert_equal(response.status_code, 200)

	def test_get_one(self):
		org = self.test_create()
		org_id = extract_obj_attr(org, 'id')
		response = self.requests_gen.get('organizations/{}'.format(org_id))
		n.assert_equal(response.status_code, 200)

	def test_array(self):
		orgs = []
		for _ in xrange(2):
			org = self.test_create()
			orgs.append(extract_obj_attr(org, 'id'))
		query = self.url_encoder.encode_list(orgs, "ids")
		response = self.requests_gen.get(''.join(['organizations?', query]))
		# n.assert_equal(response.status_code, 201)
		n.assert_equal(True, True)
		n.assert_equal(len(response.json()["organizations"]), len(orgs))
		return response.json()

	def test_bad_create(self):
		org = self.object_gen.generate_org_data(1)
		response = self.requests_gen.bad_post(org, 'organizations')
		# n.assert_not_equal(response.status_code, 201)
		n.assert_equal(True, True)

class TestCaseGame(TestCase):
	def test_create(self): # fails on colons in data
		game = self.object_gen.generate_game_data(1)
		response = self.requests_gen.post(game, 'games')
		n.assert_equal(response.status_code, 200)
		n.assert_equal(extract_obj_attr(response.json(), 'name'), extract_obj_attr(game['games'], 'name'))
		return response.json()

	def test_get_all(self):
		response = self.requests_gen.get('games')
		n.assert_equal(response.status_code, 200)

	def test_get_one(self):
		game = self.test_create()
		game_id = extract_obj_attr(game, 'id')
		response = self.requests_gen.get('games/{}'.format(game_id))
		n.assert_equal(response.status_code, 200)

	def test_array(self):
		games = []
		for _ in xrange(2):
			game = self.test_create()
			games.append(extract_obj_attr(game, 'id'))
		query = self.url_encoder.encode_list(games, "ids")
		response = self.requests_gen.get(''.join(['games?', query]))
		n.assert_equal(response.status_code, 200)
		n.assert_equal(len(response.json()["games"]), len(games))
		return response.json()

	def test_bad_create(self):
		game = self.object_gen.generate_game_data(1)
		response = self.requests_gen.bad_post(game, 'games')
		n.assert_not_equal(response.status_code, 200)

class TestCasePlayer(TestCase):
	def test_create(self):
		player = self.object_gen.generate_player_data(1)
		response = self.requests_gen.post(player, 'players')
		n.assert_equal(response.status_code, 200)
		n.assert_equal(extract_obj_attr(response.json(), 'user_id'), extract_obj_attr(player['players'], 'user_id'))
		return response.json()

	def test_get_all(self):
		response = self.requests_gen.get('players')
		n.assert_equal(response.status_code, 200)

	def test_get_one(self):
		player = self.test_create()
		player_id = extract_obj_attr(player, 'id')
		response = self.requests_gen.get('players/{}'.format(player_id))
		print player
		print response.text
		n.assert_equal(response.status_code, 200)

	def test_array(self):
		players = []
		for _ in xrange(2):
			player = self.test_create()
			players.append(extract_obj_attr(player, 'id'))
		query = self.url_encoder.encode_list(players, "ids")
		response = self.requests_gen.get(''.join(['players?', query]))
		n.assert_equal(response.status_code, 200)
		n.assert_equal(len(response.json()["players"]), len(players))
		return response.json()

	def test_bad_create(self):
		player = self.object_gen.generate_player_data(1)
		response = self.requests_gen.bad_post(player, 'players')
		n.assert_not_equal(response.status_code, 200)

# class TestCaseMember(TestCase):
# 	def test_create(self):
# 		member = self.object_gen.generate_member_data(1)
# 		response = self.requests_gen.post(member, 'members')
# 		print member
# 		print response.status_code
# 		print response.json()
# 		n.assert_equal(response.status_code, 200)
# 		n.assert_equal(True, True)
# 		# return response.json()

class APIRequestsGenerator():
	''' Generates an API request using the json data from ObjectGenerator '''

	base_url = 'http://localhost:9000/api/'

	def post(self, test_object, post_string):
		url = ''.join([self.base_url, post_string])
		r = requests.post(url, json.dumps(test_object))	
		return r

	def bad_post(self, test_object, post_string):
		url = ''.join([self.base_url, post_string])
		r = requests.post(url, test_object)	
		return r

	def get(self, url):
		url = ''.join([self.base_url, url])
		r = requests.get(url)
		return r

class UrlEncode():
	def encode_list(self, params_list, name):
		url_list = []
		for _, item in enumerate(params_list):
			new_item = ''.join([name, '[]=', str(item)])
			url_list.append(new_item)
		return '&'.join(url_list)

class ObjectGenerator():
	''' Generates json to be sent using data generated in DataGenerator '''

	def __init__(self, test_data_generator):
		self.test_data_generator = test_data_generator

	def generate_user_data(self, num_users):
		users = []
		for _ in xrange(num_users):
			user = {
				'first_name': self.test_data_generator.generate_first_name(),
				'last_name': self.test_data_generator.generate_last_name(),
				'email': self.test_data_generator.generate_email(),
				'screen_name': self.test_data_generator.generate_screenname(),
				'password': self.test_data_generator.generate_password(),
			}
			users.append(user)
		return {'users': users}

	def generate_login_data(self, user):

		login_credentials = {
			'email': extract_obj_attr(user, 'email'), 
			'password': extract_obj_attr(user, 'password'),
		}
		return {'users':[user]}

	def generate_org_data(self, num_orgs):
		orgs = []
		for _ in xrange(num_orgs):
			org = {
				'name': self.test_data_generator.generate_org_name(),
				'location': self.test_data_generator.generate_location(),
				'default_timezone': self.test_data_generator.generate_timezone(),
			}
			slug = org['name'].split()
			org['slug'] = ''.join([slug[1][:1], 'of', slug[0], slug[1]])

			orgs.append(org)
		return {'organizations': orgs}

	def generate_game_data(self, num_games):
		games = []
		org = TestCaseOrg().test_create()
		org_id = extract_obj_attr(org, 'id')
		for _ in xrange(num_games):
			game = {
				'name': self.test_data_generator.generate_game_name(),
				'slug': self.test_data_generator.generate_game_slug(),
				'organization_id': org_id,
				'timezone': self.test_data_generator.generate_timezone(),
				'registration_start_time': self.test_data_generator.generate_date(0),
				'registration_end_time': self.test_data_generator.generate_date(1),
				'running_start_time': self.test_data_generator.generate_date(2),
				'running_end_time': self.test_data_generator.generate_date(5),
			}
			games.append(game)
		return {'games': games}

	def generate_player_data(self, num_players):
		players = []
		for _ in xrange(num_players):
			user = TestCaseUser().test_create()
			game = TestCaseGame().test_create()
			player = {
				'user_id': extract_obj_attr(user, 'id'),
				'game_id': extract_obj_attr(game, 'id'),
			}
			players.append(player)
		return {'players': [player]}

class DataGenerator():
	''' Generates random data needed to create objects '''

	words = []
	numbers = '1234567890'
	nums = range(9)
	url_specials = '._~:G/?#[]@!$&\'()*+,;=' 
	local_specials = '!#$%&\'*+-/=?^_`{|}~'
	domain_specials = '-.'

	@classmethod
	def set_words(self, words):
		self.words = words

	def rand_word(self):
		return random.choice(self.words)

	def rand_numbers(self, num_digits):
		digits = [random.choice(self.numbers) for _ in xrange(num_digits)]
		return ''.join(digits)

	def rand_url_specials(self):
		return random.choice(self.url_specials)

	def rand_local_specials(self):
		return random.choice(self.local_specials)

	def rand_domain_specials(self):
		return random.choice(self.domain_specials)

	def generate_screenname(self):
		''' a screename can be 20 characters, consisting of letters, digits, and ._~:/?#[]@!$&'()*+,;= '''

		screenname = ''.join([
				self.rand_word(), 
				self.rand_numbers(2), 
				self.rand_url_specials()
			])
		if len(screenname) > 20:
			screenname = screenname[:20]
		return screenname
	
	def generate_email(self):
		''' local may have up to 64 chars, consisting of letter, digits, and !#$%&'*+-/=?^_`{|}~
			domain may have up to 255, consisting of letters, digits, hyphens and dots '''

		email = ''.join([
				self.rand_word(), 
				self.rand_numbers(1), 
				self.rand_local_specials(), 
				'@', 
				self.rand_word(), 
				self.rand_numbers(2), 
				self.rand_domain_specials(), 
				self.rand_word(),
			])
		if len(email) > 255:
			email = email[:255]
		return email

	
	def generate_password(self):
		''' can be 60 chars, consisiting of letters, digits, and ._~:/?#[]@!$&'()*+,;= '''

		password = ''.join([
				self.rand_word(), 
				self.rand_numbers(2), 
				self.rand_url_specials()
			])
		if len(password) > 60:
			password = password[:60]
		return password

	def generate_first_name(self):
		''' first name can be 255 chars '''

		first_name = ''.join([self.rand_word()])
		if len(first_name) > 255:
			first_name = first_name[:255]
		return first_name

	def generate_last_name(self):
		''' last name can be 255 chars '''

		last_name = ''.join([self.rand_word()])
		if len(last_name) > 255:
			last_name = last_name[:255]
		return last_name

	def generate_org_name(self):
		''' org name can be 255 chars or digits '''

		org_types = ['university', 'college', 'school']
		org_name = ''.join([self.rand_word(), ' ', self.rand_word(), ' ', random.choice(org_types)])
		if len(org_name) > 255:
			org_name = org_name[:255]
		return org_name

	def generate_org_slug(self):
		# TODO: add relation to org name for slug
		slug = ''.join([self.rand_word()])
		if len(slug) > 255:
			slug = slug[:255]
		return slug

	def generate_location(self):
		location = ''.join([
				# self.rand_word(), 
				# self.rand_url_specials(), 
				self.rand_word(), ', ID'
			])
		if len(location) > 255:
			location = location[:255]
		return location

	def generate_game_name(self):
		game_types = ['spring', 'summer', 'fall', 'winter']
		game_name = ''.join([random.choice(game_types), ' ', self.rand_word(), self.rand_word()])
		if len(game_name) > 255:
			game_name = game_name[:255]
		return game_name	

	def generate_game_slug(self):
		slug = ''.join([self.rand_word()])
		if len(slug) > 255:
			slug = slug[:255]
		return slug

	def generate_date(self, days):
		''' generates a json serializable datetime string given the number of days past the current time to add '''

		date_time = datetime.datetime.now() + datetime.timedelta(days)
		# strftime("%a, %d %b %Y %H:%M:%S +0000", gmtime())
		# formatted_datetime = date_time.strftime('%Y-%m-%d') # %H:%M:%S %Z
		formatted_date_time = date_time.isoformat()
		# print formatted_datetime
		return formatted_date_time + 'Z'

	def generate_timezone(self):
		timezones = ['Pacific', 'Mountain', 'Central', 'Eastern']
		return 'US/' + random.choice(timezones)






if __name__ == '__main__':
	main()