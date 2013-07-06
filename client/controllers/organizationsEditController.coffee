define ["ember", "ember-data"], (Em, DS) ->
  OrganizationsNewController = Ember.ObjectController.extend
    editOrg: ->
      this.clearErrors()
      if this.name != ''
        record = @get('model')
        record.setProperties
          name: @get("name")
          slug: @get("slug")
        record.transaction.commit()
        record.becameError =  =>
          @set 'errors', 'SERVER ERROR'
        record.didUpdate = =>
          @transitionTo('orgs.show', record);
      else
        @set 'errors', 'Empty Fields'
    errors: null,
    clearErrors: ->
      @set 'errors', null
    errorMessages: (->
      @get 'errors'
    ).property 'errors' 