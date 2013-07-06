define [], ->
  Utilities =
    avatarTag: (hash, size, options) ->
      sizes =
        small: 50
        large: 100
        profile: 150
      px = sizes[size]
      "<img class=\"" + options.hash.class +  "\" src=\"http://www.gravatar.com/avatar/" + hash + "?s=" + px + "\"/>"

