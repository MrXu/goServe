System.register(['angular2/platform/browser', './components/TopNavLayout'], function(exports_1, context_1) {
    "use strict";
    var __moduleName = context_1 && context_1.id;
    var browser_1, TopNavLayout_1;
    return {
        setters:[
            function (browser_1_1) {
                browser_1 = browser_1_1;
            },
            function (TopNavLayout_1_1) {
                TopNavLayout_1 = TopNavLayout_1_1;
            }],
        execute: function() {
            browser_1.bootstrap(TopNavLayout_1.TopNavLayout);
        }
    }
});
//# sourceMappingURL=main.js.map