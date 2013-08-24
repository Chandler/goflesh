App.Auth = Ember.Auth.create
  modules: ['emberData']

  userModel: 'App.User' # default null
  # requestAdapter: 'jquery' # default 'jquery'
  # responseAdapter: 'json' # default 'json'
  # strategyAdapter: 'token' # default 'token'
  signInEndPoint: 'api/users/authenticate'

  tokenKey: 'api_key'
  tokenIdKey: 'id'

  sessionAdapter: 'cookie'

  modules: ['emberData', 'authRedirectable', 'actionRedirectable', 'rememberable']

  # authRedirectable:
  #   route: 'sign-in'

  # actionRedirectable:
  #   signInRoute: 'users'
  #   signInSmart: true
  #   signInBlacklist: ['sign-in']
  #   signOutRoute: 'posts'

  # rememberable:
  #   tokenKey: 'api_key'
  #   tokenIdKey: 'id'
  #   period: 7
  #   autoRecall: true