App.SendPasswordResetController = BaseController.extend
  email: null
  reset: (arg) ->
    @clearErrors()
    @clearMessages()
    $.ajax "/api/password_reset",
      data: JSON.stringify(email: this.email)
      type: "POST"
      processData: false
      contentType: "application/json" 
    .done =>
      @set 'messages', this.email
    .fail => 
      @set 'errors', 'SERVER ERROR' 
