System.register(['angular2/platform/browser', 'rxjs/Rx', 'angular2/http', './components/TopNavLayout'], function(exports_1, context_1) {
    "use strict";
    var __moduleName = context_1 && context_1.id;
    var browser_1, http_1, TopNavLayout_1;
    return {
        setters:[
            function (browser_1_1) {
                browser_1 = browser_1_1;
            },
            function (_1) {},
            function (http_1_1) {
                http_1 = http_1_1;
            },
            function (TopNavLayout_1_1) {
                TopNavLayout_1 = TopNavLayout_1_1;
            }],
        execute: function() {
            browser_1.bootstrap(TopNavLayout_1.TopNavLayout, [http_1.HTTP_PROVIDERS]);
        }
    }
});
//# sourceMappingURL=main.js.map