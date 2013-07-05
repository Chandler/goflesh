define ["ember"], (Em) ->
  LoginController = Ember.ObjectController.extend
    login: (arg) ->
      console.log "test"
      Em.App.Auth.signIn
        data:
          email: 'cbabraham@gmail.com'
          password: 'asdf'