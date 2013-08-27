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
      @transitionToRoute('user.home', App.User.find(App.Auth.get('userId')))
