
App.OrganizationHomeController = BaseController.extend
  needs: 'organization'
  organization: null
  organizationBinding: 'controllers.organization'
  contentBinding: 'organization.games'
  

App.OrganizationSettingsController = BaseController.extend
  editableRecordFields: ['name', 'slug']
  needs: 'organization'
  organizationBinding: 'controllers.organization'
  contentBinding: 'controllers.organization.content'