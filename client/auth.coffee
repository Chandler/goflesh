define ["ember-auth"], (Auth) ->
  Auth.create
    requestAdapter: 'jquery' # default 'jquery'
    responseAdapter: 'json' # default 'json'
    strategyAdapter: 'token' # default 'token'
    signInEndPoint: '/users/authenticate'

    tokenKey: 'auth_token'
    tokenIdKey: 'user_id'

    modules: ['emberData', 'authRedirectable', 'actionRedirectable', 'rememberable']

    authRedirectable:
      route: 'sign-in'

    actionRedirectable:
      signInRoute: 'users'
      signInSmart: true
      signInBlacklist: ['sign-in']
      signOutRoute: 'posts'

    rememberable:
      tokenKey: 'remember_token'
      period: 7
      autoRecall: true