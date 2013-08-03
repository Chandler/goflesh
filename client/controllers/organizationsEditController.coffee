define ["ember"], (Em) ->
  OrganizationsNewController = Em.ObjectController.extend
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
          @transitionTo('organizations.show', record);
      else
        @set 'errors', 'Empty Fields'
    errors: null,
    clearErrors: ->
      @set 'errors', null
    errorMessages: (->
      @get 'errors'
    ).property 'errors' 