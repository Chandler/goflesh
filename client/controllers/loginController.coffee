App.LoginController = BaseController.extend
  email: null
  password: null
  login: (arg) ->
    @clearErrors()
    response = App.Auth.signIn
      data:
        email: @email
        password: @password

    App.Auth.on 'signInSuccess', =>
      @transitionToRoute('user.home', App.User.find(App.Auth.get('userId')))
    
    App.Auth.on 'signInError',(e) =>
      App.Auth.destroySession()
      @set 'errors', 'password and username do not match'
