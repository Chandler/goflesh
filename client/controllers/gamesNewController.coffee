define ["ember", "ember-data"], (Em, DS) ->
  GamesNewController = Ember.ObjectController.extend
    name: '',
    slug: '',
    createGame: ->
      this.clearErrors()
      if this.name != ''
        model = this.get('model')
        record = model.createRecord
          name: this.name
          slug: this.slug
          # organization_id = ???
        record.transaction.commit()
        record.becameError =  =>
          @set 'errors', 'SERVER ERROR'
        record.didCreate = =>
          @transitionToRoute('games.show', record.id);
      else
        @set 'errors', 'Empty Fields'