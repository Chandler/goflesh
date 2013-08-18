#plugins that need to run once to attach themselves to their parents.
require(['jquery-cookie','handlebars_helpers', 'ember-grid'])

#require parts of our framework not on the ember namespace
#we shouldn't need this, but things break without it
#tracking here: https://github.com/Chandler/flesh/issues/5
require(['NewController', 'BaseController'])

#Any file you create that needs be on the ember namespace must be listed here
ember_namespace = [
  "Auth",
  "Router",
  "Store",
  
  # Models
  "Player",
  "Member",
  "Game",
  "User",
  "Organization",
  
  # Routes
  "DiscoveryRoute",
  
  "GameRoute",
  "GamesNewRoute",
  
  "UserRoute",
  "UsersNewRoute",
  
  "OrganizationRoute",
  "OrganizationsNewRoute",
  "IndexRoute",

  # Controllers
  "ApplicationController"
  "LoginController",
  "DiscoveryController",

  "OrganizationHomeController",
  "OrganizationSettingsController",
  "OrganizationsController",
  "OrganizationsNewController",
  
  "UsersController",
  "UsersNewController",
  "UserHomeController",
  "UserSettingsController",
  
  "GamesController",
  "GamesNewController",
  "GameSettingsController",
  "GameHomeController",

  # Views
  "RegisterKillView",
  "TimeSeriesView",
  "PlayerTableRowView",
  "PlayerListTableView",  
  "ApplicationView",
  "LoginView",
  "DiscoveryView",
  "ListItemView",
  "GameItemView"

]
    
  


