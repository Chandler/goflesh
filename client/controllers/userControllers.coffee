App.UserSettingsController = BaseController.extend
  needs: 'user'
  user: null
  userBinding: 'controllers.user'
  contentBinding: 'controllers.user.content'
  
App.UserController = BaseController.extend()

App.UsersNewController = NewController.extend
  editableRecordFields: ['first_name', 'last_name', 'email', 'screen_name', 'password']
  first_name: '',
  last_name: '',
  email: '',
  screen_name: '',
  password: ''


App.UsersController = Ember.ObjectController.extend
  selectedUser: null

App.UserHomeController = Ember.Controller.extend
  needs: 'user'
  user: null
  userBinding: 'controllers.user'


