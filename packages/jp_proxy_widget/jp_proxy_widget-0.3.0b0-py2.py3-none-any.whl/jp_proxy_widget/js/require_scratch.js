
var dummy_text = 'define([], function() { console.log("loading dummy..."); return " dummy value!! "; })';

define("JSProxyLoad", [], function() {
    console.log("define executing");
    //return "hello there!";
    return {
        normalize: function(name, _) {
            return name; // is this needed?
        },
        load: function (name, req, onload, config) {
            var text = module_name_to_text[name];
            onload.fromText(text);
            // delete the text? xxxx
        }
    };
});

// logs 'require executing'

require.defined("JSProxyLoad")
// false

require(["JSProxyLoad"], function(x) {
    console.log("got " + x);
});

require.defined("JSProxyLoad");



var loader_is_defined = false;

// Define the loader at the last minute...
var define_loader = function(element) {
    if (!loader_is_defined) {
        module_name_to_text = [];
        element.definejs(JSProxyLoad, [], function() {
            return {
                normalize: function(name, _) {
                    return name; // is this needed?
                },
                load: function (name, req, onload, config) {
                    var text = module_name_to_text[name];
                    onload.fromText(text);
                    // delete the text? xxxx
                }
            };
        });
        // self test.
        var test_name = "xxxdummy";
        var dummy_text = 'define([], function() { console.log("loading dummy..."); return " dummy value!! "; })';
        module_name_to_text[test_name] = dummy_text;
        element.requirejs([JSProxyLoad + "!" + test_name], function(the_value) {
            console.log("JSProxyLoad for xxx_dummy succeeded. " + the_value);
        });
    }
    loader_is_defined = true;
}

element = {};
element.requirejs = require;
element.definejs = define;
JSProxyLoad = "JSProxyLoad";

define_loader();

// basic usage
define("xxx1", [], function() { console.log("loading xxx1..."); return "xxx1 dummy value!! "; });

define("xxx2", [], function() { console.log("loading xxx2..."); return "xxx2 dummy value!! "; });

require(["xxx1", "xxx2"], function (v1, v2){
    console.log("values "+v1+" :: "+v2);
});

var redefiner = function(name) {
    var define_replacement = function (a, b, c) {
        var defined_name = name;
        var requirements;
        var defining_function;
        if ((typeof a) == "string") {
            console.log("renamed define "+a+" : "+name);
            defined_name = a;
            requirements = b;
            defining_function = c;
        } else {
            console.log("de-anonymized define: " + name);
            requirements = a;
            defining_function = b;
        }
        console.log("defining " + defined_name);
        define(defined_name, requirements, defining_function);
        if (defined_name != name) {
            console.log("and also defining " + name);
            define(name, requirements, defining_function);
        }
    };
    if ((((typeof define) !== "undefined") && (define !== null)) && (define.amd !== null)) {
        define_replacement.amd = define.amd;
    }
    return define_replacement;
};

// anonymous
(function (define) {
    define([], function() { console.log("loading yyy1..."); return "yyy1 dummy value!! "; });
})(redefiner("yyy1"));

// non-anonymous
(function (define) {
    define("zzz3", [], function() { console.log("loading yyy3..."); return "yyy3 dummy value!! "; });
})(redefiner("yyy3"));

require(["xxx1", "yyy1", "yyy3"], function (v1, v2, v3){
    console.log("values "+v1+" :: "+v2+" :: "+v3);
});

sd = 'define([], function() { console.log("loading www1..."); return "www1 dummy value!! "; });';

var js_text_fn = Function("define", sd);
js_text_fn(redefiner("www1"));


require(["xxx1", "yyy1", "yyy3", "www1"], function (v1, v2, v3, v4){
    console.log("values "+v1+" :: "+v2+" :: "+v3+" :: "+v4);
});
