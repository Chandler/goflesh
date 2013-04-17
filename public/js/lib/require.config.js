var jam = {
    "packages": [
        {
            "name": "handlebars.runtime",
            "location": "lib/handlebars.runtime",
            "main": "handlebars.runtime-1.0.rc.1.js"
        },
        {
            "name": "jquery",
            "location": "lib/jquery",
            "main": "dist/jquery.js"
        }
    ],
    "version": "0.2.17",
    "shim": {
        "handlebars.runtime": {
            "exports": "Handlebars"
        }
    }
};

if (typeof require !== "undefined" && require.config) {
    require.config({
    "packages": [
        {
            "name": "handlebars.runtime",
            "location": "lib/handlebars.runtime",
            "main": "handlebars.runtime-1.0.rc.1.js"
        },
        {
            "name": "jquery",
            "location": "lib/jquery",
            "main": "dist/jquery.js"
        }
    ],
    "shim": {
        "handlebars.runtime": {
            "exports": "Handlebars"
        }
    }
});
}
else {
    var require = {
    "packages": [
        {
            "name": "handlebars.runtime",
            "location": "lib/handlebars.runtime",
            "main": "handlebars.runtime-1.0.rc.1.js"
        },
        {
            "name": "jquery",
            "location": "lib/jquery",
            "main": "dist/jquery.js"
        }
    ],
    "shim": {
        "handlebars.runtime": {
            "exports": "Handlebars"
        }
    }
};
}

if (typeof exports !== "undefined" && typeof module !== "undefined") {
    module.exports = jam;
}