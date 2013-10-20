App.LoginController = BaseController.extend
  email: null
  password: null
  login: (arg) ->
    @clearErrors()
    response = App.Auth.signIn
      data:
        email: @email
        screen_name: @email
        password: @password

    App.Auth.on 'signInSuccess', =>
      user_id = App.Auth.get('userId');
      user = App.User.find(user_id);
      App.Player.find({user_id: [user_id]})
      @transitionToRoute('user.home', user)

      
    App.Auth.on 'signInError',(e) =>
      App.Auth.destroySession()
      @set 'errors', 'password and username do not match'
