define ["ember", "ember-data", "BaseController"], (Em, DS, BaseController) ->
  NewController = Em.ObjectController.extend
    create: ->
      this.clearErrors()
      if this.name != ''
        model = this.get('model')
        record = model.createRecord
          name: this.name
          slug: this.slug
        record.transaction.commit()
        record.becameError =  =>
          @set 'errors', 'SERVER ERROR'
        record.didCreate = =>
          @transitionToRoute('orgs/' + record.id);
      else
        @set 'errors', 'Empty Fields'
