define ["ember"], (Em) ->
  LoginController = Ember.ObjectController.extend
    login: (arg) ->
      console.log "test"
      Em.App.Auth.signIn
        data:
          email: 'cats'
          password: 'cats'
