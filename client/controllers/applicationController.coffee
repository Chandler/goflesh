define ["ember"], (Em) ->
  ApplicationController = Em.Controller.extend
    signOut:  ->
      console.log "ok"
      Em.App.Auth.get('module.rememberable').forget()
      Em.App.Auth.set('signedIn', false)
