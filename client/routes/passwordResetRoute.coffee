App.PasswordResetRoute = Ember.Route.extend
  model: (params) ->
    console.log 'router'
    result = App.PasswordReset.find(params)
    # @transitionTo('login')
    console.log 'found object'
    result
    # if result
      # console.log result
    # console.log result['store']
    # console.log 'happy'
    # result.commit()
  
    # events: 
    #   error: (error, transition) ->
    #     console.log 'error'
    # $.ajax "/api/password_resets?code="+params["code"],
    #   type: "GET"
    #   contentType: "application/json" 
    # .done =>
    # 	console.log 'yay'
    # 	redirect: ->
    #   @transitionTo('login')
    #   # @set 'messages', this.email
    # .fail => 
    # 	console.log 'boo'
    # 	redirect: ->
    # 	  @transitionTo('discovery')
    #   @set 'errors', 'SERVER ERROR' 


      # if Em.Error
      #   # console.log Em.Error('message')
      #   console.log 'sad face'
      # else
      #   console.log 'yay'

     # redirect: ->
     #   if errors
     #     @transitionTo('discovery')
     #   else
     #     @transitionTo('login')

    
    # error: (error)->
  	 #  console.log error
    # setupController: (controller, model) ->
    #   @_super arguments...
    #   @controllerFor('passwordReset').set 'selectedPasswordReset', model
    #   console.log 'yayed'
