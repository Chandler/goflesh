define ["ember", "BaseController"], (Em, BaseController) ->
  GameSettingsController = BaseController.extend
    needs: 'game'
    game: null
    gameBinding: 'controllers.game'
    edit: ->
      this.clearErrors()
      if @get('game.name') != ''
        record = @get('game.content')
        @get('store').get('defaultTransaction').commit()
        record.on 'becameError', =>
          @set 'errors', 'SERVER ERROR'
        record.on 'didUpdate', =>
          @transitionTo('game.home', record);
      else
        @set 'errors', 'Empty Field'
