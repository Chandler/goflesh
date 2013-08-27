App.ApplicationController = Ember.Controller.extend
  signOut:  ->
    App.Auth.destroySession()
  currentUser: (->
    App.Auth.get('user')
  ).property()