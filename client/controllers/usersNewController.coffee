define ["NewController"], (NewController) ->
  UsersNewController = NewController.extend
    recordProperties = ['first_name', 'last_name', 'email', 'screen_name', 'password']
    first_name: '',
    last_name: '',
    email: '',
    screen_name: '',
    password: ''
