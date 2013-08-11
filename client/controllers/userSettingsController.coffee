define ["ember", "BaseController"], (Em, BaseController) ->
  UserSettingsController = BaseController.extend
    needs: 'user'
    user: null
    userBinding: 'controllers.user'
    edit: ->
      this.clearErrors()
      if @get('user.name') != ''
        record = @get('user.content')
        @get('store').get('defaultTransaction').commit()
        record.on 'becameError', =>
          @set 'errors', 'SERVER ERROR'
        record.on 'didUpdate', =>
          @transitionTo('user.home', record);
      else
        @set 'errors', 'Empty Field'
