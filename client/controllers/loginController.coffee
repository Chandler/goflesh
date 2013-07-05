define ["ember"], (Em) ->
  LoginController = Ember.ObjectController.extend
    Login: (arg) ->
      @set 'filterString', arg