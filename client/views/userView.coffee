App.UserView = Ember.View.extend
  templateName: "user"
  upload: ->
    $("input[name=avatar]").click()

  didInsertElement: ->
    $('.file_upload').on 'click', (e) =>
      e.stopPropagation()
      @upload()

    #our avatar.io key, hardcoding for now
    token = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJwcml2YXRlX3Rva2VuIjoiYjdhZGFkNjI2YWI0NTc0ZmMwYTE5M2MyOWQ0YjBiMTkxYjQxYWNlOWQwMzUzNTU5YzhhMzRhMDhkOWRiNGI2ZCJ9.V0JQ5FygLaWiXn1yzN-gdmywkki6V1V8r2y1TqSMPUU"
    client = new AvatarsIO(token) # obtain at http://avatars.io/
    $ ->
      uploader = client.create("#avatar") # selector for input[type="file"] field, here #avatar, for example
      uploader.setIdentifier(App.Auth.get('user.id'));
      uploader.setAllowedExtensions ["png", "jpg"] # optional, defaults to png, gif, jpg, jpeg
      uploader.on "complete", (url) ->
        $('.profile_image').attr('src', url)
