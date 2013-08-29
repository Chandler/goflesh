App.UserSettingsController = BaseController.extend
  needs: 'user'
  user: null
  userBinding: 'controllers.user'
  contentBinding: 'user'

App.UsersNewController = NewController.extend
  editableRecordFields: ['first_name', 'last_name', 'email', 'screen_name', 'password']
  first_name: 'n',
  last_name: 'n',
  email: 'n',
  screen_name: 'n',
  password: 'n'

App.UserController = BaseObjectController.extend
  userIsCurrentUser: (->
    id = App.Auth.get('user.id')
    return (id && id == @get('content.id'))
  ).property()

App.UsersController = Ember.ObjectController.extend
  selectedUser: null

App.UserHomeController = Ember.Controller.extend
  needs: 'user'
  user: null
  userBinding: 'controllers.user'


