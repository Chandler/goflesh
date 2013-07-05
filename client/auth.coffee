define ["ember-auth"], (Auth) ->
  Auth.create
    # requestAdapter: 'jquery' # default 'jquery'
    # responseAdapter: 'json' # default 'json'
    # strategyAdapter: 'token' # default 'token'
    signInEndPoint: 'api/users/authenticate'

    tokenKey: 'api_key'
    tokenIdKey: 'id'

    modules: ['emberData', 'authRedirectable', 'actionRedirectable', 'rememberable']

    # authRedirectable:
    #   route: 'sign-in'

    # actionRedirectable:
    #   signInRoute: 'users'
    #   signInSmart: true
    #   signInBlacklist: ['sign-in']
    #   signOutRoute: 'posts'

    rememberable:
      tokenKey: 'test'
      period: 7
      autoRecall: true