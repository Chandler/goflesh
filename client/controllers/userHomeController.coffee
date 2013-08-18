App.UserHomeController = Ember.Controller.extend
  needs: 'user'
  user: null
  userBinding: 'controllers.user'