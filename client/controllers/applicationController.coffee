define ["ember"], (Em) ->
  ApplicationController = Em.Controller.extend
    signOut:  ->
      Em.App.Auth.destroySession()
    currentUser: (->
      Em.App.Auth.get('user')
    ).property()