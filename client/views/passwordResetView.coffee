App.PasswordResetView = Ember.View.extend
  templateName: "passwordReset"
  didInsertElement: ->
    console.log("inserted")
    router = this.get('controller.target.router');
    # router.transitionTo('login')
