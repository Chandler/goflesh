define ["ember"], (Em) ->
  OrganizationHomeController = Em.Controller.extend
    needs: 'organization'
    organization: null
    organizationBinding: 'controllers.organization'