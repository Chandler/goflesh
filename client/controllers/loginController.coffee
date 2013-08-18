App.LoginController = BaseController.extend
  email: null
  password: null
  login: (arg) ->
    @clearErrors()
    App.Auth.signIn
      data:
        email: @email
        password: @password

    App.Auth.on 'signInSuccess', =>
      @transitionTo('user.home', App.Auth.get('user'))
