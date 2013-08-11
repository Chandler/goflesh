define ["ember"], (Em) ->
  UserHomeController = Em.Controller.extend
    needs: 'user'
    user: null
    userBinding: 'controllers.user'