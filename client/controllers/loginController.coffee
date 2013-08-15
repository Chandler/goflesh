define ["ember", "BaseController"], (Em, BaseController) ->
  LoginController = BaseController.extend
    email: null
    password: null
    login: (arg) ->
      @clearErrors()
      Em.App.Auth.signIn
        data:
          email: @email
          password: @password

      Em.App.Auth.on 'signInSuccess', =>
        @transitionTo('user.home', Em.App.Auth.get('user'))
