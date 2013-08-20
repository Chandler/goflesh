import json
import requests
import random
import nose.tools as n
import urllib
import datetime


def main():
	print 'testing'
	
def create_dicts(filename):
	with open(filename, 'r') as f:
		dictionary = json.load(f)
	words = dictionary["words"]["all"]
	return words

class TestCase():
	def __init__(self):
		words = create_dicts('../test.json')
		DataGenerator.set_words(words)
		self.data_gen = DataGenerator()
		self.object_gen = ObjectGenerator(self.data_gen) 
		self.api_gen = APIRequestsGenerator() 
		self.url_encoder = UrlEncode()

class TestCaseUser(TestCase):
	def test_create(self):
		user = self.object_gen.generate_user_data()
		response = self.api_gen.post(user, 'users')
		n.assert_equal(response.status_code, 200)
		return response.json()

	def test_get_all(self):
		user = self.object_gen.generate_user_data()
		response = self.api_gen.get('users')
		n.assert_equal(response.status_code, 200)

	def test_get_one(self):
		user = self.test_create()
		user_id = str(user[0]["id"])
		response = self.api_gen.get(''.join(['users/', user_id]))
		n.assert_equal(response.status_code, 200)

	def test_array(self):
		users = []
		for _ in range(5):
			user = self.test_create()
			users.append(str(user[0]["id"]))
		query = self.url_encoder.encode_list(users, "ids")
		response = self.api_gen.get(''.join(['users?', query]))
		n.assert_equal(response.status_code, 200)
		n.assert_equal(len(response.json()["users"]), len(users))

	def test_login(self):
		user = self.object_gen.generate_user_data()
		response = self.api_gen.post(user, 'users')
		n.assert_equal(response.status_code, 200)
		login_credentials = self.object_gen.generate_login_data(user)
		response = self.api_gen.post(user, 'login')
		n.assert_equal(response.status_code, 200)
		

	def test_bad_create(self):
		user = self.object_gen.generate_user_data()
		response = self.api_gen.bad_post(user, 'users')
		n.assert_not_equal(response.status_code, 200)

class TestCaseOrg(TestCase):
	def test_create(self): # fails on colons in data
		org = self.object_gen.generate_org_data()
		response = self.api_gen.post(org, 'organizations')
		n.assert_equal(response.status_code, 200)
		return response.json()

	def test_get_all(self):
		org = self.object_gen.generate_org_data()
		response = self.api_gen.get('organizations')
		n.assert_equal(response.status_code, 200)

	def test_get_one(self):
		org = self.test_create()
		org_id = str(org[0]["id"])
		response = self.api_gen.get(''.join(['organizations/', org_id]))
		n.assert_equal(response.status_code, 200)

	def test_array(self):
		orgs = []
		for _ in range(5):
			org = self.test_create()
			orgs.append(str(org[0]["id"]))
		query = self.url_encoder.encode_list(orgs, "ids")
		response = self.api_gen.get(''.join(['organizations?', query]))
		n.assert_equal(response.status_code, 200)
		n.assert_equal(len(response.json()["organizations"]), len(orgs))

	def test_bad_create(self):
		org = self.object_gen.generate_org_data()
		response = self.api_gen.bad_post(org, 'organizations')
		n.assert_not_equal(response.status_code, 200)

class TestCaseGame(TestCase):
	def test_create(self): # fails on colons in data
		game = self.object_gen.generate_game_data()
		response = self.api_gen.post(game, 'games')
		n.assert_equal(response.status_code, 200)
		return response.json()

class TestCaseMember(TestCase):
	def test_create(self):
		member = self.object_gen.generate_member_data()
		response = self.api_gen.post(member, 'members')
		n.assert_equal(response.status_code, 200)

class UrlEncode():
	def encode_list(self, params_list, name):
		url_list = []
		length = len(params_list)
		for i,item in enumerate(params_list):
			new_item = ''.join([name, '[]=', item])
			url_list.append(new_item)
		return '&'.join(url_list)

''' Generates random data needed to create objects '''
class DataGenerator():
	words = []
	numbers = '1234567890'
	nums = range(9)
	url_specials = '._~:/?#[]@!$&\'()*+,;=' 
	local_specials = '!#$%&\'*+-/=?^_`{|}~'
	domain_specials = '-.'

	@classmethod
	def set_words(self, words):
		self.words = words

	def rand_word(self):
		return random.choice(self.words)

	def rand_numbers(self, num_digits):
		# digits = []
		# for _ in range(digits):
		# 	digits.append(random.choice(self.numbers))
		# this list comprehension is totally equal to these 3 lines. look for this pattern!
		digits = [random.choice(self.numbers) for _ in xrange(num_digits)]
		return ''.join(digits)

	def rand_url_specials(self):
		return random.choice(self.url_specials)

	def rand_local_specials(self):
		return random.choice(self.local_specials)

	def rand_domain_specials(self):
		return random.choice(self.domain_specials)

	''' a screename can be 20 characters, consisting of letters, digits, and
	 ._~:/?#[]@!$&'()*+,;= '''
	def generate_screenname(self):
		screenname = ''.join([self.rand_word(), self.rand_numbers(2), self.rand_url_specials()])
		while len(screenname) > 20:
			screenname = screenname[1:]
		return screenname

	''' local may have up to 64 chars, consisting of letter, digits, and
	 !#$%&'*+-/=?^_`{|}~
	 domain may have up to 255, consisting of letters, digits, hyphens and dots '''
	def generate_email(self):
		email = ''.join([self.rand_word(), self.rand_numbers(1), self.rand_local_specials(), 
			'@', self.rand_word(), self.rand_numbers(2), self.rand_domain_specials(), self.rand_word()])
		while len(email) > 255:
			email = email[1:]
		return email

	''' can be 60 chars, consisiting of letters, digits, and
	 ._~:/?#[]@!$&'()*+,;= '''
	def generate_password(self):
		password = ''.join([self.rand_word(), self.rand_numbers(2), self.rand_url_specials()])
		while len(password) > 60:
			password = password[1:]
		return password

	''' first name can be 255 chars '''
	def generate_first_name(self):
		first_name = ''.join([self.rand_word()])
		while len(first_name) > 255:
			first_name = first_name[1:]
		return first_name

	''' last name can be 255 chars '''
	def generate_last_name(self):
		last_name = ''.join([self.rand_word()])
		while len(last_name) > 255:
			last_name = last_name[1:]
		return last_name

	''' org name can be 255 chars or digits '''
	def generate_org_name(self):
		org_types = ['university', 'college', 'school']
		org_name = ''.join([self.rand_word(), ' ', random.choice(org_types)])
		while len(org_name) > 255:
			org_name = org_name[1:]
		return org_name

	def generate_org_slug(self):
		slug = ''.join([self.rand_word()])
		while len(slug) > 255:
			slug = slug[1:]
		return slug

	def generate_location(self):
		location = ''.join([self.rand_word(), self.rand_url_specials(), self.rand_word(), ', ID'])
		while len(location) > 255:
			location = location[1:]
		return location

	def generate_game_name(self):
		game_types = ['spring', 'summer', 'fall', 'winter']
		game_name = ''.join([random.choice(game_types), ' ', self.rand_word()])
		while len(game_name) > 255:
			game_name = game_name[1:]
		return game_name	

	def generate_game_slug(self):
		slug = ''.join([self.rand_word()])
		while len(slug) > 255:
			slug = slug[1:]
		return slug

	''' generates a json serializable datetime string given the number of days past the current time to add '''
	def generate_datetime(self, days):
		date_time = datetime.datetime.now() + datetime.timedelta(days)
		formatted_date_time = date_time.isoformat()
		return json.dumps(formatted_date_time)


''' Generates json to be sent using data generated in DataGenerator '''
class ObjectGenerator():
	def __init__(self, test_data_generator):
		self.test_data_generator = test_data_generator

	def generate_user_data(self):
		user = {
			'first_name': self.test_data_generator.generate_first_name(),
			'last_name': self.test_data_generator.generate_last_name(),
			'email': self.test_data_generator.generate_email(),
			'screen_name': self.test_data_generator.generate_screenname(),
			'password': self.test_data_generator.generate_password(),

		}
		return {'users': [user]}

	def generate_login_data(self, user):
		login_credentials = {
			'email': user['users'][0]['email'],
			'password': user['users'][0]['password'],
		}
		return {'users':[user]}

	def generate_org_data(self):
		org = {
			'name': self.test_data_generator.generate_org_name(),
			'location': self.test_data_generator.generate_location(),
		}
		slug = org['name'].split()
		org['slug'] = ''.join([slug[1][:1], 'of', slug[0]])
		return {'organizations': [org]}

	def generate_game_data(self):
		org = TestCaseOrg().test_create()
		org_id = org[0]['id']
		# print org
		game = {
			'name': self.test_data_generator.generate_game_name(),
			'slug': self.test_data_generator.generate_game_slug(),
			'organization_id': org_id,
			'timezone': 'US/Pacific',
			'registration_start_time': self.test_data_generator.generate_datetime(0),
			'registration_end_time': self.test_data_generator.generate_datetime(1),
			'running_start_time': self.test_data_generator.generate_datetime(2),
			'running_end_time': self.test_data_generator.generate_datetime(5),
		}
		return {'games': [game]}

	def generate_member_data(self):
		user = TestCaseUser().test_create()
		org = TestCaseOrg().test_create()
		member = {
			'user_id': user[0]['id'],
			'organization_id': org[0]['id'],
		}
		return {'members': [member]}

''' Generates an API request using the json data from ObjectGenerator '''
class APIRequestsGenerator():
	base_url = 'http://localhost:9000/api/'

	def post(self, test_object, url):
		url = ''.join([self.base_url, url])
		r = requests.post(url, json.dumps(test_object))	
		return r

	def bad_post(self, test_object, url):
		url = ''.join([self.base_url, url])
		r = requests.post(url, test_object)	
		return r

	def get(self, url):
		url = ''.join([self.base_url, url])
		r = requests.get(url)
		return r

if __name__ == '__main__':
	main()