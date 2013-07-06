define [], ->
  Utilities =
    avatarTag: (hash, size) ->
      sizes =
        small: 50
        large: 100
      px = sizes[size]
      "<img src=\"http://www.gravatar.com/avatar/" + hash + "?s=" + px + "/>"

