define ["ember", "PasswordReset", "jquery"], (Em, PasswordReset, $) ->
  PasswordResetRoute = Em.Route.extend
    model: (params) ->
      try
        result = PasswordReset.find(params)
      catch error
      	console.log 'sucks'
        
      # result = PasswordReset.find(params)
      # console.log result
      # console.log params
      console.log 'happy'

      # $.ajax "/api/password_resets?code="+params["code"],
      #   type: "GET"
      #   contentType: "application/json" 
      # .done =>
      # 	console.log 'yay'
      # 	redirect: ->
	     #  @transitionTo('login')
      #   # @set 'messages', this.email
      # .fail => 
      # 	console.log 'boo'
      # 	redirect: ->
      # 	  @transitionTo('discovery')
        # @set 'errors', 'SERVER ERROR' 


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

    # events: -> 
    #   console.log 'sup'
    #   error: (reason, transition) ->
    #     console.log 'error'
    # error: (error)->
  	 #  console.log error
    # setupController: (controller, model) ->
    #   @_super arguments...
    #   @controllerFor('passwordReset').set 'selectedPasswordReset', model
    #   console.log 'yayed'
