define ["ember"], (Em) ->
  OrganizationSettingsController = Em.Controller.extend
    needs: 'organization'
    organization: null
    organizationBinding: 'controllers.organization'
    editOrg: ->
      console.log @get('organization')
      this.clearErrors()
      if this.name != ''
        record = @get('organization')
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